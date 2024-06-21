package player

import (
	"github.com/vasilesk/fool/pkg/card"
	"github.com/vasilesk/fool/pkg/identity"
)

type Defender interface {
	GetRoundDefenderStrategy(defender identity.Identity) RoundDefenderStrategy

	TakeLostRound(cards []card.Card)
}

type RoundDefenderStrategy interface {
	AnswerMove(move []card.Card) ([]card.Card, bool)
}

type DefenderWithIdentity interface {
	Defender
	identity.Identity
}
