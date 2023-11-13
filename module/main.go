package main

import (
	"github.com/loicpetit/go-input-and-bundle/input"
	"github.com/loicpetit/go-input-and-bundle/log"
)

func main() {
	log.Debug("Init")
	state, readError := input.Read()
	if readError != nil {
		panic(readError)
	}
	log.Debug("Nb nodes", state.NbNodes)
	log.Debug("Nb links", state.NbLinks)
	log.Debug("Nb exits", state.NbExits)
	for _, link := range state.Links {
		log.Debug("link between", link.Node1, "and", link.Node2)
	}
	for _, exit := range state.Exits {
		log.Debug("exit at", exit)
	}
}
