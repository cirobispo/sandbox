package pointing

import "github.com/cirobispo/sandbox/internal/common/pointing/hitting"

type OnPointScore func(hit hitting.HitType, side hitting.HitSide, done bool)

type PointSide int

const (
	PSStartingSide PointSide = 1
	PSOppositeSide PointSide = 2
	PSNone         PointSide = 0
)

func (s PointSide) String() string {
	switch s {
	case PSStartingSide:
		return "Starting side"
	case PSOppositeSide:
		return "Opposite side"
	default:
		return "None"
	}
}

func (s PointSide) Inverse() PointSide {
	switch s {
	case PSStartingSide:
		return PSOppositeSide
	case PSOppositeSide:
		return PSStartingSide
	default:
		return s
	}
}
