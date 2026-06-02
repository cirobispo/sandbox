package set

import (
	"errors"

	"github.com/cirobispo/sandbox/internal/comum/gaming"
	"github.com/cirobispo/sandbox/internal/comum/gaming/game"
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
	jogos                     []game.Gaming
	EventosAoAdicionarJogo    []gaming.OnAfterAddingGame
	EventosAoMudarLadoJogador []gaming.OnPlayerChangeSide
}

func SetPadrao(turnForServing *turno.Turno) ParamOption {
	return func(score *Set) bool {
		score.quemServe = turnForServing
		score.ladoServico = turno.New(turno.MudandoLado(turnForServing.LadoInicial()))
		score.jogosPorSet = 6
		score.pontoDecisivo = false
		score.tieBreak = true
		return false
	}
}

func TurnoJogosETieBreak(turnForServing *turno.Turno, size int, decidingPoint, tieBreak bool) ParamOption {
	return func(score *Set) bool {
		score.quemServe = turnForServing
		score.ladoServico = turno.New(turno.MudandoLado(turnForServing.LadoInicial()))
		score.jogosPorSet = size
		score.pontoDecisivo = decidingPoint
		score.tieBreak = tieBreak
		return true
	}
}

func New(param ParamOption) *Set {
	result := &Set{
		jogos:                     make([]game.Gaming, 0, 13),
		EventosAoAdicionarJogo:    make([]gaming.OnAfterAddingGame, 0),
		EventosAoMudarLadoJogador: make([]gaming.OnPlayerChangeSide, 0),
	}

	if param != nil {
		custom := param(result)

		serving_side := result.quemServe.LadoInicial()
		callback := placarset.SetPadrao(serving_side)
		if custom {
			callback = placarset.TamanhoETieBreak(serving_side, result.jogosPorSet, result.pontoDecisivo, result.tieBreak)
		}
		result.placar = placarset.New(callback)
	}

	result.ladoServico.AdicionarDepoisDeMudarTurno(func(ts turnos.LadoDoTurno) {
		result.executarAoMudarLadoJogador()
	})

	return result
}

func (s Set) executarAoAdicionarJogo(scoreA, scoreB int, done bool) {
	for j := range s.EventosAoAdicionarJogo {
		event := s.EventosAoAdicionarJogo[j]
		event(scoreA, scoreB, done)
	}
}

func (s Set) executarAoMudarLadoJogador() {
	for j := range s.EventosAoMudarLadoJogador {
		event := s.EventosAoMudarLadoJogador[j]
		event()
	}
}

func (s *Set) AdicionarAoAdicionarJogo(event gaming.OnAfterAddingGame) {
	s.EventosAoAdicionarJogo = append(s.EventosAoAdicionarJogo, event)
}

func (s *Set) AdicionarAoMudarLadoJogador(event gaming.OnPlayerChangeSide) {
	s.EventosAoMudarLadoJogador = append(s.EventosAoMudarLadoJogador, event)
}

func (s *Set) AdicionarJogo(g game.Gaming) error {
	if s.placar.Terminado() {
		return errors.New("set is closed.")
	}

	if !g.Score().Terminado() {
		return errors.New("game is still in play.")
	}

	s.jogos = append(s.jogos, g)
	s.quemServe.Execute()
	s.placar.AdicionarPlacar(g.Score())

	scoreA, scoreB := s.placar.Resultado()
	done := s.placar.Terminado()
	s.executarAoAdicionarJogo(scoreA, scoreB, done)
	if q := len(s.jogos) % 2; q == 1 {
		s.ladoServico.Execute()
	}

	return nil
}

func (s Set) NovoJogo() *game.Game {
	newTurn := s.quemServe.Clonar(s.quemServe.LadoCorrente())
	result := game.New(newTurn, s.pontoDecisivo)
	return result
}

func (s Set) Placar() placares.EstadoEResultadoPlacar {
	return s.placar
}
