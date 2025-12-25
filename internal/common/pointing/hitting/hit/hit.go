package hit

import (
	"github.com/cirobispo/sandbox/internal/common/pointing/hitting"
)

type Hit struct {
	pointType   hitting.HitType
	updateScore hitting.HitUpdate
	verify      hitting.VerifyEvent
}

func (g Hit) GetType() hitting.HitType {
	return g.pointType
}

func (g Hit) UpdateScore() hitting.HitUpdate {
	return g.updateScore
}

func (g Hit) PointDestination(hits *[]hitting.Hitting) hitting.HitPointDestination {
	if g.verify != nil {
		return g.verify(hits)
	}

	return returnNone(hits)
}

func returnSameSide(hits *[]hitting.Hitting) hitting.HitPointDestination {
	return hitting.HTDSameSide
}

func returnOppositeSide(hits *[]hitting.Hitting) hitting.HitPointDestination {
	return hitting.HTDOpositeSide
}

func returnNone(hits *[]hitting.Hitting) hitting.HitPointDestination {
	return hitting.HTDNone
}

func isDoubleFault(hits *[]hitting.Hitting) hitting.HitPointDestination {
	count := 0
	result := hitting.HTDNone
	for i := range *hits {
		hit := (*hits)[i]
		if tp := hit.GetType(); tp == hitting.HTServeOut || tp == hitting.HTServeNet || tp == hitting.HTFootFault {
			count++
			if count > 1 {
				result = hitting.HTDOpositeSide
				break
			}
		}
	}
	return result
}

func NewAce() Hit {
	return Hit{
		pointType:   hitting.HTAce,
		updateScore: hitting.HTUYes,
		verify:      returnSameSide,
	}
}

func NewServeLet() Hit {
	return Hit{
		pointType:   hitting.HTServeLet,
		updateScore: hitting.HTUNo,
		verify:      returnNone,
	}
}

func NewServeIn() Hit {
	return Hit{
		pointType:   hitting.HTServeIn,
		updateScore: hitting.HTUNo,
		verify:      returnNone,
	}
}

func NewServeOut() Hit {
	return Hit{
		pointType:   hitting.HTServeOut,
		updateScore: hitting.HTUCondicional,
		verify:      isDoubleFault,
	}
}

func NewServeNet() Hit {
	return Hit{
		pointType:   hitting.HTServeNet,
		updateScore: hitting.HTUCondicional,
		verify:      isDoubleFault,
	}
}

func NewReturnOut() Hit {
	return Hit{
		pointType:   hitting.HTReturnOut,
		updateScore: hitting.HTUYes,
		verify:      returnOppositeSide,
	}
}

func NewReturnNet() Hit {
	return Hit{
		pointType:   hitting.HTReturnNet,
		updateScore: hitting.HTUYes,
		verify:      returnOppositeSide,
	}
}

func NewReturnIn() Hit {
	return Hit{
		pointType:   hitting.HTReturnIn,
		updateScore: hitting.HTUNo,
		verify:      returnNone,
	}
}

func NewIn() Hit {
	return Hit{
		pointType:   hitting.HTIn,
		updateScore: hitting.HTUNo,
		verify:      returnNone,
	}
}

func NewWinner() Hit {
	return Hit{
		pointType:   hitting.HTWinner,
		updateScore: hitting.HTUYes,
		verify:      returnSameSide,
	}
}

func NewOut(verify hitting.VerifyEvent) Hit {
	return Hit{
		pointType:   hitting.HTOut,
		updateScore: hitting.HTUYes,
		verify:      verify,
	}
}

func NewNet(verify hitting.VerifyEvent) Hit {
	return Hit{
		pointType:   hitting.HTNet,
		updateScore: hitting.HTUYes,
		verify:      verify,
	}
}
