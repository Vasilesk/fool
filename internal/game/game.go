package game

import (
	"errors"
	"fmt"

	"github.com/vasilesk/fool/internal/game/cards/deck"
	"github.com/vasilesk/fool/internal/game/gameplay"
	"github.com/vasilesk/fool/internal/game/gameplay/round"
	"github.com/vasilesk/fool/internal/game/players/player"
	"github.com/vasilesk/fool/internal/game/players/selectstrategy"
	"github.com/vasilesk/fool/pkg/card"
	"github.com/vasilesk/fool/pkg/identity"
)

type Game struct {
	deck     deck.Deck
	players  []player.Player
	strategy selectstrategy.Strategy
}

func (g Game) Run() (identity.Identity, error) {
	var (
		cards    []card.Card
		attacker player.AttackerWithIdentity
		defender player.DefenderWithIdentity

		err error
	)

	inGame := true
	taken := false

	for inGame {
		attacker, defender, inGame = g.strategy.NextRound(taken)
		if taken {
			defender.TakeLostRound(cards)
		}

		cards, taken, err = round.NewRound(attacker, defender, g.deck.Trump().Suit()).Run()
		if err != nil {
			return nil, fmt.Errorf("running round: %w", err)
		}

		for _, p := range g.players {
			p.TakeDeck(g.deck.GetMax(gameplay.MaxCardsOfPlayer - p.CardsCount()))
		}
	}

	for _, p := range g.players {
		if p.CardsCount() > 0 {
			return p, nil
		}
	}

	return nil, errors.New("unable to find user")
}
