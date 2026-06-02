package ponto

import (
	"github.com/cirobispo/sandbox/internal/common/pontos/golpes/golpe"
)

func PointStarting() *PointState {
	result := NewPointState(nil).
		AddState(NewPointState(golpe.NewAce())).
		AddState(NewPointState(golpe.NewNetTouch())).
		AddState(NewPointState(golpe.NewServeNet())).
		AddState(NewPointState(golpe.NewServeLet())).
		AddState(NewPointState(golpe.NewServeIn())).
		AddState(NewPointState(golpe.NewServeOut()))

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
	result := NewPointState(golpe.NewAce())
	addFromPointStarting(result)
	return result
}

func AfterServeIn() *PointState {
	result := NewPointState(golpe.NewServeIn()).
		AddState(NewPointState(golpe.NewReturnNet())).
		AddState(NewPointState(golpe.NewReturnIn())).
		AddState(NewPointState(golpe.NewReturnOut())).
		AddState(NewPointState(golpe.NewToast()))

	return result
}

func AfterServeOut() *PointState {
	result := NewPointState(golpe.NewServeOut())
	addFromPointStarting(result)
	return result
}

func AfterServeNet() *PointState {
	result := NewPointState(golpe.NewServeNet())
	addFromPointStarting(result)

	return result
}

func AfterServeLet() *PointState {
	result := NewPointState(golpe.NewServeLet())
	addFromPointStarting(result)

	return result
}

func AfterReturnIn() *PointState {
	result := NewPointState(golpe.NewReturnIn()).
		AddState(NewPointState(golpe.NewHitNet())).
		AddState(NewPointState(golpe.NewHitBackIn())).
		AddState(NewPointState(golpe.NewHitOut())).
		AddState(NewPointState(golpe.NewToast()))

	return result
}

func AfterReturnOut() *PointState {
	result := NewPointState(golpe.NewReturnOut())
	addFromPointStarting(result)

	return result
}

func AfterReturnNet() *PointState {
	result := NewPointState(golpe.NewReturnNet())
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
	result := NewPointState(golpe.NewHitBackIn()).
		AddState(NewPointState(golpe.NewHitNet())).
		AddState(NewPointState(golpe.NewHitBackIn())).
		AddState(NewPointState(golpe.NewHitOut())).
		AddState(NewPointState(golpe.NewToast()))

	return result
}

func AfterHitBackOut() *PointState {
	result := NewPointState(golpe.NewHitOut())
	addFromPointStarting(result)

	return result
}

func AfterHitBackNet() *PointState {
	result := NewPointState(golpe.NewHitNet())
	addFromPointStarting(result)

	return result
}
