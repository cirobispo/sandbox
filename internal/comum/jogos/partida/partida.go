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
	placar                placarpartida.Partida
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

func (p Partida) executeEventosAoAdicionarSet() {
	if len(p.eventosAoAdicionarSet) > 0 {
		placarA, placarB := p.placar.Resultado()
		terminado := p.placar.Terminado()

		for j := range p.eventosAoAdicionarSet {
			event := p.eventosAoAdicionarSet[j]
			event(placarA, placarB, terminado)
		}
	}
}

func (m *Partida) AdicionarEventoAoAdicionarSet(event jogos.AoAdicionarSet) {
	m.eventosAoAdicionarSet = append(m.eventosAoAdicionarSet, event)
}

func (m Partida) verificarAlgumErro(s Setting) error {
	if m.placar.Terminado() {
		return errors.New("A partida está encerrada.")
	}

	if !s.Score().Terminado() {
		return errors.New("O set está em andamento.")
	}

	return nil
}

func (m *Partida) AdicionarSet(s Setting) error {
	if err := m.verificarAlgumErro(s); err != nil {
		return err
	}

	m.sets = append(m.sets, s)
	m.placar.AdicionarPlacar(s.Score())

	m.executeEventosAoAdicionarSet()

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
