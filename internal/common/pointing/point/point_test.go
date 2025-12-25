package point_test

import (
	"testing"

	"github.com/cirobispo/sandbox/internal/common/pointing/hitting"
	"github.com/cirobispo/sandbox/internal/common/pointing/hitting/hit"
	"github.com/cirobispo/sandbox/internal/common/pointing/point"
	"github.com/cirobispo/sandbox/internal/common/turning"
	"github.com/cirobispo/sandbox/internal/common/turning/turn"
)

type addTest struct {
	arg      hit.Hit
	expected hitting.HitPointDestination
}

var addTests = []addTest{
	addTest{hit.NewServeIn(), hitting.HTDNone},
	addTest{hit.NewIn(), hitting.HTDNone},
	addTest{hit.NewIn(), hitting.HTDNone},
	addTest{hit.NewIn(), hitting.HTDNone},
	addTest{hit.NewIn(), hitting.HTDNone},
	addTest{hit.NewIn(), hitting.HTDNone},
	addTest{hit.NewIn(), hitting.HTDNone},
	addTest{hit.NewIn(), hitting.HTDNone},
	addTest{hit.NewIn(), hitting.HTDNone},
	addTest{hit.NewIn(), hitting.HTDNone},
	addTest{hit.NewWinner(), hitting.HTDSameSide},
}

func TestAddHits(tt *testing.T) {
	t := turn.New(turning.STA)
	p := point.New(*t)
	for i := range addTests {
		test := addTests[i]
		if output := p.AddHit(test.arg); output != test.expected {
			tt.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}
