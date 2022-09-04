package standard

import (
	"fmt"

	"github.com/vasilesk/fool/internal/game/players/player"
	"github.com/vasilesk/fool/internal/game/players/selectstrategy"
)

type ffa struct {
	players []player.Player

	pos int
}

func (s *ffa) NextRound(wasLost bool) (player.AttackerWithIdentity, player.DefenderWithIdentity, bool) {
	if wasLost {
		s.switchPlaying()
	}

	fst := s.currentPlaying()

	if switched := s.switchPlaying(); !switched {
		return nil, nil, false
	}

	snd := s.currentPlaying()

	return fst, snd, true
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

func NewStrategyFFA(players []player.Player) (selectstrategy.Strategy, error) {
	if len(players) != 2 && len(players) != 3 {
		return nil, fmt.Errorf("%d", len(players))
	}

	return &ffa{players: players, pos: 0}, nil
}
