package gamescore

import (
	"github.com/cirobispo/sandbox/internal/common/pointing"
	"github.com/cirobispo/sandbox/internal/common/pointing/point"
	"github.com/cirobispo/sandbox/internal/common/turning"
)

type OnGameScore func(scoreA, scoreB int, done bool)

type GameScore struct {
	side           turning.SideTurn
	decidingPoint  bool
	scoreA, scoreB int
	points         []point.Point

	onAfterScore []OnGameScore
}

func New(startSide turning.SideTurn, decidingPoint bool) GameScore {
	return GameScore{
		side:          startSide,
		decidingPoint: decidingPoint,
		scoreA:        0,
		scoreB:        0,
		onAfterScore:  make([]OnGameScore, 0),
	}
}

func (gs *GameScore) getPointSides() (*int, *int) {
	sA, sB := &gs.scoreA, &gs.scoreB
	if gs.side == turning.STB {
		sA, sB = &gs.scoreB, &gs.scoreA
	}

	return sA, sB
}

func (gs GameScore) executeOnAfterScoreEvent(scoreA, scoreB int) {
	closeEnding := (scoreA >= 3 || scoreB >= 3) && scoreA != scoreB
	done := (scoreA >= 3 || scoreB >= 3) && closeEnding
	for i := range gs.onAfterScore {
		event := gs.onAfterScore[i]
		event(scoreA, scoreB, done)
	}
}

func (gs *GameScore) AddPoint(p point.Point) {
	gs.points = append(gs.points, p)
	if who := p.Side(); who != pointing.PSNone {
		sA, sB := gs.getPointSides()
		(*sA)++
		if who == pointing.PSB {
			(*sB)++
			(*sA)--
		}

		gs.executeOnAfterScoreEvent(gs.scoreA, gs.scoreB)
	}
}
