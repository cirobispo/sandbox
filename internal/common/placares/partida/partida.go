package partida

import (
	"errors"

	"github.com/cirobispo/sandbox/internal/common/placares"
	"github.com/cirobispo/sandbox/internal/common/turning"
)

type OnMatchScore func(scoreA, scoreB int, done bool)
type ParamOption func(score *Partida)

type Partida struct {
	sideToBegin    turning.TurningSide
	bestOf         int
	scoreA, scoreB int

	onAfterScoreEvent []OnMatchScore
}

func WithDefault() ParamOption {
	return WithSideAndSize(turning.TSA, 3)
}

func WithGrandSlam() ParamOption {
	return WithSideAndSize(turning.TSA, 5)
}

func WithSideAndSize(sideToBegin turning.TurningSide, bestOf int) ParamOption {
	return func(score *Partida) {
		score.sideToBegin = sideToBegin
		score.bestOf = bestOf
	}
}

func New(param ParamOption) Partida {
	result := Partida{
		scoreA:            0,
		scoreB:            0,
		onAfterScoreEvent: make([]OnMatchScore, 0),
	}

	if param != nil {
		param(&result)
	}

	return result
}

func (m Partida) executeOnAfterScoreEvent(scoreA, scoreB int) {
	done := m.Terminado()
	for i := range m.onAfterScoreEvent {
		event := m.onAfterScoreEvent[i]
		event(scoreA, scoreB, done)
	}
}

func (m *Partida) getScores() (*int, *int) {
	sA, sB := &m.scoreA, &m.scoreB
	if m.sideToBegin == turning.TSB {
		sA, sB = &m.scoreB, &m.scoreA
	}

	return sA, sB
}

func (m *Partida) AddOnAfterScoreEvent(event OnMatchScore) {
	m.onAfterScoreEvent = append(m.onAfterScoreEvent, event)
}

func (m *Partida) AddScore(score placares.EstadoEParametroPlacar) error {
	if m.Terminado() { // am I acepting more points?
		return errors.New("Match completed already.")
	}

	if score.Tipo() != placares.TPSet {
		return errors.New("This is not Set Score.")
	}

	if !score.Terminado() {
		return errors.New("Set is not completed.")
	}

	sA, sB := m.getScores()
	sideToAdd := sA

	if score.Lado() == placares.LPOposto {
		sideToAdd = sB
	}

	*sideToAdd += 1
	m.executeOnAfterScoreEvent(m.scoreA, m.scoreB)
	return nil
}

func (m Partida) Resultado() (int, int) {
	return m.scoreA, m.scoreB
}

func (m Partida) Terminado() bool {
	sA, sB := m.Resultado()

	amountToWin := (m.bestOf / 2) + (m.bestOf % 2)

	sideAWins := (sA >= amountToWin && sA-sB >= 1)
	sideBWins := (sB >= amountToWin && sB-sA >= 1)
	result := sideAWins || sideBWins

	return result
}

func (m Partida) Lado() placares.LadoDoPlacar {
	return placares.Lado(m)
}

func (s Partida) Tipo() placares.TipoDoPlacar {
	return placares.TPPartida
}
