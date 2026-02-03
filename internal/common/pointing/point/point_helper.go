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

var hitServInReturnIn = []TestItem{
	TestItem{hit.NewServeIn()},
	TestItem{hit.NewReturnIn()},
}

var __winnerSameSide = []TestItem{
	TestItem{hit.NewWinner()},
}

var __winnerOppositeSide = []TestItem{
	TestItem{hit.NewIn()},
	TestItem{hit.NewWinner()},
}

var hitAce = []TestItem{
	TestItem{hit.NewAce()},
}

var __longRallie = []TestItem{
	TestItem{hit.NewIn()},
	TestItem{hit.NewIn()},
	TestItem{hit.NewIn()},
	TestItem{hit.NewIn()},
	TestItem{hit.NewIn()},
	TestItem{hit.NewIn()},
	TestItem{hit.NewIn()},
	TestItem{hit.NewIn()},
}

var hit__NetSameSide = []TestItem{
	TestItem{hit.NewIn()},
	TestItem{hit.NewIn()},
	TestItem{hit.NewNet()},
}

var hit__OutSameSide = []TestItem{
	TestItem{hit.NewIn()},
	TestItem{hit.NewIn()},
	TestItem{hit.NewOut()},
}

var hit__NetOppositeSide = []TestItem{
	TestItem{hit.NewIn()},
	TestItem{hit.NewNet()},
}

var hit__OutOppositeSide = []TestItem{
	TestItem{hit.NewIn()},
	TestItem{hit.NewOut()},
}

var doubleFault = []TestItem{
	TestItem{hit.NewServeNet()},
	TestItem{hit.NewServeOut()},
}

// func addHits2Test(tt *testing.T, items *[]TestItem, ps pointing.PointSide) TestBlock {
// 	tb := TestBlock{items: *items}
// 	ctt := timingturn.NewFromTurn(countingturn.NewFromTurn(turn.New(turning.STA)))
// 	p := New(ctt)
// 	for i := range tb.items {
// 		item := tb.items[i]
// 		tt.Logf("Testing hit type: %d, giving point to side: %v\n", item.value.Type(), item.value.Side())
// 		p.AddHit(item.value)
// 	}
// 	tb.point = ps
// 	return tb
// }

//	func addPoints2Test(tt *testing.T, points *[]TestBlock, results []pointing.PointSide) {
//		for x := range *points {
//			items, result := (*points)[x], results[x]
//			addHits2Test(tt, &items.items, result)
//		}
//	}
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

func LongRallieOSPoint(subBlock TestBlock) TestBlock {
	data := hitServInReturnIn
	data = append(data, __longRallie...)
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
