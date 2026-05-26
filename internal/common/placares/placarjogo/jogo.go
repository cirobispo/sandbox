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

	eventosAoMudarPlacar []AoMudarPlacar
}

func New(startSide turning.TurningSide, decidingPoint bool) Jogo {
	return Jogo{
		ladoInicio:           startSide,
		pontoDecisivo:        decidingPoint,
		placarA:              0,
		placarB:              0,
		eventosAoMudarPlacar: make([]AoMudarPlacar, 0),
	}
}

func (j *Jogo) placares() (*int, *int) {
	sA, sB := &j.placarA, &j.placarB
	if j.ladoInicio == turning.TSB {
		sA, sB = &j.placarB, &j.placarA
	}

	return sA, sB
}

func (j *Jogo) placarInvertido(side *int) *int {
	sA, sB := j.placares()
	if side == sA {
		return sB
	}
	return sA
}

func (j *Jogo) AdicionaAoMudarPlacar(event AoMudarPlacar) {
	j.eventosAoMudarPlacar = append(j.eventosAoMudarPlacar, event)
}

func (j Jogo) executarEventosAoMudarPlacar(scoreA, scoreB int) {
	done := j.Terminado()
	for i := range j.eventosAoMudarPlacar {
		event := j.eventosAoMudarPlacar[i]
		event(scoreA, scoreB, done)
	}
}

func (j Jogo) verificarEstado(placar placares.EstadoEParametroPlacar) error {
	if j.Terminado() { // am I acepting more points?
		return errors.New("Game completed already.")
	}

	if placar.Tipo() != placares.TPPonto {
		return errors.New("This is not a score for a point.")
	}

	if !placar.Terminado() {
		return errors.New("Point is not completed.")
	}

	return nil
}

func (j *Jogo) AdicionaPlacar(score placares.EstadoEParametroPlacar) error {
	if err := j.verificarEstado(score); err != nil {
		return err
	}

	incr := 1
	sA, sB := j.placares()
	sideToAdd := sA
	if who := score.Lado(); who == placares.LPOposto {
		sideToAdd = sB
	}

	if (!j.pontoDecisivo) && (*sA > 3 || *sB > 3) {
		if *sideToAdd == 3 {
			incr = -1
			sideToAdd = j.placarInvertido(sideToAdd)
		}
	}

	*sideToAdd += incr
	j.executarEventosAoMudarPlacar(j.Resultado())

	return nil
}

func (j Jogo) Resultado() (int, int) {
	return j.placarA, j.placarB
}

func (j Jogo) Lado() placares.LadoDoPlacar {
	return placares.Lado(j)
}

func (j Jogo) Tipo() placares.TipoDoPlacar {
	return placares.TPJogo
}

func (j Jogo) Terminado() bool {
	sA, sB := j.Resultado()
	diff := 2
	if j.pontoDecisivo {
		diff--
	}

	AWins := sA >= 4 && (sA-sB) >= diff
	BWins := sB >= 4 && (sB-sA) >= diff

	result := AWins || BWins

	return result
}
