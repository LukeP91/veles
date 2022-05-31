package main

import (
	"log"
	"net/http"
)

type InMemoryInvestmentStore struct{}

func (i *InMemoryInvestmentStore) GetInvestmentAmount(id int) int {
	return 123
}

func main() {
	server := &InvestmentServer{&InMemoryInvestmentStore{}}
	log.Fatal(http.ListenAndServe(":2320", server))
}
