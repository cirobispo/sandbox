package placarjogo

import (
	"errors"

	"github.com/cirobispo/sandbox/internal/comum/placares"
	"github.com/cirobispo/sandbox/internal/comum/turnos"
)

type AoMudarPlacar func(placarA, placarB int, terminado bool)
type TestaEncerramento func(valores ...int) bool

type Jogo struct {
	ladoInicio        turnos.LadoDoTurno
	pontoDecisivo     bool
	placarA, placarB  int
	testaEncerramento TestaEncerramento

	eventosAoMudarPlacar []AoMudarPlacar
}

func New(startSide turnos.LadoDoTurno, decidingPoint bool) Jogo {
	return Jogo{
		ladoInicio:           startSide,
		pontoDecisivo:        decidingPoint,
		placarA:              0,
		placarB:              0,
		eventosAoMudarPlacar: make([]AoMudarPlacar, 0),
	}
}

func terminado(valores ...int) bool {
	sA, sB, pontoDecisivo := valores[0], valores[1], valores[2] == 1
	diff := 2
	if pontoDecisivo {
		diff--
	}

	AWins := sA >= 4 && (sA-sB) >= diff
	BWins := sB >= 4 && (sB-sA) >= diff

	result := AWins || BWins

	return result
}

func (j *Jogo) placares() (*int, *int) {
	sA, sB := &j.placarA, &j.placarB
	if j.ladoInicio == turnos.LTB {
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

func (j Jogo) executarEventosAoMudarPlacar(scoreA, scoreB int) {
	done := j.Terminado()
	for i := range j.eventosAoMudarPlacar {
		event := j.eventosAoMudarPlacar[i]
		event(scoreA, scoreB, done)
	}
}

func (j Jogo) verificarEstado(placar placares.EstadoEParametroPlacar) error {
	if j.Terminado() {
		return errors.New("Jogo já terminado.")
	}

	if placar.Tipo() != placares.TPPonto {
		return errors.New("Este placar não é de ponto.")
	}

	if !placar.Terminado() {
		return errors.New("Este ponto não foi encerrado.")
	}

	return nil
}

func (j *Jogo) AdicionaPlacar(placar placares.EstadoEParametroPlacar) error {
	if err := j.verificarEstado(placar); err != nil {
		return err
	}

	incr := 1
	sA, sB := j.placares()
	sideToAdd := sA
	if who := placar.Lado(); who == placares.LPOposto {
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

func (j *Jogo) AdicionaAoMudarPlacar(event AoMudarPlacar) {
	j.eventosAoMudarPlacar = append(j.eventosAoMudarPlacar, event)
}

func (j *Jogo) DefinirTestaEncerramento(teste TestaEncerramento) {
	j.testaEncerramento = teste
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
	placarA, placaB := j.placarA, j.placarB
	pontoDecisivo := func(v bool) int {
		if v {
			return 1
		}
		return 0
	}(j.pontoDecisivo)

	if j.testaEncerramento != nil {
		return j.testaEncerramento(placarA, placaB, pontoDecisivo)
	}

	return terminado(placarA, placaB, pontoDecisivo)
}
