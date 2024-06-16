//nolint:forbidigo
package console

import (
	"fmt"
	"slices"

	"github.com/samber/lo"

	"github.com/vasilesk/fool/internal/game/cards/beaten"
	"github.com/vasilesk/fool/pkg/card"
)

type defender struct {
	*console
}

func newDefender(c *console) *defender {
	return &defender{console: c}
}

func (d *defender) AnswerMove(move []card.Card) ([]card.Card, bool) {
	sortCards(d.cards, d.trump)

	fmt.Println(d.name, "defending from", move, "you have", d.cards, "trump is", d.trump)

	res := make([]card.Card, 0, len(move))

	for _, m := range move {
		c, err := readCard()
		if err != nil {
			return nil, false
		}

		if !slices.Contains(d.cards, c) {
			return nil, false
		}

		if !beaten.Beats(c, m, d.trump) {
			return nil, false
		}

		res = append(res, c)
	}

	d.cards = lo.Filter(d.cards, func(item card.Card, _ int) bool {
		return !slices.Contains(res, item)
	})

	return res, true
}
