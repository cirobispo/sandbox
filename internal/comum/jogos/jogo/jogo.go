package jogo

import (
	"errors"

	"github.com/cirobispo/sandbox/internal/comum/jogos"
	"github.com/cirobispo/sandbox/internal/comum/placares"
	"github.com/cirobispo/sandbox/internal/comum/placares/placarjogo"
	"github.com/cirobispo/sandbox/internal/comum/placares/placarponto"
	"github.com/cirobispo/sandbox/internal/comum/placares/placartiebreak"
	"github.com/cirobispo/sandbox/internal/comum/pontos/ponto"
	"github.com/cirobispo/sandbox/internal/comum/turnos"
	"github.com/cirobispo/sandbox/internal/comum/turnos/turno"
)

type Gaming interface {
	ServingSide() turnos.LadoDoTurno
	Score() placares.EstadoResultadoEParametroPlacar
	Points() []ponto.Ponto
}

type ParamOption func(jogo *Jogo)

type Jogo struct {
	turno                   *turno.Turno
	pontoDecisivo           bool
	placar                  placares.EstadoResultadoParametroEAdicionadorPlacar
	pontos                  []ponto.Ponto
	eventosAoAdicionarPonto []jogos.AoAdicionarPonto
}

func Regular(turno *turno.Turno, pontoDecisivo bool) ParamOption {
	ladoInicial := turno.LadoInicial()
	return func(jogo *Jogo) {
		jogo.turno = turno
		jogo.pontoDecisivo = pontoDecisivo
		jogo.placar = placarjogo.New(ladoInicial, pontoDecisivo)
	}
}

func TieBreak(turno *turno.Turno, pontoDecisivo bool) ParamOption {
	ladoInicial := turno.LadoInicial()
	return func(jogo *Jogo) {
		jogo.turno = turno
		jogo.pontoDecisivo = pontoDecisivo
		jogo.placar = placartiebreak.New(placartiebreak.ChegarEm7(ladoInicial, pontoDecisivo))
	}
}

func SuperTieBreak(turno *turno.Turno, pontoDecisivo bool) ParamOption {
	ladoInicial := turno.LadoInicial()
	return func(jogo *Jogo) {
		jogo.turno = turno
		jogo.pontoDecisivo = pontoDecisivo
		jogo.placar = placartiebreak.New(placartiebreak.ChegarEm10(ladoInicial, pontoDecisivo))
	}
}

func New(param ParamOption) *Jogo {
	result := &Jogo{
		eventosAoAdicionarPonto: make([]jogos.AoAdicionarPonto, 0),
	}

	param(result)

	return result
}

func (j Jogo) executeEventosAoAdicionarPonto() {
	if len(j.eventosAoAdicionarPonto) > 0 {
		placarA, placarB := j.placar.Resultado()
		terminado := j.placar.Terminado()

		for e := range j.eventosAoAdicionarPonto {
			event := j.eventosAoAdicionarPonto[e]
			event(placarA, placarB, terminado)
		}
	}
}

func (j *Jogo) AdicionarEventoAoAdicionarPonto(evento jogos.AoAdicionarPonto) {
	j.eventosAoAdicionarPonto = append(j.eventosAoAdicionarPonto, evento)
}

func (j *Jogo) AdicionarPonto(p ponto.Ponto) error {
	if !p.Terminado() {
		return errors.New("O ponto ainda está em andamento.")
	}

	j.pontos = append(j.pontos, p.Clonar())
	placar := placarponto.New(&p)

	j.placar.AdicionaPlacar(placar)
	j.turno.Execute()

	j.executeEventosAoAdicionarPonto()
	return nil
}

func (j Jogo) LadoDoServico() turnos.LadoDoTurno {
	return j.turno.LadoInicial()
}

func (j Jogo) Placar() placares.EstadoEResultadoPlacar {
	return j.placar
}

func (j Jogo) Pontos() []ponto.Ponto {
	result := make([]ponto.Ponto, len(j.pontos))
	copy(result, j.pontos)

	return result
}
