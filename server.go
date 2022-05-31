package main

import (
	"fmt"
	"net/http"
	"strings"
)

type InvestmentStore interface {
	GetInvestmentAmount(id string) int
	SaveInvestment(id string, amount int)
}

type InvestmentServer struct {
	store InvestmentStore
}

func (i *InvestmentServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	investment_id := strings.TrimPrefix(r.URL.Path, "/investments/")

	switch r.Method {
	case http.MethodPost:
		i.saveInvestment(w, investment_id)
	case http.MethodGet:
		i.getInvestment(w, investment_id)
	}

}

func (i *InvestmentServer) getInvestment(w http.ResponseWriter, investment_id string) {
	amount := i.store.GetInvestmentAmount(investment_id)

	if amount == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, i.store.GetInvestmentAmount(investment_id))
}

func (i *InvestmentServer) saveInvestment(w http.ResponseWriter, investment_id string) {
	i.store.SaveInvestment(investment_id, 0)
	w.WriteHeader(http.StatusAccepted)
}
