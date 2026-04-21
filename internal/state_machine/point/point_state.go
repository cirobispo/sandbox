package point

import (
	"fmt"

	"github.com/cirobispo/sandbox/internal/common/pointing/hitting"
	"github.com/cirobispo/sandbox/internal/common/pointing/hitting/hit"
	"github.com/cirobispo/sandbox/internal/common/pointing/point"
)

type ExecuteOnPoint func(p *point.Point)

type PointState struct {
	hit            hitting.Hitting
	subPointsState []*PointState
	execute        ExecuteOnPoint
}

func NewPointState(hit hitting.Hitting) *PointState {
	return &PointState{
		hit:            hit,
		subPointsState: make([]*PointState, 0),
	}
}

func (s PointState) Hit() hit.Hit {
	result := hit.New(s.hit.Type(), s.hit.Side())
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

func (s *PointState) Execute(p *point.Point) error {
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
		result = append(result, fmt.Sprint((*state).hit.Type()))
	}

	return result
}
