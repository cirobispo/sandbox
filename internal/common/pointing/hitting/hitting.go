package hitting

type HitType int
type HitUpdate int
type HitPointDestination int

type VerifyEvent func(hits *[]Hitting) HitPointDestination

type Hitting interface {
	GetType() HitType
	UpdateScore() HitUpdate
	PointDestination(hits *[]Hitting) HitPointDestination
}

const (
	HTFootFault HitType = 1
	HTAce       HitType = 2
	HTServeLet  HitType = 3
	HTServeOut  HitType = 4
	HTServeNet  HitType = 5
	HTServeIn   HitType = 6
	HTReturnOut HitType = 7
	HTReturnNet HitType = 8
	HTReturnIn  HitType = 9
	HTNet       HitType = 10
	HTIn        HitType = 11
	HTOut       HitType = 12
	HTWinner    HitType = 13
	HTToast     HitType = 14
	HTOther     HitType = 15
)

const (
	HTUNo          HitUpdate = 0
	HTUYes         HitUpdate = 1
	HTUCondicional HitUpdate = 2
)

const (
	HTDSameSide    HitPointDestination = 0
	HTDOpositeSide HitPointDestination = 1
	HTDNone        HitPointDestination = 2
)
