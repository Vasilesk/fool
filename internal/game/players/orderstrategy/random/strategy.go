package random

import (
	"math/rand"
	"time"

	"github.com/vasilesk/fool/internal/game/players/orderstrategy"
	"github.com/vasilesk/fool/internal/game/players/player"
)

type random struct {
	seed int64
}

func New() orderstrategy.Strategy {
	return NewSeed(time.Now().UnixNano())
}

func NewSeed(s int64) orderstrategy.Strategy {
	return random{seed: s}
}

func (r random) Order(players []player.Player) {
	rand.Seed(r.seed)
	rand.Shuffle(len(players), func(i, j int) { players[i], players[j] = players[j], players[i] })
}
