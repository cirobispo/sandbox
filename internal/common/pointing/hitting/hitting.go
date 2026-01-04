package hitting

import "fmt"

type HitType int
type HitSide int

const (
	HTFootFault HitType = 1
	HTServeNet  HitType = 2
	HTServeLet  HitType = 3
	HTServeIn   HitType = 4
	HTAce       HitType = 5
	HTServeOut  HitType = 6
	HTReturnOut HitType = 7
	HTReturnNet HitType = 8
	HTReturnIn  HitType = 9
	HTMiss      HitType = 10
	HTNet       HitType = 11
	HTIn        HitType = 12
	HTOut       HitType = 13
	HTWinner    HitType = 14
	HTToast     HitType = 15
	HTOther     HitType = 16
)

const (
	HTDNone        HitSide = 0
	HTDChangeSide  HitSide = 1
	HTDSameSide    HitSide = 2
	HTDOpositeSide HitSide = 3
	HTDConditional HitSide = 4
)

func (h HitType) String() string {
	return fmt.Sprintf("%c", h)
}

func (h HitSide) String() string {
	return fmt.Sprintf("%d", h)
}

type Hitting interface {
	Type() HitType
	Side() HitSide
}
