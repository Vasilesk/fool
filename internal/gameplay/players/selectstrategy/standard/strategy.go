package standard

import (
	"fmt"

	"github.com/samber/lo"

	"github.com/vasilesk/fool/internal/gameplay/players/player"
	"github.com/vasilesk/fool/internal/gameplay/players/selectstrategy"
)

type ffa struct {
	players []player.Player

	pos int
}

func NewStrategyFFA(players []player.Player) (selectstrategy.Strategy, error) {
	if len(players) != 2 && len(players) != 3 {
		return nil, fmt.Errorf("%d", len(players))
	}

	return &ffa{players: players, pos: 0}, nil
}

func (s *ffa) NextPlayers(
	prevLost bool,
) (bool, player.AttackerWithIdentity, player.DefenderWithIdentity, []player.AttackerWithIdentity) {
	if prevLost {
		s.switchPlaying()
	}

	fst := s.currentPlaying()

	if switched := s.switchPlaying(); !switched {
		return false, nil, nil, nil
	}

	snd := s.currentPlaying()

	moreAttackers := lo.FilterMap(s.players, func(p player.Player, _ int) (player.AttackerWithIdentity, bool) {
		if p.Name() != fst.Name() && p.Name() != snd.Name() {
			return player.AttackerWithIdentity(p), true
		}

		return nil, false
	})

	return true, fst, snd, lo.Shuffle(moreAttackers)
}

func (s *ffa) switchPlaying() bool {
	oldPos := s.pos

	for i := 1; i < len(s.players); i++ {
		pos := (i + s.pos) % len(s.players)
		if s.players[pos].CardsCount() > 0 {
			s.pos = pos

			break
		}
	}

	return s.pos != oldPos
}

func (s *ffa) currentPlaying() player.Player {
	return s.players[s.pos%len(s.players)]
}
