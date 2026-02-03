package hitting

type HitType int
type HitSide int

const (
	HTFootFault   HitType = 1
	HTServeNet    HitType = 2
	HTServeLet    HitType = 3
	HTServeIn     HitType = 4
	HTAce         HitType = 5
	HTServeOut    HitType = 6
	HTReturnOut   HitType = 7
	HTReturnNet   HitType = 8
	HTReturnIn    HitType = 9
	HTDoubleFault HitType = 10
	HTNet         HitType = 11
	HTIn          HitType = 12
	HTOut         HitType = 13
	HTWinner      HitType = 14
	HTToast       HitType = 15
	HTNetTouch    HitType = 16
	HTMiss        HitType = 17
)

const (
	HTDNone         HitSide = 0
	HTDChangeSide   HitSide = 1
	HTDSameSide     HitSide = 2
	HTDOppositeSide HitSide = 3
	HTDConditional  HitSide = 4
)

type Hitting interface {
	Type() HitType
	Side() HitSide
}

func (h HitSide) Inverse() HitSide {
	switch h {
	case HTDSameSide:
		return HTDOppositeSide
	case HTDOppositeSide:
		return HTDSameSide
	default:
		return h
	}
}

func (h HitType) String() string {
	switch h {
	case HTFootFault:
		return "Foot fault"
	case HTServeNet:
		return "Serve on net"
	case HTServeLet:
		return "Let"
	case HTServeIn:
		return "Serve in"
	case HTAce:
		return "Ace!"
	case HTServeOut:
		return "Serve out"
	case HTReturnOut:
		return "Return out"
	case HTReturnNet:
		return "Return net"
	case HTReturnIn:
		return "Return in"
	case HTDoubleFault:
		return "Double fault"
	case HTNet:
		return "Hit net"
	case HTIn:
		return "Hit in"
	case HTOut:
		return "Hit out"
	case HTWinner:
		return "Winner!"
	case HTToast:
		return "Toast!"
	case HTNetTouch:
		return "Touch net"
	case HTMiss:
		return "Miss"
	default:
		return "Other"
	}
}

func (h HitSide) String() string {
	switch h {
	case HTDNone:
		return "None"
	case HTDChangeSide:
		return "Change Side"
	case HTDSameSide:
		return "Same side"
	case HTDOppositeSide:
		return "Opposite side"
	case HTDConditional:
		return "Conditional"
	default:
		return "other"
	}
}
