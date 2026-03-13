package set

import (
	"errors"

	"github.com/cirobispo/sandbox/internal/common/gaming"
	"github.com/cirobispo/sandbox/internal/common/gaming/game"
	"github.com/cirobispo/sandbox/internal/common/scoring"
	"github.com/cirobispo/sandbox/internal/common/scoring/setscore"
	"github.com/cirobispo/sandbox/internal/common/turning"
	"github.com/cirobispo/sandbox/internal/common/turning/turn"
)

type ParamOption func(set *Set) bool

type Set struct {
	whoServ, sideServ *turn.Turn
	setSize           int
	decidingPoint     bool
	tieBreak          bool
	score             setscore.SetScore
	games             []game.Game
	onAddingGameEvent []gaming.OnAfterAddingGame
}

func WithDefaultSet(turnForServing *turn.Turn) ParamOption {
	return func(score *Set) bool {
		score.whoServ = turnForServing
		score.sideServ = turn.New(turn.WithTurningSide(turnForServing.StartSide()))
		score.setSize = 6
		score.decidingPoint = false
		score.tieBreak = true
		return false
	}
}

func WithTurnSizeAndTieBreak(turnForServing *turn.Turn, size int, decidingPoint, tieBreak bool) ParamOption {
	return func(score *Set) bool {
		score.whoServ = turnForServing
		score.sideServ = turn.New(turn.WithTurningSide(turnForServing.StartSide()))
		score.setSize = size
		score.decidingPoint = decidingPoint
		score.tieBreak = tieBreak
		return true
	}
}

func New(param ParamOption) Set {
	result := Set{
		games:             make([]game.Game, 0, 13),
		onAddingGameEvent: make([]gaming.OnAfterAddingGame, 0),
	}

	if param != nil {
		custom := param(&result)

		callback := setscore.WithDefaultSet(turning.TSA)
		if custom {
			serving_side := result.whoServ.StartSide()
			callback = setscore.WithSideSizeAndTieBreak(serving_side, result.setSize, result.decidingPoint, result.tieBreak)
		}
		result.score = setscore.New(callback)
	}

	return result
}

func (s *Set) AddOnAddingGameEvent(event gaming.OnAfterAddingGame) {
	s.onAddingGameEvent = append(s.onAddingGameEvent, event)
}

func (s *Set) AddGame(g game.Game) error {
	if s.Done() {
		return errors.New("set is closed.")
	}

	if !g.Score().Done() {
		return errors.New("game is still in play.")
	}

	s.games = append(s.games, g)
	s.whoServ.Execute()
	if q := len(s.games) % 2; q == 1 {
		s.sideServ.Execute()
	}

	s.score.AddScore(g.Score())

	scoreA, scoreB := s.score.Result()
	done := s.score.Done()
	s.executeOnAfterAddingGame(scoreA, scoreB, done)
	return nil
}

// func (s *Set) scoreToCompute() *int {
// 	s.sideServ.LastSide()
// 	return s.sc
// }

func (s Set) Score() scoring.ScoringResulting {
	return s.score
}

func (s Set) Done() bool {
	return s.score.Done()
}

func (s Set) executeOnAfterAddingGame(scoreA, scoreB int, done bool) {
	for j := range s.onAddingGameEvent {
		event := s.onAddingGameEvent[j]
		event(scoreA, scoreB, done)
	}
}
