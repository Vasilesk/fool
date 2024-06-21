package random

import (
	"math/rand"

	"github.com/vasilesk/fool/internal/gameplay/players/orderstrategy"
	"github.com/vasilesk/fool/internal/gameplay/players/player"
)

type random struct{}

func New() orderstrategy.Strategy {
	return random{}
}

func (r random) Order(players []player.Player) {
	rand.Shuffle(len(players), func(i, j int) { players[i], players[j] = players[j], players[i] })
}
