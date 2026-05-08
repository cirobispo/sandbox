package placarjogo

import (
	"errors"

	"github.com/cirobispo/sandbox/internal/common/placares"
	"github.com/cirobispo/sandbox/internal/common/turning"
)

type AoMudarPlacar func(placarA, placarB int, terminado bool)

type Jogo struct {
	ladoInicio       turning.TurningSide
	pontoDecisivo    bool
	placarA, placarB int

	EventosAoMudarPlacar []AoMudarPlacar
}

func New(startSide turning.TurningSide, decidingPoint bool) Jogo {
	return Jogo{
		ladoInicio:           startSide,
		pontoDecisivo:        decidingPoint,
		placarA:              0,
		placarB:              0,
		EventosAoMudarPlacar: make([]AoMudarPlacar, 0),
	}
}

func (g *Jogo) placares() (*int, *int) {
	sA, sB := &g.placarA, &g.placarB
	if g.ladoInicio == turning.TSB {
		sA, sB = &g.placarB, &g.placarA
	}

	return sA, sB
}

func (g *Jogo) placarInvertido(side *int) *int {
	sA, sB := g.placares()
	if side == sA {
		return sB
	}
	return sA
}

func (g *Jogo) AdicionaAoMudarPlacar(event AoMudarPlacar) {
	g.EventosAoMudarPlacar = append(g.EventosAoMudarPlacar, event)
}

func (g Jogo) executarEventAoMudarPlacar(scoreA, scoreB int) {
	done := g.Terminado()
	for i := range g.EventosAoMudarPlacar {
		event := g.EventosAoMudarPlacar[i]
		event(scoreA, scoreB, done)
	}
}

func (g *Jogo) AdicionaPlacar(score placares.EstadoEParametroPlacar) error {
	if g.Terminado() { // am I acepting more points?
		return errors.New("Game completed already.")
	}

	if score.Tipo() != placares.TPPonto {
		return errors.New("This is not a score for a point.")
	}

	if !score.Terminado() {
		return errors.New("Point is not completed.")
	}

	incr := 1
	sA, sB := g.placares()
	sideToAdd := sA
	if who := score.Lado(); who == placares.LPOposto {
		sideToAdd = sB
	}

	if (!g.pontoDecisivo) && (*sA > 3 || *sB > 3) {
		if *sideToAdd == 3 {
			incr = -1
			sideToAdd = g.placarInvertido(sideToAdd)
		}
	}

	*sideToAdd += incr
	g.executarEventAoMudarPlacar(g.placarA, g.placarB)

	return nil
}

func (g Jogo) Resultado() (int, int) {
	return g.placarA, g.placarB
}

func (g Jogo) Lado() placares.LadoDoPlacar {
	return placares.Lado(g)
}

func (g Jogo) Tipo() placares.TipoDoPlacar {
	return placares.TPJogo
}

func (g Jogo) Terminado() bool {
	sA, sB := g.Resultado()
	diff := 2
	if g.pontoDecisivo {
		diff--
	}

	AWins := sA >= 4 && (sA-sB) >= diff
	BWins := sB >= 4 && (sB-sA) >= diff

	result := AWins || BWins

	return result
}
