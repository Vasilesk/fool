package main

import (
	"fmt"
	"log"

	"github.com/vasilesk/fool/internal/gameplay"
	deskstandard "github.com/vasilesk/fool/internal/gameplay/cards/deck/standard"
	"github.com/vasilesk/fool/internal/gameplay/game"
	orderstrategyrandom "github.com/vasilesk/fool/internal/gameplay/players/orderstrategy/random"
	"github.com/vasilesk/fool/internal/gameplay/players/player"
	"github.com/vasilesk/fool/internal/gameplay/players/player/console"
	selectstrategystandard "github.com/vasilesk/fool/internal/gameplay/players/selectstrategy/standard"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln("error:", err)
	}

	log.Println("program ended")
}

func run() error {
	d := deskstandard.NewDeck()

	p1, err := console.NewPlayer(d.TakeMax(gameplay.MaxCardsOfPlayer), d.TrumpCard().Suit())
	if err != nil {
		return fmt.Errorf("creating player 1: %w", err)
	}

	p2, err := console.NewPlayer(d.TakeMax(gameplay.MaxCardsOfPlayer), d.TrumpCard().Suit())
	if err != nil {
		return fmt.Errorf("creating player 2: %w", err)
	}

	players := []player.Player{p1, p2}

	selstrat, err := selectstrategystandard.NewStrategyFFA(players)
	if err != nil {
		return fmt.Errorf("creating select strategy: %w", err)
	}

	g := game.NewGame(d, players, orderstrategyrandom.New(), selstrat)

	log.Println("game started")

	loser, err := g.Run()
	if err != nil {
		return fmt.Errorf("running game: %w", err)
	}

	log.Println("game ended, fool is", loser.Name())

	return nil
}
