package deck

import (
	"github.com/vasilesk/fool/pkg/card"
)

type Deck interface {
	GetMax(n int) []card.Card
	Trump() card.Card
}
