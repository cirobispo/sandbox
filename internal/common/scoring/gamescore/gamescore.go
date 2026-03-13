package gamescore

import (
	"errors"

	"github.com/cirobispo/sandbox/internal/common/scoring"
	"github.com/cirobispo/sandbox/internal/common/turning"
)

type OnGameScore func(scoreA, scoreB int, done bool)

type GameScore struct {
	startSide      turning.TurningSide
	decidingPoint  bool
	scoreA, scoreB int

	onAfterScoreEvent []OnGameScore
}

func New(startSide turning.TurningSide, decidingPoint bool) GameScore {
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
	if g.startSide == turning.TSB {
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

func (g *GameScore) AddScore(score scoring.Scoring) error {
	if g.Done() { // am I acepting more points?
		return errors.New("Game completed already.")
	}

	if score.Type() != scoring.STPoint {
		return errors.New("This is not a score for a point.")
	}

	if !score.Done() {
		return errors.New("Point is not completed.")
	}

	incr := 1
	sA, sB := g.getScores()
	sideToAdd := sA
	if who := score.Side(); who == scoring.SSOpposite {
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

	return nil
}

func (g GameScore) Result() (int, int) {
	return g.scoreA, g.scoreB
}

func (g GameScore) Side() scoring.ScoringSide {
	return scoring.Side(g)
}

func (g GameScore) Type() scoring.ScoringType {
	return scoring.STGame
}

func (g GameScore) Done() bool {
	sA, sB := g.Result()
	diff := 2
	if g.decidingPoint {
		diff--
	}

	AWins := sA >= 4 && (sA-sB) >= diff
	BWins := sB >= 4 && (sB-sA) >= diff

	result := AWins || BWins

	return result
}
