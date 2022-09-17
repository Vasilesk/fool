package orderstrategy

import "github.com/vasilesk/fool/internal/game/players/player"

type Strategy interface {
	Order([]player.Player)
}
