package console

import (
	"fmt"

	"github.com/vasilesk/fool/pkg/card"
)

func readCard() (card.Card, error) {
	var cardStr string

	_, err := fmt.Scanf("%s", &cardStr)
	if err != nil {
		return card.Card{}, fmt.Errorf("scanning input: %w", err)
	}

	return card.NewFromString(cardStr) //nolint:wrapcheck
}
