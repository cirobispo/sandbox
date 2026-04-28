package game

import (
	"errors"

	"github.com/cirobispo/sandbox/internal/common/gaming"
	"github.com/cirobispo/sandbox/internal/common/pointing/point"
	"github.com/cirobispo/sandbox/internal/common/scoring"
	"github.com/cirobispo/sandbox/internal/common/scoring/gamescore"
	"github.com/cirobispo/sandbox/internal/common/turning"
	"github.com/cirobispo/sandbox/internal/common/turning/turn"
)

type Gaming interface {
	ServingSide() turning.TurningSide
	Score() scoring.EstadoResultadoEParametroPlacar
	Points() []point.Point
}

type Game struct {
	turn               *turn.Turn
	decidingPoint      bool
	score              gamescore.GameScore
	points             []point.Point
	onAddingPointEvent []gaming.OnAfterAddingPoint
}

func New(turn *turn.Turn, decidingPoint bool) *Game {
	side := turn.StartSide()
	return &Game{
		turn:               turn,
		decidingPoint:      decidingPoint,
		score:              gamescore.New(side, decidingPoint),
		onAddingPointEvent: make([]gaming.OnAfterAddingPoint, 0),
	}
}

func (g Game) executeOnAfterAddingPoint(scoreA, scoreB int, done bool) {
	for j := range g.onAddingPointEvent {
		event := g.onAddingPointEvent[j]
		event(scoreA, scoreB, done)
	}
}

func (g *Game) AddOnAddingPointEvent(event gaming.OnAfterAddingPoint) {
	g.onAddingPointEvent = append(g.onAddingPointEvent, event)
}

func (g *Game) AddPoint(p point.Point) error {
	if !p.Done() {
		return errors.New("point is still in play.")
	}

	g.points = append(g.points, p.Clone())
	scoreToAdd, error := gamescore.PointToScore(&p)

	if error != nil {
		return errors.New("point is still in play.")
	}

	g.score.AddScore(scoreToAdd)
	g.turn.Execute()

	scoreA, scoreB := g.score.Resultado()
	done := g.score.Terminado()
	g.executeOnAfterAddingPoint(scoreA, scoreB, done)
	return nil
}

func (g Game) ServingSide() turning.TurningSide {
	return g.turn.StartSide()
}

func (g Game) Score() scoring.EstadoEResultadoPlacar {
	return g.score
}

func (g Game) Points() []point.Point {
	result := make([]point.Point, 0, len(g.points))
	copy(result, g.points)

	return result
}
