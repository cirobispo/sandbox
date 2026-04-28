package setscore

import (
	"errors"

	"github.com/cirobispo/sandbox/internal/common/scoring"
	"github.com/cirobispo/sandbox/internal/common/turning"
)

type OnSetScore func(scoreA, scoreB int, isTieBreak, done bool)
type ParamOption func(score *SetScore)

type SetScore struct {
	sideToBegin      turning.TurningSide
	maxEven          int
	confirmationSize int
	scoreA, scoreB   int

	onAfterScoreEvent []OnSetScore
}

func WithDefaultSet(sideToBegin turning.TurningSide) ParamOption {
	return func(score *SetScore) {
		score.sideToBegin = sideToBegin
		score.maxEven = 6
		score.confirmationSize = 2
	}
}

func WithSideSizeAndTieBreak(sideToBegin turning.TurningSide, size int, decidingGame, tieBreakForLastEven bool) ParamOption {
	return func(score *SetScore) {
		score.sideToBegin = sideToBegin
		score.maxEven = size
		score.confirmationSize = 2
		if decidingGame {
			score.confirmationSize--
		}
	}
}

func New(param ParamOption) SetScore {
	result := SetScore{
		scoreA:            0,
		scoreB:            0,
		onAfterScoreEvent: make([]OnSetScore, 0),
	}

	if param != nil {
		param(&result)
	}

	return result
}

func (s SetScore) executeOnAfterScoreEvent(scoreA, scoreB int) {
	done, isTieBreak := s.Terminado(), s.IsTieBreak()
	for i := range s.onAfterScoreEvent {
		event := s.onAfterScoreEvent[i]
		event(scoreA, scoreB, isTieBreak, done)
	}
}

func (s *SetScore) getScores() (*int, *int) {
	sA, sB := &s.scoreA, &s.scoreB
	if s.sideToBegin == turning.TSB {
		sA, sB = &s.scoreB, &s.scoreA
	}

	return sA, sB
}

func (s *SetScore) AddOnAfterScoreEvent(event OnSetScore) {
	s.onAfterScoreEvent = append(s.onAfterScoreEvent, event)
}

func (s *SetScore) AddScore(score scoring.EstadoEParametroPlacar) error {
	if s.Terminado() { // am I acepting more points?
		return errors.New("Score completed already.")
	}

	if score.Tipo() != scoring.TPJogo {
		return errors.New("This is not a Game Score.")
	}

	if !score.Terminado() { // am I acepting more points?
		return errors.New("Game is not completed.")
	}

	sA, sB := s.getScores()
	sideToAdd := sA

	if score.Lado() == scoring.LPOposto {
		sideToAdd = sB
	}

	*sideToAdd += 1
	s.executeOnAfterScoreEvent(s.scoreA, s.scoreB)
	return nil
}

func (s SetScore) Resultado() (int, int) {
	return s.scoreA, s.scoreB
}

func (s SetScore) Lado() scoring.LadoDoPlacar {
	return scoring.Lado(s)
}

func (s SetScore) Tipo() scoring.TipoDoPlacar {
	return scoring.TPSet
}

func (s SetScore) Terminado() bool {
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

func (s SetScore) IsTieBreak() bool {
	sA, sB := s.Resultado()
	tie := s.maxEven
	if s.confirmationSize == 1 {
		tie = s.maxEven - 1
	}
	result := (sA >= tie && sB >= tie)

	return result
}
