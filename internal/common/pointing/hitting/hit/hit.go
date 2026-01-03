package hit

import (
	"github.com/cirobispo/sandbox/internal/common/pointing/hitting"
)

type Hit struct {
	pointType hitting.HitType
	pointSide hitting.HitSide
}

func (h Hit) Type() hitting.HitType {
	return h.pointType
}

func (h Hit) Side() hitting.HitSide {
	return h.pointSide
}

func NewFootFault() Hit {
	return Hit{
		pointType: hitting.HTFootFault,
		pointSide: hitting.HTDConditional,
	}
}

func NewAce() Hit {
	return Hit{
		pointType: hitting.HTAce,
		pointSide: hitting.HTDSameSide,
	}
}

func NewServeOut() Hit {
	return Hit{
		pointType: hitting.HTServeOut,
		pointSide: hitting.HTDConditional,
	}
}

func NewServeIn() Hit {
	return Hit{
		pointType: hitting.HTServeIn,
		pointSide: hitting.HTDNone,
	}
}

func NewServeLet() Hit {
	return Hit{
		pointType: hitting.HTServeLet,
		pointSide: hitting.HTDNone,
	}
}

func NewServeNet() Hit {
	return Hit{
		pointType: hitting.HTServeNet,
		pointSide: hitting.HTDConditional,
	}
}

func NewReturnNet() Hit {
	return Hit{
		pointType: hitting.HTReturnNet,
		pointSide: hitting.HTDOpositeSide,
	}
}

func NewReturnIn() Hit {
	return Hit{
		pointType: hitting.HTReturnIn,
		pointSide: hitting.HTDNone,
	}
}

func NewReturnOut() Hit {
	return Hit{
		pointType: hitting.HTReturnOut,
		pointSide: hitting.HTDOpositeSide,
	}
}

func NewNet() Hit {
	return Hit{
		pointType: hitting.HTNet,
		pointSide: hitting.HTDOpositeSide,
	}
}

func NewIn() Hit {
	return Hit{
		pointType: hitting.HTIn,
		pointSide: hitting.HTDNone,
	}
}

func NewWinner() Hit {
	return Hit{
		pointType: hitting.HTWinner,
		pointSide: hitting.HTDSameSide,
	}
}

func NewOut() Hit {
	return Hit{
		pointType: hitting.HTOut,
		pointSide: hitting.HTDOpositeSide,
	}
}

func NewMiss() Hit {
	return Hit{
		pointType: hitting.HTOut,
		pointSide: hitting.HTDOpositeSide,
	}
}
