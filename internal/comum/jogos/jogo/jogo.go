package jogo

import (
	"errors"

	"github.com/cirobispo/sandbox/internal/comum/jogos"
	"github.com/cirobispo/sandbox/internal/comum/placares"
	"github.com/cirobispo/sandbox/internal/comum/placares/placarjogo"
	"github.com/cirobispo/sandbox/internal/comum/placares/placarponto"
	"github.com/cirobispo/sandbox/internal/comum/pontos/ponto"
	"github.com/cirobispo/sandbox/internal/comum/turnos"
	"github.com/cirobispo/sandbox/internal/comum/turnos/turno"
)

type Gaming interface {
	ServingSide() turnos.LadoDoTurno
	Score() placares.EstadoResultadoEParametroPlacar
	Points() []ponto.Ponto
}

type Jogo struct {
	turno                   *turno.Turno
	pontoDecisivo           bool
	placar                  placarjogo.Jogo
	pontos                  []ponto.Ponto
	eventosAoAdicionarPonto []jogos.AoAdicionarPonto
}

func New(turn *turno.Turno, decidingPoint bool) *Jogo {
	side := turn.LadoInicial()
	return &Jogo{
		turno:                   turn,
		pontoDecisivo:           decidingPoint,
		placar:                  placarjogo.New(side, decidingPoint),
		eventosAoAdicionarPonto: make([]jogos.AoAdicionarPonto, 0),
	}
}

func (j Jogo) executeEventosAoAdicionarPonto(scoreA, scoreB int, done bool) {
	for e := range j.eventosAoAdicionarPonto {
		event := j.eventosAoAdicionarPonto[e]
		event(scoreA, scoreB, done)
	}
}

func (j *Jogo) AdicionarEventoAoAdicionarPonto(event jogos.AoAdicionarPonto) {
	j.eventosAoAdicionarPonto = append(j.eventosAoAdicionarPonto, event)
}

func (j *Jogo) AdicionarPonto(p ponto.Ponto) error {
	if !p.Terminado() {
		return errors.New("O ponto ainda está em andamento.")
	}

	j.pontos = append(j.pontos, p.Clonar())
	placar := placarponto.New(&p)

	j.placar.AdicionaPlacar(placar)
	j.turno.Execute()

	placarA, placarB := j.placar.Resultado()
	terminado := j.placar.Terminado()
	j.executeEventosAoAdicionarPonto(placarA, placarB, terminado)
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
