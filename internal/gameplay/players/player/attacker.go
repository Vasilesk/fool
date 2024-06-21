package player

import (
	"github.com/vasilesk/fool/internal/gameplay/cards/beaten"
	"github.com/vasilesk/fool/pkg/card"
	"github.com/vasilesk/fool/pkg/identity"
)

type Attacker interface {
	GetRoundAttackerStrategy(defender identity.Identity) RoundAttackerStrategy
}

type RoundAttackerStrategy interface {
	MakeMove(wereBeaten []beaten.Beaten) ([]card.Card, error)
}

type AttackerWithIdentity interface {
	Attacker
	identity.Identity
}
