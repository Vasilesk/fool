package round

import (
	"fmt"

	"github.com/vasilesk/fool/internal/game/cards/beaten"
	"github.com/vasilesk/fool/internal/game/gameplay"
	"github.com/vasilesk/fool/internal/game/players/player"
	"github.com/vasilesk/fool/pkg/card"
)

type Round struct {
	attacker player.AttackerWithIdentity
	defender player.DefenderWithIdentity

	trump card.Suit
}

func NewRound(
	attacker player.AttackerWithIdentity,
	defender player.DefenderWithIdentity,
	trump card.Suit,
) Round {
	return Round{
		attacker: attacker,
		defender: defender,
		trump:    trump,
	}
}

//nolint:nonamedreturns
func (r Round) Run() (cards []card.Card, taken bool, err error) {
	const maxBeaten = gameplay.MaxCardsOfPlayer

	wereBeaten := make([]beaten.Beaten, 0, maxBeaten)

	attack := r.attacker.GetRoundAttackerStrategy(r.defender)
	defend := r.defender.GetRoundDefenderStrategy(r.attacker)

	move, err := attack.MakeMove(wereBeaten)
	if err != nil {
		return nil, false, fmt.Errorf("making first attack: %w", err)
	}

	for len(move) > 0 {
		answer, answered := defend.AnswerMove(move, wereBeaten)
		if !answered {
			return makeTaken(wereBeaten, move), true, nil
		}

		wereBeaten = append(wereBeaten, answer...)
		if len(wereBeaten) == maxBeaten {
			return beaten.GetCards(wereBeaten), false, nil
		}

		move, err = attack.MakeMove(wereBeaten)
		if err != nil {
			return nil, false, fmt.Errorf("making attack: %w", err)
		}
	}

	return beaten.GetCards(wereBeaten), false, nil
}

func makeTaken(wereBeaten []beaten.Beaten, wereNotBeaten []card.Card) []card.Card {
	taken := beaten.GetCards(wereBeaten)
	taken = append(taken, wereNotBeaten...)

	return taken
}
