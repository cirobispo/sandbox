package point

import (
	"github.com/cirobispo/sandbox/internal/common/pointing"
	"github.com/cirobispo/sandbox/internal/common/pointing/hitting/hit"
)

type TestItem struct {
	Value hit.Hit
}

type TestBlock struct {
	Items []TestItem
	Point pointing.PointSide
}

var hitServInReturnIn = []TestItem{{hit.NewServeIn()}, {hit.NewReturnIn()}}

var __winnerSameSide = []TestItem{{hit.NewWinner()}}

var __winnerOppositeSide = []TestItem{{hit.NewHitBackIn()}, {hit.NewWinner()}}

var hitAce = []TestItem{{hit.NewAce()}}

var __longRallie = []TestItem{{hit.NewHitBackIn()}, {hit.NewHitBackIn()}}

var hit__NetSameSide = []TestItem{{hit.NewHitBackIn()}, {hit.NewHitBackIn()}, {hit.NewHitNet()}}

var hit__OutSameSide = []TestItem{{hit.NewHitBackIn()}, {hit.NewHitBackIn()}, {hit.NewHitOut()}}

var hit__NetOppositeSide = []TestItem{{hit.NewHitBackIn()}, {hit.NewHitNet()}}

var hit__OutOppositeSide = []TestItem{{hit.NewHitBackIn()}, {hit.NewHitOut()}}

var doubleFault = []TestItem{{hit.NewServeNet()}, {hit.NewServeOut()}}

func DoubleFault() TestBlock {
	data := make([]TestItem, 0, 2)
	data = append(data, doubleFault...)
	return TestBlock{Items: data}
}

func WinnerSSPoint() TestBlock {
	data := hitServInReturnIn
	data = append(data, __winnerSameSide...)
	return TestBlock{Items: data}
}

func WinnerOSPoint() TestBlock {
	data := hitServInReturnIn
	data = append(data, __winnerOppositeSide...)
	return TestBlock{Items: data}
}

func AcePoint() TestBlock {
	data := hitAce
	return TestBlock{Items: data}
}

func LongRallieOSPoint(hits int, subBlock TestBlock) TestBlock {
	data := hitServInReturnIn
	for hits > 0 {
		data = append(data, __longRallie...)
		hits--
	}
	data = append(data, subBlock.Items...)

	return TestBlock{Items: data, Point: subBlock.Point}
}

func NetSameSide(toCompose bool) TestBlock {
	data := make([]TestItem, 0, 3)
	if !toCompose {
		data = append(data, hitServInReturnIn...)
	}
	data = append(data, hit__NetSameSide...)
	return TestBlock{Items: data}
}

func OutSameSide(toCompose bool) TestBlock {
	data := make([]TestItem, 0, 3)
	if !toCompose {
		data = append(data, hitServInReturnIn...)
	}
	data = append(data, hit__OutSameSide...)
	return TestBlock{Items: data}
}

func NetOppositeSide(toCompose bool) TestBlock {
	data := make([]TestItem, 0, 3)
	if !toCompose {
		data = append(data, hitServInReturnIn...)
	}
	data = append(data, hit__NetOppositeSide...)
	return TestBlock{Items: data}
}

func OutOppositeSide(toCompose bool) TestBlock {
	data := make([]TestItem, 0, 3)
	if !toCompose {
		data = append(data, hitServInReturnIn...)
	}
	data = append(data, hit__OutOppositeSide...)
	return TestBlock{Items: data}
}
