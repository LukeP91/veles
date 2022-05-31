package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubInvestmentsStore struct {
	investments map[int]int
}

func (s *StubInvestmentsStore) GetInvestmentAmount(id int) int {
	amount := s.investments[id]
	return amount
}

func TestGETPlayers(t *testing.T) {
	store := StubInvestmentsStore{
		map[int]int{
			1: 20,
			2: 30,
		},
	}
	server := &InvestmentServer{&store}

	t.Run("returns Investment 1 amount", func(t *testing.T) {
		request := newGetInvestmentAmountRequest("1")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

        assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "20")
	})

	t.Run("returns Investment 2 amount", func(t *testing.T) {
		request := newGetInvestmentAmountRequest("2")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

        assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "30")
	})

	t.Run("returns 404 on missing investments", func(t *testing.T) {
		request := newGetInvestmentAmountRequest("3")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

        assertStatus(t, response.Code, http.StatusNotFound)
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

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}
