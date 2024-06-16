package game

import (
	"errors"
	"fmt"

	"github.com/vasilesk/fool/internal/game/cards/deck"
	"github.com/vasilesk/fool/internal/game/gameplay"
	"github.com/vasilesk/fool/internal/game/gameplay/round"
	"github.com/vasilesk/fool/internal/game/players/orderstrategy"
	"github.com/vasilesk/fool/internal/game/players/player"
	"github.com/vasilesk/fool/internal/game/players/selectstrategy"
	"github.com/vasilesk/fool/pkg/card"
	"github.com/vasilesk/fool/pkg/identity"
)

type Game struct {
	deck           deck.Deck
	players        []player.Player
	orderStrategy  orderstrategy.Strategy
	selectStrategy selectstrategy.Strategy
}

func NewGame(
	deck deck.Deck,
	players []player.Player,
	orderStrategy orderstrategy.Strategy,
	selectStrategy selectstrategy.Strategy,
) *Game {
	return &Game{
		deck:           deck,
		players:        players,
		orderStrategy:  orderStrategy,
		selectStrategy: selectStrategy,
	}
}

func (g Game) Run() (identity.Identity, error) {
	var (
		cards []card.Card

		err error
	)

	taken := false

	for {
		attacker, defender, inGame := g.selectStrategy.NextRound(taken)
		if !inGame {
			break
		}

		cards, taken, err = round.NewRound(attacker, defender, g.deck.Trump().Suit()).Run()
		if err != nil {
			return nil, fmt.Errorf("running round: %w", err)
		}

		if taken {
			defender.TakeLostRound(cards)
		}

		for _, p := range g.players {
			if p.CardsCount() < gameplay.MaxCardsOfPlayer {
				p.TakeDeck(g.deck.TakeMax(gameplay.MaxCardsOfPlayer - p.CardsCount()))
			}
		}
	}

	for _, p := range g.players {
		if p.CardsCount() > 0 {
			return p, nil
		}
	}

	return nil, errors.New("unable to find user")
}
