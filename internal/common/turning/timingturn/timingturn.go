package timingturn

import (
	"time"

	"github.com/cirobispo/sandbox/internal/common/turning"
	"github.com/cirobispo/sandbox/internal/common/turning/turn"
)

type TimedTurning interface {
	turning.Turning
	Duration(t turn.Turn) int
}

func newFromTurn(t *turn.Turn) *turn.Turn {
	turn.AddData[time.Time](t, "start", time.Now())
	result := t

	result.AddAfterChangeEvent(func(st turning.SideTurn) {
		start := turn.GetData[time.Time](t, "start")
		duration := time.Since(start)

		turn.AddData[time.Duration](t, "duration", duration)
	})

	return result
}

func New(start turning.SideTurn) *turn.Turn {
	t := turn.New(start)
	return NewFromTurn(t)
}

func NewFromTurn(t *turn.Turn) *turn.Turn {
	return newFromTurn(t)
}

func Duration(t *turn.Turn) time.Duration {
	duration := turn.GetData[time.Duration](t, "duration")
	return duration
}
