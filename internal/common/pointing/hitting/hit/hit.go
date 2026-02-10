package hit

import (
	"github.com/cirobispo/sandbox/internal/common/pointing/hitting"
)

type Hit struct {
	pointType hitting.HitType
	side      hitting.HitSide
}

func (h Hit) Type() hitting.HitType {
	return h.pointType
}

func (h Hit) Side() hitting.HitSide {
	return h.side
}

func New(t hitting.HitType, s hitting.HitSide) Hit {
	return Hit{
		pointType: t,
		side:      s,
	}
}

func NewFootFault() Hit {
	return Hit{
		pointType: hitting.HTFootFault,
		side:      hitting.HTDConditional,
	}
}

func NewAce() Hit {
	return Hit{
		pointType: hitting.HTAce,
		side:      hitting.HTDSameSide,
	}
}

func NewServeOut() Hit {
	return Hit{
		pointType: hitting.HTServeOut,
		side:      hitting.HTDConditional,
	}
}

func NewServeIn() Hit {
	return Hit{
		pointType: hitting.HTServeIn,
		side:      hitting.HTDChangeSide,
	}
}

func NewServeLet() Hit {
	return Hit{
		pointType: hitting.HTServeLet,
		side:      hitting.HTDNone,
	}
}

func NewServeNet() Hit {
	return Hit{
		pointType: hitting.HTServeNet,
		side:      hitting.HTDConditional,
	}
}

func NewReturnNet() Hit {
	return Hit{
		pointType: hitting.HTReturnNet,
		side:      hitting.HTDOppositeSide,
	}
}

func NewReturnIn() Hit {
	return Hit{
		pointType: hitting.HTReturnIn,
		side:      hitting.HTDChangeSide,
	}
}

func NewReturnOut() Hit {
	return Hit{
		pointType: hitting.HTReturnOut,
		side:      hitting.HTDOppositeSide,
	}
}

func NewDoubleFault() Hit {
	return Hit{
		pointType: hitting.HTDoubleFault,
		side:      hitting.HTDOppositeSide,
	}
}

func NewHitNet() Hit {
	return Hit{
		pointType: hitting.HTNet,
		side:      hitting.HTDOppositeSide,
	}
}

func NewHitBackIn() Hit {
	return Hit{
		pointType: hitting.HTIn,
		side:      hitting.HTDChangeSide,
	}
}

func NewWinner() Hit {
	return Hit{
		pointType: hitting.HTWinner,
		side:      hitting.HTDSameSide,
	}
}

func NewHitOut() Hit {
	return Hit{
		pointType: hitting.HTOut,
		side:      hitting.HTDOppositeSide,
	}
}

func NewMiss() Hit {
	return Hit{
		pointType: hitting.HTOut,
		side:      hitting.HTDOppositeSide,
	}
}

func NewToast() Hit {
	return Hit{
		pointType: hitting.HTToast,
		side:      hitting.HTDSameSide,
	}
}

func NewNetTouch() Hit {
	return Hit{
		pointType: hitting.HTNetTouch,
		side:      hitting.HTDOppositeSide,
	}
}
