package game

import (
	"github.com/cirobispo/sandbox/internal/common/pointing/point"
	"github.com/cirobispo/sandbox/internal/common/scoring/gamescore"
	"github.com/cirobispo/sandbox/internal/common/turning"
)

type Game struct {
	score gamescore.GameScore
}

func New(startSide turning.SideTurn, decidingPoint bool) Game {
	return Game{
		score: gamescore.New(startSide, decidingPoint),
	}
}

func (g *Game) AddPoint(p point.Point) {
	g.score.AddPoint(p)
}
