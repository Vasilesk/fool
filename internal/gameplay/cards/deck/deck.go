package deck

import (
	"github.com/vasilesk/fool/pkg/card"
)

type Deck interface {
	TakeMax(lim int) []card.Card
	TrumpCard() card.Card
}
