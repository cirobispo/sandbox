package placarjogo

import (
	"errors"

	"github.com/cirobispo/sandbox/internal/comum/placares"
	"github.com/cirobispo/sandbox/internal/comum/turnos"
)

type AoMudarPlacar func(placarA, placarB int, terminado bool)

type Jogo struct {
	ladoInicial      turnos.Lado
	pontoDecisivo    bool
	placarA, placarB int

	eventosAoMudarPlacar []AoMudarPlacar
}

func New(ladoInicial turnos.Lado, pontoDecisivo bool) *Jogo {
	return &Jogo{
		ladoInicial:          ladoInicial,
		pontoDecisivo:        pontoDecisivo,
		placarA:              0,
		placarB:              0,
		eventosAoMudarPlacar: make([]AoMudarPlacar, 0),
	}
}

func (j *Jogo) placares() (*int, *int) {
	sA, sB := &j.placarA, &j.placarB
	if j.ladoInicial == turnos.LadoB {
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

func (j Jogo) executarEventosAoMudarPlacar() {
	if len(j.eventosAoMudarPlacar) > 0 {
		placarA, placarB := j.Resultado()
		terminado := j.Terminado()
		for i := range j.eventosAoMudarPlacar {
			event := j.eventosAoMudarPlacar[i]
			event(placarA, placarB, terminado)
		}
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

func (j *Jogo) AdicionarPlacar(placarPonto placares.EstadoEParametroPlacar) error {
	if err := j.verificarEstado(placarPonto); err != nil {
		return err
	}

	incr := 1
	sA, sB := j.placares()
	adicionarEm := sA
	if who := placarPonto.Lado(); who == placares.LPOposto {
		adicionarEm = sB
	}

	if (!j.pontoDecisivo) && (*sA > 3 || *sB > 3) {
		if *adicionarEm == 3 {
			incr = -1
			adicionarEm = j.placarInvertido(adicionarEm)
		}
	}

	*adicionarEm += incr
	j.executarEventosAoMudarPlacar()

	return nil
}

func (j *Jogo) AdicionaAoMudarPlacar(event AoMudarPlacar) {
	j.eventosAoMudarPlacar = append(j.eventosAoMudarPlacar, event)
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
	sA, sB := j.placarA, j.placarB
	diff := 2
	if j.pontoDecisivo {
		diff--
	}

	AWins := sA >= 4 && (sA-sB) >= diff
	BWins := sB >= 4 && (sB-sA) >= diff

	result := AWins || BWins

	return result
}
