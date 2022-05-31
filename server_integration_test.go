package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSavingInvestmentsAndRetrivingThem(t *testing.T) {
	store := InMemoryInvestmentStore{}
	server := InvestmentServer{&store}
	investment_id := "a4321421"

	server.ServeHTTP(httptest.NewRecorder(), newPostSaveRequest(investment_id, 30))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newGetInvestmentAmountRequest(investment_id))
	assertStatus(t, response.Code, http.StatusOK)

	assertResponseBody(t, response.Body.String(), "30")
}
