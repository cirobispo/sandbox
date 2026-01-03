package point_test

import (
	"testing"

	"github.com/cirobispo/sandbox/internal/common/pointing/hitting"
	"github.com/cirobispo/sandbox/internal/common/pointing/hitting/hit"
)

type testItem struct {
	arg1, arg2, arg3 hit.Hit
	expected         hitting.HitSide
}

type testBlock struct {
	items []testItem
}

// var longRallie = []testItem{
// 	testItem{hit.NewServeIn(), hitting.HTDNone},
// 	testItem{hit.NewIn(), hitting.HTDNone},
// 	testItem{hit.NewIn(), hitting.HTDNone},
// 	testItem{hit.NewIn(), hitting.HTDNone},
// 	testItem{hit.NewIn(), hitting.HTDNone},
// 	testItem{hit.NewIn(), hitting.HTDNone},
// 	testItem{hit.NewIn(), hitting.HTDNone},
// 	testItem{hit.NewIn(), hitting.HTDNone},
// 	testItem{hit.NewIn(), hitting.HTDNone},
// 	testItem{hit.NewIn(), hitting.HTDNone},
// 	testItem{hit.NewWinner(), hitting.HTDSameSide},
// }

// var hitWinner = []testItem{
// 	testItem{hit.NewServeIn(), hitting.HTDNone},
// 	testItem{hit.NewIn(), hitting.HTDNone},
// 	testItem{hit.NewWinner(), hitting.HTDSameSide},
// }

// var testblock = []testBlock{
// 	testBlock{items: hitWinner},
// 	testBlock{items: longRallie},
// }

func TestAddHits(tt *testing.T) {
	// t := turn.New(turning.STA)
	// p := point.New(*t)
	// for b := range testblock {
	// 	block := testblock[b]
	// 	for i := range block.items {
	// 		test := block.items[i]
	// 		if output := p.AddHit(test.arg3); output != test.expected {
	// 			tt.Errorf("Output %q not equal to expected %q", output, test.expected)
	// 		}
	// 	}
	// }
}
