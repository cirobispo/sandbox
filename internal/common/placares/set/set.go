package set

import (
	"errors"

	"github.com/cirobispo/sandbox/internal/common/placares"
	"github.com/cirobispo/sandbox/internal/common/turning"
)

type OnSetScore func(scoreA, scoreB int, isTieBreak, done bool)
type ParamOption func(score *Set)

type Set struct {
	sideToBegin      turning.TurningSide
	maxEven          int
	confirmationSize int
	scoreA, scoreB   int

	onAfterScoreEvent []OnSetScore
}

func WithDefaultSet(sideToBegin turning.TurningSide) ParamOption {
	return func(score *Set) {
		score.sideToBegin = sideToBegin
		score.maxEven = 6
		score.confirmationSize = 2
	}
}

func WithSideSizeAndTieBreak(sideToBegin turning.TurningSide, size int, decidingGame, tieBreakForLastEven bool) ParamOption {
	return func(score *Set) {
		score.sideToBegin = sideToBegin
		score.maxEven = size
		score.confirmationSize = 2
		if decidingGame {
			score.confirmationSize--
		}
	}
}

func New(param ParamOption) Set {
	result := Set{
		scoreA:            0,
		scoreB:            0,
		onAfterScoreEvent: make([]OnSetScore, 0),
	}

	if param != nil {
		param(&result)
	}

	return result
}

func (s Set) executeOnAfterScoreEvent(scoreA, scoreB int) {
	done, isTieBreak := s.Terminado(), s.IsTieBreak()
	for i := range s.onAfterScoreEvent {
		event := s.onAfterScoreEvent[i]
		event(scoreA, scoreB, isTieBreak, done)
	}
}

func (s *Set) getScores() (*int, *int) {
	sA, sB := &s.scoreA, &s.scoreB
	if s.sideToBegin == turning.TSB {
		sA, sB = &s.scoreB, &s.scoreA
	}

	return sA, sB
}

func (s *Set) AddOnAfterScoreEvent(event OnSetScore) {
	s.onAfterScoreEvent = append(s.onAfterScoreEvent, event)
}

func (s *Set) AddScore(score placares.EstadoEParametroPlacar) error {
	if s.Terminado() { // am I acepting more points?
		return errors.New("Score completed already.")
	}

	if score.Tipo() != placares.TPJogo {
		return errors.New("This is not a Game Score.")
	}

	if !score.Terminado() { // am I acepting more points?
		return errors.New("Game is not completed.")
	}

	sA, sB := s.getScores()
	sideToAdd := sA

	if score.Lado() == placares.LPOposto {
		sideToAdd = sB
	}

	*sideToAdd += 1
	s.executeOnAfterScoreEvent(s.scoreA, s.scoreB)
	return nil
}

func (s Set) Resultado() (int, int) {
	return s.scoreA, s.scoreB
}

func (s Set) Lado() placares.LadoDoPlacar {
	return placares.Lado(s)
}

func (s Set) Tipo() placares.TipoDoPlacar {
	return placares.TPSet
}

func (s Set) Terminado() bool {
	sA, sB := s.Resultado()

	diff := s.confirmationSize
	if sA > s.maxEven || sB > s.maxEven {
		diff = 1
	}

	sideAWins := (sA >= s.maxEven && sA-sB >= diff)
	sideBWins := (sB >= s.maxEven && sB-sA >= diff)
	result := sideAWins || sideBWins

	return result
}

func (s Set) IsTieBreak() bool {
	sA, sB := s.Resultado()
	tie := s.maxEven
	if s.confirmationSize == 1 {
		tie = s.maxEven - 1
	}
	result := (sA >= tie && sB >= tie)

	return result
}
