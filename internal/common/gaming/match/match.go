package match

import (
	"errors"

	"github.com/cirobispo/sandbox/internal/common/gaming"
	"github.com/cirobispo/sandbox/internal/common/gaming/game"
	"github.com/cirobispo/sandbox/internal/common/gaming/set"
	"github.com/cirobispo/sandbox/internal/common/scoring"
	"github.com/cirobispo/sandbox/internal/common/scoring/matchscore"
	"github.com/cirobispo/sandbox/internal/common/turning"
)

type ParamOption func(set *Match) bool

type Setting interface {
	//	ServingSide() turning.TurningSide
	Score() scoring.EstadoEParametroPlacar
	Games() []game.Game
}

type Match struct {
	custom           bool
	matchSize        int
	setSize          int
	decidingPoint    bool
	tieBreak         bool
	score            matchscore.MatchScore
	sets             []Setting
	onAddingSetEvent []gaming.OnAfterAddingSet
}

func DefaultMatch() ParamOption {
	return func(score *Match) bool {
		score.matchSize = 6
		score.setSize = 12
		score.decidingPoint = false
		score.tieBreak = true
		return false
	}
}

func SizeSetSizeDecidindPointTieBreak(matchSize, setSize int, decidingPoint, tieBreak bool) ParamOption {
	return func(m *Match) bool {
		m.matchSize = matchSize
		m.setSize = setSize
		m.decidingPoint = decidingPoint
		m.tieBreak = tieBreak
		return true
	}
}

func New(param ParamOption) *Match {
	result := &Match{
		sets:             make([]Setting, 0, 13),
		onAddingSetEvent: make([]gaming.OnAfterAddingSet, 0),
	}

	if param != nil {
		result.custom = param(result)

		callback := matchscore.WithDefault()
		if result.custom {
			callback = matchscore.WithSideAndSize(turning.TSA, result.matchSize)
		}
		result.score = matchscore.New(callback)
	}

	return result
}

func (m Match) executeOnAfterAddingSet(scoreA, scoreB int, done bool) {
	for j := range m.onAddingSetEvent {
		event := m.onAddingSetEvent[j]
		event(scoreA, scoreB, done)
	}
}

func (m *Match) AddOnAddingSetEvent(event gaming.OnAfterAddingSet) {
	m.onAddingSetEvent = append(m.onAddingSetEvent, event)
}

func (m Match) checkAnyError(s Setting) error {
	if m.score.Terminado() {
		return errors.New("set is closed.")
	}

	if !s.Score().Terminado() {
		return errors.New("game is still in play.")
	}

	return nil
}

func (m *Match) AddSet(s Setting) error {
	if err := m.checkAnyError(s); err != nil {
		return err
	}

	m.sets = append(m.sets, s)
	m.score.AddScore(s.Score())

	scoreA, scoreB := m.score.Resultado()
	done := m.score.Terminado()
	m.executeOnAfterAddingSet(scoreA, scoreB, done)

	return nil
}

func (m Match) NewSet() *set.Set {
	result := set.New(set.WithDefaultSet(nil))
	if m.custom {
		result = set.New(set.WithTurnSizeAndTieBreak(nil, m.setSize, m.decidingPoint, m.tieBreak))
	}

	return result
}

func (m Match) Score() scoring.EstadoEResultadoPlacar {
	return m.score
}
