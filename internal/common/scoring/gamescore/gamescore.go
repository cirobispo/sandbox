package gamescore

import (
	"math"

	"github.com/cirobispo/sandbox/internal/common/pointing"
	"github.com/cirobispo/sandbox/internal/common/pointing/point"
	"github.com/cirobispo/sandbox/internal/common/scoring"
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

func (g *GameScore) getScores() (*int, *int) {
	sA, sB := &g.scoreA, &g.scoreB
	if g.startSide == turning.STB {
		sA, sB = &g.scoreB, &g.scoreA
	}

	return sA, sB
}

func (g *GameScore) inverseScore(side *int) *int {
	sA, sB := g.getScores()
	if side == sA {
		return sB
	}
	return sA
}

func (g *GameScore) AddOnAfterScoreEvent(event OnGameScore) {
	g.onAfterScoreEvent = append(g.onAfterScoreEvent, event)
}

func (g GameScore) executeOnAfterScoreEvent(scoreA, scoreB int) {
	done := g.Done()
	for i := range g.onAfterScoreEvent {
		event := g.onAfterScoreEvent[i]
		event(scoreA, scoreB, done)
	}
}

func (g *GameScore) AddPoint(p point.Point) {
	if g.Done() { // verify only it still acepting more points.
		return
	}

	if who := p.Side(); who != pointing.PSNone {
		incr := 1
		sA, sB := g.getScores()
		sideToAdd := sA
		if who == pointing.PSOppositeSide {
			sideToAdd = sB
		}

		if (!g.decidingPoint) && (*sA > 3 || *sB > 3) {
			if *sideToAdd == 3 {
				incr = -1
				sideToAdd = g.inverseScore(sideToAdd)
			}
		}
		*sideToAdd += incr
		g.executeOnAfterScoreEvent(g.scoreA, g.scoreB)
	}
}

func (g GameScore) Done() bool {
	sA, sB := g.Result()
	result := (g.decidingPoint && (sA > 3 || sB > 3)) ||
		(!g.decidingPoint && (sA > 3 || sB > 3) && (math.Abs(float64(sA-sB)) > 1))
	return result
}

func (g GameScore) Result() (int, int) {
	a, b := g.getScores()
	return *a, *b
}

func (g GameScore) Side() scoring.ScoringSide {
	return scoring.Side(g)
}

func (g GameScore) Type() scoring.ScoringType {
	return scoring.STGame
}
