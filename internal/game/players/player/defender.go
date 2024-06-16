package player

import (
	"github.com/vasilesk/fool/internal/game/cards/beaten"
	"github.com/vasilesk/fool/pkg/card"
	"github.com/vasilesk/fool/pkg/identity"
)

type Defender interface {
	GetRoundDefenderStrategy(defender identity.Identity) RoundDefenderStrategy

	TakeLostRound(cards []card.Card)
}

type RoundDefenderStrategy interface {
	AnswerMove(move []card.Card, wereBeaten []beaten.Beaten) (moreBeaten []beaten.Beaten, answered bool)
}

type DefenderWithIdentity interface {
	Defender
	identity.Identity
}
