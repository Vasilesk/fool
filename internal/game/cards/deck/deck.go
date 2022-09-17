package deck

import (
	"github.com/vasilesk/fool/pkg/card"
)

type Deck interface {
	TakeMax(lim int) []card.Card
	Trump() card.Card
}
