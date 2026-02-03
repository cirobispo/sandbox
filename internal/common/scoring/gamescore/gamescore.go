package gamescore

import (
	"math"

	"github.com/cirobispo/sandbox/internal/common/pointing"
	"github.com/cirobispo/sandbox/internal/common/pointing/point"
	"github.com/cirobispo/sandbox/internal/common/turning"
)

type OnGameScore func(scoreA, scoreB int, done bool)

type GameScore struct {
	startSide      turning.SideTurn
	decidingPoint  bool
	scoreA, scoreB int

	onAfterScoreEvent []OnGameScore
}

func New(startSide turning.SideTurn, decidingPoint bool) GameScore {
	return GameScore{
		startSide:         startSide,
		decidingPoint:     decidingPoint,
		scoreA:            0,
		scoreB:            0,
		onAfterScoreEvent: make([]OnGameScore, 0),
	}
}

func (gs *GameScore) getScores() (*int, *int) {
	sA, sB := &gs.scoreA, &gs.scoreB
	if gs.startSide == turning.STB {
		sA, sB = &gs.scoreB, &gs.scoreA
	}

	return sA, sB
}

func (gs *GameScore) AddOnAfterScoreEvent(event OnGameScore) {
	gs.onAfterScoreEvent = append(gs.onAfterScoreEvent, event)
}

func (gs GameScore) executeOnAfterScoreEvent(scoreA, scoreB int) {
	closeEnding := (scoreA >= 3 || scoreB >= 3) && scoreA != scoreB
	done := (scoreA >= 3 || scoreB >= 3) && closeEnding
	for i := range gs.onAfterScoreEvent {
		event := gs.onAfterScoreEvent[i]
		event(scoreA, scoreB, done)
	}
}

func (gs *GameScore) AddPoint(p point.Point) {
	if who := p.PointSide(); who != pointing.PSNone {
		sA, sB := gs.getScores()
		(*sA)++
		if who == pointing.PSOppositeSide {
			(*sB)++
			(*sA)--
		}

		gs.executeOnAfterScoreEvent(gs.scoreA, gs.scoreB)
	}
}

func (gs GameScore) Result() (int, int) {
	a, b := gs.getScores()
	return *a, *b
}

func (gs GameScore) Done() bool {
	sA, sB := gs.Result()
	result := (gs.decidingPoint && (sA > 3 || sB > 3)) ||
		(!gs.decidingPoint && (sA > 3 || sB > 3) && (math.Abs(float64(sA-sB)) > 1))
	return result
}
