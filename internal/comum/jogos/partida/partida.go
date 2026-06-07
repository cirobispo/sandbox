package partida

import (
	"errors"

	"github.com/cirobispo/sandbox/internal/comum/jogos"
	"github.com/cirobispo/sandbox/internal/comum/jogos/jogo"
	"github.com/cirobispo/sandbox/internal/comum/jogos/set"
	"github.com/cirobispo/sandbox/internal/comum/placares"
	"github.com/cirobispo/sandbox/internal/comum/placares/placarpartida"
	"github.com/cirobispo/sandbox/internal/comum/turnos"
)

type ParamOption func(p *Partida) bool

type Setting interface {
	//	ServingSide() turning.TurningSide
	Score() placares.EstadoEParametroPlacar
	Games() []jogo.Jogo
}

type Partida struct {
	custom                bool
	tamanhoPartida        int
	tamanhoSet            int
	pontoDecisivo         bool
	tieBreak              bool
	placar                placarpartida.PlacarPartida
	sets                  []Setting
	eventosAoAdicionarSet []jogos.AoAdicionarSet
}

func PartidaPadrao() ParamOption {
	return func(score *Partida) bool {
		score.tamanhoPartida = 6
		score.tamanhoSet = 12
		score.pontoDecisivo = false
		score.tieBreak = true
		return false
	}
}

func PartidaCustomizada(tamanhoPartida, tamanhoSet int, pontoDecisivo, tieBreak bool) ParamOption {
	return func(m *Partida) bool {
		m.tamanhoPartida = tamanhoPartida
		m.tamanhoSet = tamanhoSet
		m.pontoDecisivo = pontoDecisivo
		m.tieBreak = tieBreak
		return true
	}
}

func New(param ParamOption) *Partida {
	result := &Partida{
		sets:                  make([]Setting, 0, 13),
		eventosAoAdicionarSet: make([]jogos.AoAdicionarSet, 0),
	}

	if param != nil {
		result.custom = param(result)

		callback := placarpartida.Padrao()
		if result.custom {
			callback = placarpartida.TamanhoELado(turnos.LTA, result.tamanhoPartida)
		}
		result.placar = placarpartida.New(callback)
	}

	return result
}

func (m Partida) executeEventosAoAdicionarSet(scoreA, scoreB int, done bool) {
	for j := range m.eventosAoAdicionarSet {
		event := m.eventosAoAdicionarSet[j]
		event(scoreA, scoreB, done)
	}
}

func (m *Partida) AdicionarEventoAoAdicionarSet(event jogos.AoAdicionarSet) {
	m.eventosAoAdicionarSet = append(m.eventosAoAdicionarSet, event)
}

func (m Partida) verificarAlgumErro(s Setting) error {
	if m.placar.Terminado() {
		return errors.New("set is closed.")
	}

	if !s.Score().Terminado() {
		return errors.New("game is still in play.")
	}

	return nil
}

func (m *Partida) AdicionarSet(s Setting) error {
	if err := m.verificarAlgumErro(s); err != nil {
		return err
	}

	m.sets = append(m.sets, s)
	m.placar.AdicionarPlacar(s.Score())

	scoreA, scoreB := m.placar.Resultado()
	done := m.placar.Terminado()
	m.executeEventosAoAdicionarSet(scoreA, scoreB, done)

	return nil
}

func (m Partida) CriaNovoSet() *set.Set {
	result := set.New(set.SetPadrao(nil))
	if m.custom {
		result = set.New(set.TurnoJogosETieBreak(nil, m.tamanhoSet, m.pontoDecisivo, m.tieBreak))
	}

	return result
}

func (m Partida) Placar() placares.EstadoEResultadoPlacar {
	return m.placar
}
