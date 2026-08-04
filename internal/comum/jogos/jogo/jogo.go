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
	ServingSide() turnos.Lado
	Score() placares.EstadoResultadoEParametroPlacar
	Points() []ponto.Ponto
}

type ParamOption func(jogo *Jogo)

type Jogo struct {
	ladoInicial             turnos.Lado
	qualSacador             *turno.Turno
	pontoDecisivo           bool
	placar                  placares.EstadoResultadoParametroEAdicionadorPlacar
	pontos                  []ponto.Ponto
	eventosAoAdicionarPonto []jogos.AoAdicionarPonto
}

func Regular(ladoInicial turnos.Lado, pontoDecisivo bool) ParamOption {
	return func(jogo *Jogo) {
		jogo.ladoInicial = ladoInicial
		jogo.qualSacador = turno.New(turno.DefinindoLado(ladoInicial))
		jogo.pontoDecisivo = pontoDecisivo
		jogo.placar = placarjogo.New(ladoInicial, pontoDecisivo)
	}
}

func TieBreak(ladoInicial turnos.Lado, pontoDecisivo bool) ParamOption {
	return func(jogo *Jogo) {
		jogo.ladoInicial = ladoInicial
		jogo.qualSacador = turno.New(turno.DefinindoLado(ladoInicial))
		jogo.pontoDecisivo = pontoDecisivo
		placarTieBreak := placartiebreak.New(placartiebreak.ChegarEm7(ladoInicial, pontoDecisivo))
		placarTieBreak.AdicionaAoMudarPlacar(func(placarA, placarB int, terminado bool) {
			total := placarA + placarB
			if total%2 == 1 {
				jogo.qualSacador.Execute()
			}
		})
		jogo.placar = placarTieBreak
	}
}

func SuperTieBreak(ladoInicial turnos.Lado, pontoDecisivo bool) ParamOption {
	return func(jogo *Jogo) {
		jogo.ladoInicial = ladoInicial
		jogo.qualSacador = turno.New(turno.DefinindoLado(ladoInicial))
		jogo.pontoDecisivo = pontoDecisivo
		placarJogo := placartiebreak.New(placartiebreak.ChegarEm10(ladoInicial, pontoDecisivo))
		placarJogo.AdicionaAoMudarPlacar(func(placarA, placarB int, terminado bool) {
			total := placarA + placarB
			if total%2 == 1 {
				jogo.qualSacador.Execute()
			}
		})
		jogo.placar = placarJogo
	}
}

func New(param ParamOption) *Jogo {
	result := &Jogo{eventosAoAdicionarPonto: make([]jogos.AoAdicionarPonto, 0)}
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

func (j Jogo) verificarEstado(p *ponto.Ponto) error {
	if !p.Terminado() {
		return errors.New("O ponto ainda está em andamento.")
	}

	if j.placar.Terminado() {
		return errors.New("O jogo já foi encerrado.")
	}

	return nil
}

func (j *Jogo) AdicionarPonto(p *ponto.Ponto) error {
	if err := j.verificarEstado(p); err != nil {
		return err
	}

	j.pontos = append(j.pontos, p.Clonar())
	j.placar.AdicionarPlacar(placarponto.New(p, j.ladoInicial, j.qualSacador.LadoCorrente()))

	j.executeEventosAoAdicionarPonto()
	return nil
}

func (j Jogo) LadoDoServico() turnos.Lado {
	return j.ladoInicial
}

func (j Jogo) Placar() placares.EstadoEResultadoPlacar {
	return j.placar
}

func (j Jogo) Pontos() []ponto.Ponto {
	result := make([]ponto.Ponto, len(j.pontos))
	copy(result, j.pontos)

	return result
}

func (j Jogo) QualSacador() turnos.Lado {
	return j.qualSacador.LadoCorrente()
}
