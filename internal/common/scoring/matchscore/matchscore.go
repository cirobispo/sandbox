package matchscore

import (
	"errors"

	"github.com/cirobispo/sandbox/internal/common/scoring"
	"github.com/cirobispo/sandbox/internal/common/turning"
)

type OnMatchScore func(scoreA, scoreB int, done bool)
type ParamOption func(score *MatchScore)

type MatchScore struct {
	sideToBegin    turning.SideTurn
	bestOf         int
	scoreA, scoreB int

	onAfterScoreEvent []OnMatchScore
}

func WithDefault() ParamOption {
	return WithSideAndSize(turning.STA, 3)
}

func WithGrandSlam() ParamOption {
	return WithSideAndSize(turning.STA, 5)
}

func WithSideAndSize(sideToBegin turning.SideTurn, bestOf int) ParamOption {
	return func(score *MatchScore) {
		score.sideToBegin = sideToBegin
		score.bestOf = bestOf
	}
}

func New(param ParamOption) MatchScore {
	result := MatchScore{
		scoreA:            0,
		scoreB:            0,
		onAfterScoreEvent: make([]OnMatchScore, 0),
	}

	if param != nil {
		param(&result)
	}

	return result
}

func (m MatchScore) executeOnAfterScoreEvent(scoreA, scoreB int) {
	done := m.Done()
	for i := range m.onAfterScoreEvent {
		event := m.onAfterScoreEvent[i]
		event(scoreA, scoreB, done)
	}
}

func (m *MatchScore) getScores() (*int, *int) {
	sA, sB := &m.scoreA, &m.scoreB
	if m.sideToBegin == turning.STB {
		sA, sB = &m.scoreB, &m.scoreA
	}

	return sA, sB
}

func (m *MatchScore) AddOnAfterScoreEvent(event OnMatchScore) {
	m.onAfterScoreEvent = append(m.onAfterScoreEvent, event)
}

func (m *MatchScore) AddScore(score scoring.Scoring) error {
	if score.Type() != scoring.STSet {
		return errors.New("This is not a score for a match.")
	}
	if m.Done() || !score.Done() { // am I acepting more points?
		return errors.New("Score completed already.")
	}

	sA, sB := m.getScores()
	sideToAdd := sA

	if score.Side() == scoring.SSB {
		sideToAdd = sB
	}

	*sideToAdd += 1
	m.executeOnAfterScoreEvent(m.scoreA, m.scoreB)
	return nil
}

func (m MatchScore) Result() (int, int) {
	a, b := m.getScores()
	return *a, *b
}

func (m MatchScore) Done() bool {
	sA, sB := m.Result()

	amountToWin := (m.bestOf / 2) + (m.bestOf % 2)

	sideAWins := (sA >= amountToWin && sA-sB >= 1)
	sideBWins := (sB >= amountToWin && sB-sA >= 1)
	result := sideAWins || sideBWins

	return result
}

func (m MatchScore) Side() scoring.ScoringSide {
	return scoring.Side(m)
}

func (s MatchScore) Type() scoring.ScoringType {
	return scoring.STMatch
}
