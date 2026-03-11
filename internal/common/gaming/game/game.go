package game

import (
	"github.com/cirobispo/sandbox/internal/common/gaming"
	"github.com/cirobispo/sandbox/internal/common/pointing/point"
	"github.com/cirobispo/sandbox/internal/common/scoring/gamescore"
	"github.com/cirobispo/sandbox/internal/common/turning"
	"github.com/cirobispo/sandbox/internal/common/turning/turn"
)

type Game struct {
	turn               *turn.Turn
	decidingPoint      bool
	score              gamescore.GameScore
	points             []point.Point
	onAddingPointEvent []gaming.OnAfterAddingPoint
}

func New(turn *turn.Turn, decidingPoint bool) Game {
	return Game{
		turn:               turn,
		decidingPoint:      decidingPoint,
		score:              gamescore.New(turning.STA, decidingPoint),
		onAddingPointEvent: make([]gaming.OnAfterAddingPoint, 0),
	}
}

func (g *Game) AddOnAddingPointEvent(event gaming.OnAfterAddingPoint) {
	g.onAddingPointEvent = append(g.onAddingPointEvent, event)
}

func (g *Game) AddPoint(p point.Point) {
	g.points = append(g.points, p.Clone())
	g.score.AddScore(gamescore.PointToScore(&p, g.decidingPoint))
	g.turn.Execute()

	scoreA, scoreB := g.score.Result()
	done := g.score.Done()
	g.executeOnAfterAddingPoint(scoreA, scoreB, done)
}

func (g Game) Score() gamescore.GameScore {
	return g.score
}

func (g Game) NewTurn() *turn.Turn {
	return g.turn.Clone(g.turn.LastSide())
}

func (g Game) executeOnAfterAddingPoint(scoreA, scoreB int, done bool) {
	for j := range g.onAddingPointEvent {
		event := g.onAddingPointEvent[j]
		event(scoreA, scoreB, done)
	}
}
