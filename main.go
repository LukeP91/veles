package main

import (
	"log"
	"net/http"
)

type InMemoryInvestmentStore struct{}

func (i *InMemoryInvestmentStore) GetInvestmentAmount(id string) int {
	return 123
}

func (i *InMemoryInvestmentStore) SaveInvestment(id string, amount int) {}

func main() {
	server := &InvestmentServer{&InMemoryInvestmentStore{}}
	log.Fatal(http.ListenAndServe(":2320", server))
}
