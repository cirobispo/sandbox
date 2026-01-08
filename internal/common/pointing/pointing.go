package pointing

import "github.com/cirobispo/sandbox/internal/common/pointing/hitting"

type OnPointScore func(hit hitting.HitType, side hitting.HitSide)

type PointSide int

const (
	PSA    PointSide = 0
	PSB    PointSide = 1
	PSNone PointSide = 2
)

type Pointing interface {
}

type OnOcours func()
