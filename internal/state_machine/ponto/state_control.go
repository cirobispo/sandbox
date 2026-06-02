package ponto

import (
	"fmt"

	"github.com/cirobispo/sandbox/internal/comum/gaming/game"
	"github.com/cirobispo/sandbox/internal/comum/pontos/ponto"
)

type ExecuteOnGame func(g *game.Game)

type PointStateControl struct {
	point        *ponto.Ponto
	currentState *PointState
	states       []*PointState
}

func NewPointStateControl(point *ponto.Ponto) PointStateControl {
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
	c.point.AdicionaGolpe(s.Hit())
	s.Execute(c.point)

	if c.point.Terminado() {
		fmt.Printf("Ponto encerrado com %d hit(s)\n", c.point.Tamanho())
		fmt.Println()
		items := c.point.Golpes()
		for j := range items {
			item := items[j]
			fmt.Printf("%s\n", item.Tipo())
		}
	}
}

func (c *PointStateControl) CurrentState() *PointState {
	return c.currentState
}

func (c PointStateControl) BallInPlay() bool {
	return !c.point.Terminado()
}

func (c PointStateControl) FindState(state string) *PointState {
	var result *PointState
	for j := range c.states {
		item := c.states[j]
		if item.hit != nil && item.Hit().Tipo().String() == state {
			result = item
			break
		}
	}

	return result
}
