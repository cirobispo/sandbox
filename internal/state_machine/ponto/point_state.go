package ponto

import (
	"fmt"

	"github.com/cirobispo/sandbox/internal/comum/pontos/golpes"
	"github.com/cirobispo/sandbox/internal/comum/pontos/golpes/golpe"
	"github.com/cirobispo/sandbox/internal/comum/pontos/ponto"
)

type ExecuteOnPoint func(p *ponto.Ponto)

type PointState struct {
	hit            golpes.Golpes
	subPointsState []*PointState
	execute        ExecuteOnPoint
}

func NewPointState(hit golpes.Golpes) *PointState {
	return &PointState{
		hit:            hit,
		subPointsState: make([]*PointState, 0),
	}
}

func (s PointState) Hit() golpe.Golpe {
	result := golpe.New(s.hit.Tipo(), s.hit.Lado())
	return result
}

func (s *PointState) AddState(state *PointState) *PointState {
	s.subPointsState = append(s.subPointsState, state)
	return s
}

func (s *PointState) SubStates() []PointState {
	result := make([]PointState, 0, len(s.subPointsState))
	for j := range s.subPointsState {
		sub_item := *s.subPointsState[j]
		result = append(result, sub_item)
	}

	return result
}

func (s *PointState) Execute(p *ponto.Ponto) error {
	if s.execute == nil {
		return fmt.Errorf("execute function undefined.")
	}

	s.execute(p)
	return nil
}

func (s PointState) StatesToChoose() []string {
	result := make([]string, 0, len(s.subPointsState))
	for j := range s.subPointsState {
		state := s.subPointsState[j]
		result = append(result, fmt.Sprint((*state).hit.Tipo()))
	}

	return result
}
