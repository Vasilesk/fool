//nolint:forbidigo
package console

import (
	"errors"
	"fmt"
	"slices"
	"sort"

	"github.com/samber/lo"

	"github.com/vasilesk/fool/internal/game/cards/beaten"
	"github.com/vasilesk/fool/internal/game/players/player"
	"github.com/vasilesk/fool/pkg/card"
	"github.com/vasilesk/fool/pkg/identity"
)

type console struct {
	name  string
	cards []card.Card
	trump card.Suit
}

func NewPlayer(cards []card.Card, trump card.Suit) (player.Player, error) {
	fmt.Println("enter your name")

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
	return newAttack(c)
}

func (c *console) GetRoundDefenderStrategy(_ identity.Identity) player.RoundDefenderStrategy {
	return newDefend(c)
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

///////

type attack struct {
	*console
}

func newAttack(c *console) *attack {
	return &attack{console: c}
}

func (a *attack) MakeMove(_ []beaten.Beaten) ([]card.Card, error) {
	sortCards(a.cards, a.trump)

	fmt.Println(a.name, "make your move, you have", a.cards, "trump is", a.trump)

	c, err := readCard() // todo: for ...: more cards, process errs
	if err != nil {
		fmt.Println("no card provided")

		return nil, nil //nolint:nilerr
	}

	fmt.Println("your move is", c)

	if !slices.Contains(a.cards, c) {
		return nil, errors.New("no suitable card")
	}

	a.cards = lo.Filter(a.cards, func(item card.Card, _ int) bool {
		return item != c
	})

	return []card.Card{c}, nil
}

///////

type defend struct {
	*console
}

func newDefend(c *console) *defend {
	return &defend{console: c}
}

func (d *defend) AnswerMove(move []card.Card, _ []beaten.Beaten) ([]beaten.Beaten, bool) {
	sortCards(d.cards, d.trump)

	fmt.Println(d.name, "defending from", move, "you have", d.cards, "trump is", d.trump)

	res := make([]beaten.Beaten, 0, len(move))

	for _, m := range move {
		c, err := readCard()
		if err != nil {
			return nil, false
		}

		if !slices.Contains(d.cards, c) {
			return nil, false
		}

		b, err := beaten.NewBeaten(m, c, d.trump)
		if err != nil {
			return nil, false
		}

		res = append(res, b)
	}

	beatWith := lo.Map(res, func(item beaten.Beaten, _ int) card.Card {
		return item.GetAnswer()
	})

	d.cards = lo.Filter(d.cards, func(item card.Card, _ int) bool {
		return !slices.Contains(beatWith, item)
	})

	return res, true
}

func sortCards(cc []card.Card, t card.Suit) {
	sort.Slice(cc, func(i, j int) bool {
		if beaten.Beats(cc[i], cc[j], t) {
			return true
		}

		return cc[i].Suit() > cc[j].Suit() || cc[i].Weight() > cc[j].Weight()
	})
}
