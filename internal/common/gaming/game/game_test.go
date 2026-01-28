package game_test

import (
	"fmt"
	"testing"

	"github.com/cirobispo/sandbox/internal/common/gaming/game"
	"github.com/cirobispo/sandbox/internal/common/pointing"
	"github.com/cirobispo/sandbox/internal/common/pointing/hitting/hit"
	"github.com/cirobispo/sandbox/internal/common/pointing/point"
	"github.com/cirobispo/sandbox/internal/common/turning"
	"github.com/cirobispo/sandbox/internal/common/turning/turn"
)

type testItem struct {
	value hit.Hit
}

type testBlock struct {
	items []testItem
	point point.Point
}

var hitServInReturnIn = []testItem{
	testItem{hit.NewServeIn()},
	testItem{hit.NewReturnIn()},
}

var __winnerSameSide = []testItem{
	testItem{hit.NewWinner()},
}

var __winnerOppositeSide = []testItem{
	testItem{hit.NewIn()},
	testItem{hit.NewWinner()},
}

var hitAce = []testItem{
	testItem{hit.NewServeIn()},
}

var longRallie = []testItem{
	testItem{hit.NewServeIn()},
	testItem{hit.NewReturnIn()},
	testItem{hit.NewIn()},
	testItem{hit.NewIn()},
	testItem{hit.NewIn()},
	testItem{hit.NewIn()},
	testItem{hit.NewIn()},
	testItem{hit.NewIn()},
	testItem{hit.NewIn()},
	testItem{hit.NewIn()},
	testItem{hit.NewIn()},
	testItem{hit.NewNet()},
}

var hit__NetSameSide = []testItem{
	testItem{hit.NewIn()},
	testItem{hit.NewIn()},
	testItem{hit.NewNet()},
}

var hit__OutSameSide = []testItem{
	testItem{hit.NewIn()},
	testItem{hit.NewIn()},
	testItem{hit.NewOut()},
}

var _game = game.New(turn.New(turning.STA), false)

func addToTest(items *[]testItem) testBlock {
	tb := testBlock{items: *items}
	p := point.New(_game.NewTurn())
	for i := range tb.items {
		item := tb.items[i]
		p.AddHit(item.value)
	}
	tb.point = p
	_game.AddPoint(p)
	return tb
}

func winnerSSPoint() testBlock {
	data := hitServInReturnIn
	data = append(data, __winnerSameSide...)
	return addToTest(&data)
}

func acePoint() testBlock {
	return addToTest(&hitAce)
}

func longRallieOSPoint() testBlock {
	return addToTest(&longRallie)
}

func netSameSide() testBlock {
	data := hitServInReturnIn
	data = append(data, hit__NetSameSide...)
	return addToTest(&data)
}

func outSameSide() testBlock {
	data := hitServInReturnIn
	data = append(data, hit__OutSameSide...)
	return addToTest(&data)
}

func testPoint(tt *testing.T, block testBlock, name string) {
	tt.Logf("Start of test for point %s\n", name)
	pointResult := block.point.PointSide()
	if pointResult == pointing.PSNone {
		tt.Errorf("%s is not correct on result point: %v", name, block.point)
	}

	a, b := _game.Result()
	tt.Logf("Game result score A: %d score B : %d\n", a, b)
}

func testGame(tt *testing.T, blocks *[]testBlock, name string) {
	tt.Logf("Start of test for game %s\n", name)
	for i := range *blocks {
		block := (*blocks)[i]
		testPoint(tt, block, fmt.Sprintf("%d", i+1))
	}
}

func TestWinner(tt *testing.T) {
	testPoint(tt, winnerSSPoint(), "Winner")
}

func TestAce(tt *testing.T) {
	testPoint(tt, acePoint(), "Ace")
}

func TestLongRallie(tt *testing.T) {
	testPoint(tt, longRallieOSPoint(), "Long rallie")
}

func TestNetSameSide(tt *testing.T) {
	testPoint(tt, netSameSide(), "Net same side")
}

func TestOutSameSide(tt *testing.T) {
	testPoint(tt, outSameSide(), "Out same side")
}

func TestGame(tt *testing.T) {
	var blocks []testBlock = []testBlock{
		acePoint(),
		winnerSSPoint(),
		longRallieOSPoint(),
		winnerSSPoint(),
		netSameSide(),
		winnerSSPoint(),
		outSameSide(),
		winnerSSPoint(),
	}
	testGame(tt, &blocks, "6x2")
}
