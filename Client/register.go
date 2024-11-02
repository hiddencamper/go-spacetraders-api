package main

import (
	APITools "github.com/hiddencamper/go-spacetraders-api/APITools"
)

type RegisterView struct {
	ListFactions APITools.ListFactions
	AgentSymbol  string
	email        string
	cursor       int
	subpage      int
}
