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

func (gs *GameScore) inverseScore(side *int) *int {
	sA, sB := gs.getScores()
	if side == sA {
		return sB
	}
	return sA
}

func (gs *GameScore) AddOnAfterScoreEvent(event OnGameScore) {
	gs.onAfterScoreEvent = append(gs.onAfterScoreEvent, event)
}

func (gs GameScore) executeOnAfterScoreEvent(scoreA, scoreB int) {
	done := gs.Done()
	for i := range gs.onAfterScoreEvent {
		event := gs.onAfterScoreEvent[i]
		event(scoreA, scoreB, done)
	}
}

func (gs *GameScore) AddPoint(p point.Point) {
	if !gs.Done() { // verify only it still acepting more points.
		if who := p.PointSide(); who != pointing.PSNone {
			incr := 1
			sA, sB := gs.getScores()
			sideToAdd := sA
			if who == pointing.PSOppositeSide {
				sideToAdd = sB
			}

			if (!gs.decidingPoint) && (*sA > 3 || *sB > 3) {
				if *sideToAdd == 3 {
					incr = -1
					sideToAdd = gs.inverseScore(sideToAdd)
				}
			}
			*sideToAdd += incr
			gs.executeOnAfterScoreEvent(gs.scoreA, gs.scoreB)
		}
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
