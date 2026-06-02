package ponto

import (
	"github.com/cirobispo/sandbox/internal/common/pontos"
	"github.com/cirobispo/sandbox/internal/common/pontos/golpes/golpe"
)

type TestItem struct {
	Value golpe.Hit
}

type TestBlock struct {
	Items []TestItem
	Point pontos.LadoDoPonto
}

var hitServInReturnIn = []TestItem{{golpe.NewServeIn()}, {golpe.NewReturnIn()}}

var __winnerSameSide = []TestItem{{golpe.NewWinner()}}

var __winnerOppositeSide = []TestItem{{golpe.NewHitBackIn()}, {golpe.NewWinner()}}

var hitAce = []TestItem{{golpe.NewAce()}}

var __longRallie = []TestItem{{golpe.NewHitBackIn()}, {golpe.NewHitBackIn()}}

var hit__NetSameSide = []TestItem{{golpe.NewHitBackIn()}, {golpe.NewHitBackIn()}, {golpe.NewHitNet()}}

var hit__OutSameSide = []TestItem{{golpe.NewHitBackIn()}, {golpe.NewHitBackIn()}, {golpe.NewHitOut()}}

var hit__NetOppositeSide = []TestItem{{golpe.NewHitBackIn()}, {golpe.NewHitNet()}}

var hit__OutOppositeSide = []TestItem{{golpe.NewHitBackIn()}, {golpe.NewHitOut()}}

var doubleFault = []TestItem{{golpe.NewServeNet()}, {golpe.NewServeOut()}}

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
