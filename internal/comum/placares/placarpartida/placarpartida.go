package placarpartida

import (
	"errors"

	"github.com/cirobispo/sandbox/internal/comum/placares"
	"github.com/cirobispo/sandbox/internal/comum/turnos"
)

type AoPontuarNaPartida func(placarA, placarB int, terminado bool)
type ParamOption func(p *Partida)

type Partida struct {
	ladoInicial      turnos.Lado
	melhorDe         int
	placarA, placarB int

	eventosAoPontuarNaPartida []AoPontuarNaPartida
}

func MelhorDeTres() ParamOption {
	return TamanhoELado(turnos.LadoA, 3)
}

func MelhorDeCinco() ParamOption {
	return TamanhoELado(turnos.LadoA, 5)
}

func TamanhoELado(ladoIncial turnos.Lado, melhorDe int) ParamOption {
	return func(score *Partida) {
		score.ladoInicial = ladoIncial
		score.melhorDe = melhorDe
	}
}

func New(param ParamOption) Partida {
	result := Partida{
		placarA:                   0,
		placarB:                   0,
		eventosAoPontuarNaPartida: make([]AoPontuarNaPartida, 0),
	}

	if param != nil {
		param(&result)
	}

	return result
}

func (p Partida) executeEventosAoPontuarNaPartida(placarA, placarB int) {
	done := p.Terminado()
	for i := range p.eventosAoPontuarNaPartida {
		event := p.eventosAoPontuarNaPartida[i]
		event(placarA, placarB, done)
	}
}

func (p *Partida) placares() (*int, *int) {
	sA, sB := &p.placarA, &p.placarB
	if p.ladoInicial == turnos.LadoB {
		sA, sB = &p.placarB, &p.placarA
	}

	return sA, sB
}

func (p *Partida) AdicionarEventoAoPontuarNaPartida(event AoPontuarNaPartida) {
	p.eventosAoPontuarNaPartida = append(p.eventosAoPontuarNaPartida, event)
}

func (p Partida) verificarEstado(placar placares.EstadoEParametroPlacar) error {
	if p.Terminado() { // am I acepting more points?
		return errors.New("Match completed already.")
	}

	if placar.Tipo() != placares.TPSet {
		return errors.New("This is not Set Score.")
	}

	if !placar.Terminado() {
		return errors.New("Set is not completed.")
	}

	return nil
}

func (p *Partida) AdicionarPlacar(placar placares.EstadoEParametroPlacar) error {
	if err := p.verificarEstado(placar); err != nil {
		return err
	}

	sA, sB := p.placares()
	sideToAdd := sA

	if placar.Lado() == placares.LPOposto {
		sideToAdd = sB
	}

	*sideToAdd += 1
	p.executeEventosAoPontuarNaPartida(p.placarA, p.placarB)
	return nil
}

func (p Partida) Resultado() (int, int) {
	return p.placarA, p.placarB
}

func (p Partida) Terminado() bool {
	sA, sB := p.Resultado()

	amountToWin := (p.melhorDe / 2) + (p.melhorDe % 2)

	sideAWins := (sA >= amountToWin && sA-sB >= 1)
	sideBWins := (sB >= amountToWin && sB-sA >= 1)
	result := sideAWins || sideBWins

	return result
}

func (p Partida) Lado() placares.LadoDoPlacar {
	return placares.Lado(p)
}

func (p Partida) Tipo() placares.TipoDoPlacar {
	return placares.TPPartida
}
