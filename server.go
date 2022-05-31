package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type InvestmentStore interface {
	GetInvestmentAmount(id int) int
}

type InvestmentServer struct {
	store InvestmentStore
}

func (i *InvestmentServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	investment_id, _ := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/investments/"))

    amount := i.store.GetInvestmentAmount(investment_id)

    if amount == 0 {
        w.WriteHeader(http.StatusNotFound)
    }

	fmt.Fprint(w, i.store.GetInvestmentAmount(investment_id))
}
