package round

import (
	"errors"
	"fmt"

	"github.com/samber/lo"

	"github.com/vasilesk/fool/internal/gameplay"
	"github.com/vasilesk/fool/internal/gameplay/cards/beaten"
	"github.com/vasilesk/fool/internal/gameplay/players/player"
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

	attack := r.attacker.GetRoundAttackerStrategy(r.defender)
	defend := r.defender.GetRoundDefenderStrategy(r.attacker)

	move, err := attack.MakeMove(nil)
	if err != nil {
		return nil, false, fmt.Errorf("making first attack: %w", err)
	}

	if len(move) == 0 {
		return nil, false, errors.New("empty move")
	}

	wereBeaten := make([]beaten.Beaten, 0, maxBeaten)

	for len(move) > 0 {
		answer, answered := defend.AnswerMove(move)
		if !answered {
			return makeTaken(wereBeaten, move), true, nil
		}

		if len(answer) != len(move) {
			return nil, false, errors.New("answer length mismatch")
		}

		wereBeaten = append(wereBeaten,
			lo.FilterMap(lo.Zip2(move, answer), func(item lo.Tuple2[card.Card, card.Card], _ int) (beaten.Beaten, bool) {
				b, err := beaten.NewBeaten(item.A, item.B, r.trump)

				return b, err == nil
			})...,
		)

		if len(wereBeaten) >= maxBeaten {
			return beaten.GetCards(wereBeaten), false, nil
		} // todo: rules here

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
