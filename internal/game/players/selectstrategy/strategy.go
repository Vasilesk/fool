package selectstrategy

import (
	"github.com/vasilesk/fool/internal/game/players/player"
)

type Strategy interface {
	NextRound(wasLost bool) (
		attacker player.AttackerWithIdentity,
		defender player.DefenderWithIdentity,
		inGame bool,
	)
}
