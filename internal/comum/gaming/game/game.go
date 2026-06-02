package game

import (
	"errors"

	"github.com/cirobispo/sandbox/internal/comum/gaming"
	"github.com/cirobispo/sandbox/internal/comum/placares"
	"github.com/cirobispo/sandbox/internal/comum/placares/placarjogo"
	"github.com/cirobispo/sandbox/internal/comum/pontos/ponto"
	"github.com/cirobispo/sandbox/internal/comum/turnos"
	"github.com/cirobispo/sandbox/internal/comum/turnos/turno"
)

type Gaming interface {
	ServingSide() turnos.LadoDoTurno
	Score() placares.EstadoResultadoEParametroPlacar
	Points() []ponto.Ponto
}

type Game struct {
	turn               *turno.Turno
	decidingPoint      bool
	score              placarjogo.Jogo
	points             []ponto.Ponto
	onAddingPointEvent []gaming.OnAfterAddingPoint
}

func New(turn *turno.Turno, decidingPoint bool) *Game {
	side := turn.LadoInicial()
	return &Game{
		turn:               turn,
		decidingPoint:      decidingPoint,
		score:              placarjogo.New(side, decidingPoint),
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

func (g *Game) AddPoint(p ponto.Ponto) error {
	if !p.Terminado() {
		return errors.New("point is still in play.")
	}

	g.points = append(g.points, p.Clonar())
	scoreToAdd, error := placarjogo.PontoParaPlacar(&p)

	if error != nil {
		return errors.New("point is still in play.")
	}

	g.score.AdicionaPlacar(scoreToAdd)
	g.turn.Execute()

	scoreA, scoreB := g.score.Resultado()
	done := g.score.Terminado()
	g.executeOnAfterAddingPoint(scoreA, scoreB, done)
	return nil
}

func (g Game) ServingSide() turnos.LadoDoTurno {
	return g.turn.LadoInicial()
}

func (g Game) Score() placares.EstadoEResultadoPlacar {
	return g.score
}

func (g Game) Points() []ponto.Ponto {
	result := make([]ponto.Ponto, 0, len(g.points))
	copy(result, g.points)

	return result
}
