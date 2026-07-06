package set

import (
	"errors"

	"github.com/cirobispo/sandbox/internal/comum/jogos"
	"github.com/cirobispo/sandbox/internal/comum/jogos/jogo"
	"github.com/cirobispo/sandbox/internal/comum/placares"
	"github.com/cirobispo/sandbox/internal/comum/placares/placarset"
	"github.com/cirobispo/sandbox/internal/comum/turnos"
	"github.com/cirobispo/sandbox/internal/comum/turnos/turno"
)

type ParamOption func(set *Set) bool

type Set struct {
	quemServe, ladoServico    *turno.Turno
	jogosPorSet               int
	pontoDecisivo             bool
	tieBreak                  bool
	placar                    placarset.Set
	jogos                     []jogo.Gaming
	EventosAoAdicionarJogo    []jogos.OnAfterAddingGame
	EventosAoMudarLadoJogador []jogos.OnPlayerChangeSide
}

func SetPadrao(turnForServing *turno.Turno) ParamOption {
	return func(score *Set) bool {
		score.quemServe = turnForServing
		score.ladoServico = turno.New(turno.DefinindoLado(turnForServing.LadoInicial()))
		score.jogosPorSet = 6
		score.pontoDecisivo = false
		score.tieBreak = true
		return false
	}
}

func TurnoJogosETieBreak(turnForServing *turno.Turno, size int, decidingPoint, tieBreak bool) ParamOption {
	return func(score *Set) bool {
		score.quemServe = turnForServing
		score.ladoServico = turno.New(turno.DefinindoLado(turnForServing.LadoInicial()))
		score.jogosPorSet = size
		score.pontoDecisivo = decidingPoint
		score.tieBreak = tieBreak
		return true
	}
}

func New(param ParamOption) *Set {
	result := &Set{
		jogos:                     make([]jogo.Gaming, 0, 13),
		EventosAoAdicionarJogo:    make([]jogos.OnAfterAddingGame, 0),
		EventosAoMudarLadoJogador: make([]jogos.OnPlayerChangeSide, 0),
	}

	if param != nil {
		custom := param(result)

		serving_side := result.quemServe.LadoInicial()
		callback := placarset.SetPadrao(serving_side)
		if custom {
			callback = placarset.TamanhoVantagemETieBreak(serving_side, result.jogosPorSet, result.pontoDecisivo, result.tieBreak)
		}
		result.placar = placarset.New(callback)
	}

	result.ladoServico.AdicionarDepoisDeMudarTurno(func(ts turnos.LadoDoTurno) {
		result.executarAoMudarLadoJogador()
	})

	return result
}

func (s Set) executarAoAdicionarJogo() {
	if len(s.EventosAoAdicionarJogo) > 0 {
		placarA, placarB := s.placar.Resultado()
		terminado := s.placar.Terminado()

		for j := range s.EventosAoAdicionarJogo {
			event := s.EventosAoAdicionarJogo[j]
			event(placarA, placarB, terminado)
		}
	}
}

func (s Set) executarAoMudarLadoJogador() {
	for j := range s.EventosAoMudarLadoJogador {
		event := s.EventosAoMudarLadoJogador[j]
		event()
	}
}

func (s *Set) AdicionarAoAdicionarJogo(event jogos.OnAfterAddingGame) {
	s.EventosAoAdicionarJogo = append(s.EventosAoAdicionarJogo, event)
}

func (s *Set) AdicionarAoMudarLadoJogador(event jogos.OnPlayerChangeSide) {
	s.EventosAoMudarLadoJogador = append(s.EventosAoMudarLadoJogador, event)
}

func (s Set) verificarEstado(j jogo.Gaming) error {
	if s.placar.Terminado() {
		return errors.New("O set está encerrado.")
	}

	if !j.Score().Terminado() {
		return errors.New("O jogo está em andamento.")
	}

	return nil
}

func (s *Set) AdicionarJogo(j jogo.Gaming) error {
	if err := s.verificarEstado(j); err != nil {
		return err
	}

	s.jogos = append(s.jogos, j)
	s.quemServe.Execute()
	s.placar.AdicionarPlacar(j.Score())

	s.executarAoAdicionarJogo()
	if q := len(s.jogos) % 2; q == 1 {
		s.ladoServico.Execute()
	}

	return nil
}

func (s Set) NovoJogo() *jogo.Jogo {
	newTurn := s.quemServe.Clonar(s.quemServe.LadoCorrente())
	result := jogo.New(jogo.Regular(newTurn, s.pontoDecisivo))
	if s.placar.IsTieBreak() {
		result = jogo.New(jogo.TieBreak(newTurn, s.pontoDecisivo))
	}
	return result
}

func (s Set) Placar() placares.EstadoEResultadoPlacar {
	return s.placar
}
