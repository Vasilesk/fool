package orderstrategy

import "github.com/vasilesk/fool/internal/gameplay/players/player"

type Strategy interface {
	Order(players []player.Player)
}
