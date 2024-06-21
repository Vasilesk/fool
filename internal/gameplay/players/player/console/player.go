//nolint:forbidigo
package console

import (
	"fmt"

	"github.com/vasilesk/fool/internal/gameplay/players/player"
	"github.com/vasilesk/fool/pkg/card"
	"github.com/vasilesk/fool/pkg/identity"
)

type console struct {
	name  string
	cards []card.Card
	trump card.Suit
}

func NewPlayer(cards []card.Card, trump card.Suit) (player.Player, error) {
	fmt.Println("Enter your name")

	var name string
	if _, err := fmt.Scanf("%s", &name); err != nil {
		return nil, fmt.Errorf("scanning name input: %w", err)
	}

	return &console{
		name:  name,
		cards: cards,
		trump: trump,
	}, nil
}

func (c *console) Name() string {
	return c.name
}

func (c *console) GetRoundAttackerStrategy(_ identity.Identity) player.RoundAttackerStrategy {
	return newAttacker(c)
}

func (c *console) GetRoundDefenderStrategy(_ identity.Identity) player.RoundDefenderStrategy {
	return newDefender(c)
}

func (c *console) TakeLostRound(cards []card.Card) {
	c.cards = append(c.cards, cards...)

	fmt.Printf("player %q takes lost round cards: %v\n", c.name, cards)
}

func (c *console) TakeDeck(cards []card.Card) {
	c.cards = append(c.cards, cards...)

	fmt.Printf("player %q takes deck cards: %v\n", c.name, cards)
}

func (c *console) CardsCount() int {
	return len(c.cards)
}
