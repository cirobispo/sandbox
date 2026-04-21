package point

import (
	"github.com/cirobispo/sandbox/internal/common/pointing/hitting/hit"
)

func PointStarting() *PointState {
	result := NewPointState(nil).
		AddState(NewPointState(hit.NewAce())).
		AddState(NewPointState(hit.NewNetTouch())).
		AddState(NewPointState(hit.NewServeNet())).
		AddState(NewPointState(hit.NewServeLet())).
		AddState(NewPointState(hit.NewServeIn())).
		AddState(NewPointState(hit.NewServeOut()))

	return result
}

func addFromPointStarting(stateToAdd *PointState) {
	ps := PointStarting()
	for j := range ps.subPointsState {
		item := ps.subPointsState[j]
		stateToAdd.AddState(item)
	}
}

func AfterHitAce() *PointState {
	result := NewPointState(hit.NewAce())
	addFromPointStarting(result)
	return result
}

func AfterServeIn() *PointState {
	result := NewPointState(hit.NewServeIn()).
		AddState(NewPointState(hit.NewReturnNet())).
		AddState(NewPointState(hit.NewReturnIn())).
		AddState(NewPointState(hit.NewReturnOut())).
		AddState(NewPointState(hit.NewToast()))

	return result
}

func AfterServeOut() *PointState {
	result := NewPointState(hit.NewServeOut())
	addFromPointStarting(result)
	return result
}

func AfterServeNet() *PointState {
	result := NewPointState(hit.NewServeNet())
	addFromPointStarting(result)

	return result
}

func AfterServeLet() *PointState {
	result := NewPointState(hit.NewServeLet())
	addFromPointStarting(result)

	return result
}

func AfterReturnIn() *PointState {
	result := NewPointState(hit.NewReturnIn()).
		AddState(NewPointState(hit.NewHitNet())).
		AddState(NewPointState(hit.NewHitBackIn())).
		AddState(NewPointState(hit.NewHitOut())).
		AddState(NewPointState(hit.NewToast()))

	return result
}

func AfterReturnOut() *PointState {
	result := NewPointState(hit.NewReturnOut())
	addFromPointStarting(result)

	return result
}

func AfterReturnNet() *PointState {
	result := NewPointState(hit.NewReturnNet())
	addFromPointStarting(result)

	return result
}

func addFromHitBackIn(stateToAdd *PointState) {
	hbi := AfterHitBackIn()
	for j := range hbi.subPointsState {
		item := hbi.subPointsState[j]
		stateToAdd.AddState(item)
	}
}

func AfterHitBackIn() *PointState {
	result := NewPointState(hit.NewHitBackIn()).
		AddState(NewPointState(hit.NewHitNet())).
		AddState(NewPointState(hit.NewHitBackIn())).
		AddState(NewPointState(hit.NewHitOut())).
		AddState(NewPointState(hit.NewToast()))

	return result
}

func AfterHitBackOut() *PointState {
	result := NewPointState(hit.NewHitOut())
	addFromPointStarting(result)

	return result
}

func AfterHitBackNet() *PointState {
	result := NewPointState(hit.NewHitNet())
	addFromPointStarting(result)

	return result
}
