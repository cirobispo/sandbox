package game

import (
	"github.com/cirobispo/sandbox/internal/common/pointing/point"
	"github.com/cirobispo/sandbox/internal/common/scoring/gamescore"
	"github.com/cirobispo/sandbox/internal/common/turning"
	"github.com/cirobispo/sandbox/internal/common/turning/turn"
)

type Game struct {
	turn          *turn.Turn
	decidingPoint bool
	score         gamescore.GameScore
	points        []point.Point
}

func New(turn *turn.Turn, decidingPoint bool) Game {
	return Game{
		turn:          turn,
		decidingPoint: decidingPoint,
		score:         gamescore.New(turning.STA, decidingPoint),
	}
}

func (g *Game) AddPoint(p point.Point) {
	g.points = append(g.points, p)
	g.score.AddPoint(p)
	g.turn.Execute()
}

func (g Game) Result() (int, int) {
	return g.score.Result()
}

func (g Game) NewTurn() *turn.Turn {
	return g.turn.Clone(g.turn.LastSide())
}
