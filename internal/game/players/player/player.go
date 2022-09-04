package player

import (
	"github.com/vasilesk/fool/internal/game/cards/beaten"
	"github.com/vasilesk/fool/pkg/card"
	"github.com/vasilesk/fool/pkg/identity"
)

type PlayersGetter interface {
	GetPlayers(identities []identity.Identity, trump card.Card) []Player
}

type Player interface {
	identity.Identity

	Attacker
	Defender

	TakeDeck(cards []card.Card)
	CardsCount() int
}

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

type StateForDefender interface {
	GetAttacker() identity.Identity
}
