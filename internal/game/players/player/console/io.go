//nolint:forbidigo
package console

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/samber/lo"

	"github.com/vasilesk/fool/internal/game/cards/beaten"
	"github.com/vasilesk/fool/pkg/card"
)

func readCardsByNumbers(cards []card.Card) ([]card.Card, error) {
	fmt.Println("Choose card (enter option number):")
	lo.ForEach(cards, func(item card.Card, index int) {
		fmt.Println(item, "-", index+1)
	})
	fmt.Println("nothing - 0")

	var numbersStr string
	if _, err := fmt.Scanf("%s", &numbersStr); err != nil {
		return nil, fmt.Errorf("scanning input: %w", err)
	}

	res := make([]card.Card, 0, len(cards))

	for _, numStr := range strings.Split(numbersStr, " ") {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			return nil, fmt.Errorf("parsing number: %w", err) // todo: change to loop
		}

		if num == 0 {
			return nil, nil
		}

		ind := num - 1

		if ind < 0 || ind >= len(cards) {
			return nil, fmt.Errorf("invalid number: %d", num) // todo: change to loop
		}

		res = append(res, cards[ind])
	}

	return res, nil
}

func readCard() (card.Card, error) {
	var cardStr string

	if _, err := fmt.Scanf("%s", &cardStr); err != nil {
		return card.Card{}, fmt.Errorf("scanning input: %w", err)
	}

	return card.NewFromString(cardStr) //nolint:wrapcheck
}

func sortCards(cc []card.Card, t card.Suit) {
	sort.Slice(cc, func(i, j int) bool {
		if beaten.Beats(cc[i], cc[j], t) {
			return true
		}

		return cc[i].Suit() > cc[j].Suit() || cc[i].Weight() > cc[j].Weight()
	})
}
