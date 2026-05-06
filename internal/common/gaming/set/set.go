package set

import (
	"errors"

	"github.com/cirobispo/sandbox/internal/common/gaming"
	"github.com/cirobispo/sandbox/internal/common/gaming/game"
	"github.com/cirobispo/sandbox/internal/common/placares"
	setscore "github.com/cirobispo/sandbox/internal/common/placares/set"
	"github.com/cirobispo/sandbox/internal/common/turning"
	"github.com/cirobispo/sandbox/internal/common/turning/turn"
)

type ParamOption func(set *Set) bool

type Set struct {
	whoServ, sideServ       *turn.Turn
	setSize                 int
	decidingPoint           bool
	tieBreak                bool
	score                   setscore.Set
	games                   []game.Gaming
	onAddingGameEvent       []gaming.OnAfterAddingGame
	onPlayerChangeSideEvent []gaming.OnPlayerChangeSide
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

func New(param ParamOption) *Set {
	result := &Set{
		games:                   make([]game.Gaming, 0, 13),
		onAddingGameEvent:       make([]gaming.OnAfterAddingGame, 0),
		onPlayerChangeSideEvent: make([]gaming.OnPlayerChangeSide, 0),
	}

	if param != nil {
		custom := param(result)

		serving_side := result.whoServ.StartSide()
		callback := setscore.WithDefaultSet(serving_side)
		if custom {
			callback = setscore.WithSideSizeAndTieBreak(serving_side, result.setSize, result.decidingPoint, result.tieBreak)
		}
		result.score = setscore.New(callback)
	}

	result.sideServ.AddOnAfterChange(func(ts turning.TurningSide) {
		result.executeOnPlayerChangeSide()
	})

	return result
}

func (s Set) executeOnAfterAddingGame(scoreA, scoreB int, done bool) {
	for j := range s.onAddingGameEvent {
		event := s.onAddingGameEvent[j]
		event(scoreA, scoreB, done)
	}
}

func (s Set) executeOnPlayerChangeSide() {
	for j := range s.onPlayerChangeSideEvent {
		event := s.onPlayerChangeSideEvent[j]
		event()
	}
}

func (s *Set) AddOnAddingGameEvent(event gaming.OnAfterAddingGame) {
	s.onAddingGameEvent = append(s.onAddingGameEvent, event)
}

func (s *Set) AddOnPlayerChangeEvent(event gaming.OnPlayerChangeSide) {
	s.onPlayerChangeSideEvent = append(s.onPlayerChangeSideEvent, event)
}

func (s *Set) AddGame(g game.Gaming) error {
	if s.score.Terminado() {
		return errors.New("set is closed.")
	}

	if !g.Score().Terminado() {
		return errors.New("game is still in play.")
	}

	s.games = append(s.games, g)
	s.whoServ.Execute()
	s.score.AddScore(g.Score())

	scoreA, scoreB := s.score.Resultado()
	done := s.score.Terminado()
	s.executeOnAfterAddingGame(scoreA, scoreB, done)
	if q := len(s.games) % 2; q == 1 {
		s.sideServ.Execute()
	}

	return nil
}

func (s Set) NewGame() *game.Game {
	newTurn := s.whoServ.Clone(s.whoServ.CurrentSide())
	result := game.New(newTurn, s.decidingPoint)
	return result
}

func (s Set) Score() placares.EstadoEResultadoPlacar {
	return s.score
}
