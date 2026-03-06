package setscore

import (
	"github.com/cirobispo/sandbox/internal/common/scoring"
	"github.com/cirobispo/sandbox/internal/common/turning"
)

type OnSetScore func(scoreA, scoreB int, isTieBreak, done bool)
type ParamOption func(score *SetScore)

type SetScore struct {
	sideToBegin      turning.SideTurn
	maxEven          int
	confirmationSize int
	scoreA, scoreB   int

	onAfterScoreEvent []OnSetScore
}

func WithDefaultSet(sideToBegin turning.SideTurn) ParamOption {
	return func(score *SetScore) {
		score.sideToBegin = sideToBegin
		score.maxEven = 6
		score.confirmationSize = 2
	}
}

func WithSideSizeAndTieBreak(sideToBegin turning.SideTurn, size int, decidingGame, tieBreakForLastEven bool) ParamOption {
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

func (ss SetScore) executeOnAfterScoreEvent(scoreA, scoreB int) {
	done, isTieBreak := ss.Done(), ss.IsTieBreak()
	for i := range ss.onAfterScoreEvent {
		event := ss.onAfterScoreEvent[i]
		event(scoreA, scoreB, isTieBreak, done)
	}
}

func (ss *SetScore) getScores() (*int, *int) {
	sA, sB := &ss.scoreA, &ss.scoreB
	if ss.sideToBegin == turning.STB {
		sA, sB = &ss.scoreB, &ss.scoreA
	}

	return sA, sB
}

func (ss *SetScore) AddOnAfterScoreEvent(event OnSetScore) {
	ss.onAfterScoreEvent = append(ss.onAfterScoreEvent, event)
}

func (ss *SetScore) AddGameScore(score scoring.Scoring) {
	if ss.Done() || !score.Done() { // am I acepting more points?
		return
	}

	sA, sB := ss.getScores()
	sideToAdd := sA

	if score.Side() == scoring.SSB {
		sideToAdd = sB
	}

	*sideToAdd += 1
	ss.executeOnAfterScoreEvent(ss.scoreA, ss.scoreB)
}

func (ss SetScore) Result() (int, int) {
	a, b := ss.getScores()
	return *a, *b
}

func (ss SetScore) Done() bool {
	sA, sB := ss.Result()

	sideAWins := (sA >= ss.maxEven && sA-sB >= ss.confirmationSize)
	sideBWins := (sB >= ss.maxEven && sB-sA >= ss.confirmationSize)
	result := sideAWins || sideBWins

	return result
}

func (ss SetScore) IsTieBreak() bool {
	sA, sB := ss.Result()

	result := (sA >= ss.maxEven && sB >= ss.maxEven)

	return result
}

func (ss SetScore) Side() scoring.ScoringSide {
	return scoring.Side(ss)
}
