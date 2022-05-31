package main

import (
    "bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubInvestmentsStore struct {
	investments map[string]int
	saveCalls   []string
}

func (s *StubInvestmentsStore) GetInvestmentAmount(id string) int {
	amount := s.investments[id]
	return amount
}

func (s *StubInvestmentsStore) SaveInvestment(name string, amount int) {
	s.investments[name] = amount
	s.saveCalls = append(s.saveCalls, name)
}

func TestGETInvestments(t *testing.T) {
	store := StubInvestmentsStore{
		map[string]int{
			"a": 20,
			"b": 30,
		},
		nil,
	}
	server := &InvestmentServer{&store}

	t.Run("returns Investment 1 amount", func(t *testing.T) {
		request := newGetInvestmentAmountRequest("a")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "20")
	})

	t.Run("returns Investment 2 amount", func(t *testing.T) {
		request := newGetInvestmentAmountRequest("b")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "30")
	})

	t.Run("returns 404 on missing investments", func(t *testing.T) {
		request := newGetInvestmentAmountRequest("c")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusNotFound)
	})
}

func TestStoreInvestment(t *testing.T) {
	store := StubInvestmentsStore{
		map[string]int{},
		nil,
	}
	server := &InvestmentServer{&store}

	t.Run("it saves investment when POST", func(t *testing.T) {
		id := "a42141"
		amount := 30
		request := newPostSaveRequest(id, amount)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusAccepted)

		if len(store.saveCalls) != 1 {
			t.Errorf("got %d calls to SaveInvestment want %d", len(store.saveCalls), 1)
		}

		if store.saveCalls[0] != id {
			t.Errorf("did not store correct investment id got %q want %q", store.saveCalls[0], id)
		}

		if store.investments[id] != amount {
			t.Errorf("did not store correct investment amount got %q want %q", store.investments[id], amount)
		}
	})
}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}

func newGetInvestmentAmountRequest(id string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/investments/%s", id), nil)
	return req
}

func newPostSaveRequest(name string, amount int) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/investments/%s", name), bytes.NewBuffer(amount))
	return req
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}
