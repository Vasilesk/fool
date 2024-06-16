package main

import (
	"fmt"
	"log"

	"github.com/vasilesk/fool/internal/game"
	deskstandard "github.com/vasilesk/fool/internal/game/cards/deck/standard"
	"github.com/vasilesk/fool/internal/game/gameplay"
	orderstrategyrandom "github.com/vasilesk/fool/internal/game/players/orderstrategy/random"
	"github.com/vasilesk/fool/internal/game/players/player"
	"github.com/vasilesk/fool/internal/game/players/player/console"
	selectstrategystandard "github.com/vasilesk/fool/internal/game/players/selectstrategy/standard"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln("error:", err)
	}

	log.Println("program ended")
}

func run() error {
	d := deskstandard.NewDeck()

	p1, err := console.NewPlayer(d.TakeMax(gameplay.MaxCardsOfPlayer), d.Trump().Suit())
	if err != nil {
		return fmt.Errorf("creating player 1: %w", err)
	}

	p2, err := console.NewPlayer(d.TakeMax(gameplay.MaxCardsOfPlayer), d.Trump().Suit())
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
