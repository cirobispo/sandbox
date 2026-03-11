package pointing

import "github.com/cirobispo/sandbox/internal/common/pointing/hitting"

type OnPointScore func(hit hitting.HitType, side hitting.HitSide, done bool)

type PointSide int

const (
	PSServing  PointSide = 1
	PSOpposite PointSide = 2
	PSNone     PointSide = 0
)

func (s PointSide) String() string {
	switch s {
	case PSServing:
		return "Serving side"
	case PSOpposite:
		return "Opposite side"
	default:
		return "None"
	}
}

func (s PointSide) Inverse() PointSide {
	switch s {
	case PSServing:
		return PSOpposite
	case PSOpposite:
		return PSServing
	default:
		return s
	}
}
