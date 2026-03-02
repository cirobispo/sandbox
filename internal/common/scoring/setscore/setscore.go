package setscore

import (
	"math"

	"github.com/cirobispo/sandbox/internal/common/gaming/game"
	"github.com/cirobispo/sandbox/internal/common/turning"
)

type OnSetScore func(scoreA, scoreB int, isTieBreak, done bool)

type SetScore struct {
	startSide      turning.SideTurn
	setSize        int
	decidingSet    bool
	scoreA, scoreB int

	onAfterScoreEvent []OnSetScore
}

func New(startSide turning.SideTurn, decidingSet bool) SetScore {
	return SetScore{
		startSide:         startSide,
		setSize:           6,
		decidingSet:       decidingSet,
		scoreA:            0,
		scoreB:            0,
		onAfterScoreEvent: make([]OnSetScore, 0),
	}
}

func (ss *SetScore) getScores() (*int, *int) {
	sA, sB := &ss.scoreA, &ss.scoreB
	if ss.startSide == turning.STB {
		sA, sB = &ss.scoreB, &ss.scoreA
	}

	return sA, sB
}

func (ss *SetScore) inverseScore(side *int) *int {
	sA, sB := ss.getScores()
	if side == sA {
		return sB
	}
	return sA
}

func (ss *SetScore) AddOnAfterScoreEvent(event OnSetScore) {
	ss.onAfterScoreEvent = append(ss.onAfterScoreEvent, event)
}

func (ss SetScore) executeOnAfterScoreEvent(scoreA, scoreB int) {
	done, isTieBreak := ss.Done(), ss.isTieBreak()
	for i := range ss.onAfterScoreEvent {
		event := ss.onAfterScoreEvent[i]
		event(scoreA, scoreB, isTieBreak, done)
	}
}

func (ss *SetScore) AddGame(g *game.Game) {
	if ss.Done() { // am I acepting more points?
		return
	}

	sA, sB := ss.getScores()
	sideToAdd := sA

	if (!ss.decidingSet) && (*sA > ss.setSize || *sB > ss.setSize) {
		if *sideToAdd == 5 {
			sideToAdd = ss.inverseScore(sideToAdd)
		}
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
	result := (ss.decidingSet && (sA >= ss.setSize || sB >= ss.setSize)) ||
		(!ss.decidingSet && (sA >= ss.setSize || sB >= ss.setSize) && (math.Abs(float64(sA-sB)) > 1))
	return result
}

func (ss SetScore) isTieBreak() bool {
	sA, sB := ss.Result()
	result := (sA == ss.setSize && sB == ss.setSize)

	return result
}
