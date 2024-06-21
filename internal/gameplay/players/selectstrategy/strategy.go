package selectstrategy

import (
	"github.com/vasilesk/fool/internal/gameplay/players/player"
)

type Strategy interface {
	NextPlayers(wasLost bool) (
		inGame bool,
		attacker player.AttackerWithIdentity,
		defender player.DefenderWithIdentity,
		moreAttackers []player.AttackerWithIdentity,
	)
}
