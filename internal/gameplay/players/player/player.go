package player

import (
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
