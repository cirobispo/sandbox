package point

import (
	"fmt"

	"github.com/cirobispo/sandbox/internal/common/gaming/game"
	"github.com/cirobispo/sandbox/internal/common/pointing/point"
)

type ExecuteOnGame func(g *game.Game)

type PointStateControl struct {
	point        *point.Point
	currentState *PointState
	states       []*PointState
}

func NewPointStateControl(point *point.Point) PointStateControl {
	states := []*PointState{PointStarting(),
		AfterServeIn(), AfterServeNet(), AfterServeOut(), AfterServeLet(), AfterHitAce(),
		AfterReturnIn(), AfterReturnNet(), AfterReturnOut(),
		AfterHitBackIn(), AfterHitBackOut(), AfterHitBackNet(),
	}

	return PointStateControl{
		point:        point,
		currentState: states[0],
		states:       states,
	}
}

func (c *PointStateControl) UpdateState(s *PointState) {
	c.currentState = s
	c.point.AddHit(s.Hit())
	s.Execute(c.point)

	if c.point.Done() {
		fmt.Printf("Ponto encerrado com %d hit(s)\n", c.point.Length())
		fmt.Println()
		items := c.point.Hits()
		for j := range items {
			item := items[j]
			fmt.Printf("%s\n", item.Type())
		}
	}
}

func (c *PointStateControl) CurrentState() *PointState {
	return c.currentState
}

func (c PointStateControl) BallInPlay() bool {
	return !c.point.Done()
}

func (c PointStateControl) FindState(state string) *PointState {
	var result *PointState
	for j := range c.states {
		item := c.states[j]
		if item.hit != nil && item.Hit().Type().String() == state {
			result = item
			break
		}
	}

	return result
}
