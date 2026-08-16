package ponto

import (
	"github.com/cirobispo/sandbox/internal/comum/golpes/golpe"
)

func PointStarting() *PointState {
	result := NewPointState(nil).
		AddState(NewPointState(golpe.NewAce())).
		AddState(NewPointState(golpe.NewToqueNaRede())).
		AddState(NewPointState(golpe.NewServicoNaRede())).
		AddState(NewPointState(golpe.NewLET())).
		AddState(NewPointState(golpe.NewServicoDentro())).
		AddState(NewPointState(golpe.NewServicoFora()))

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
	result := NewPointState(golpe.NewServicoDentro()).
		AddState(NewPointState(golpe.NewRetornoNaRede())).
		AddState(NewPointState(golpe.NewRetornoDentro())).
		AddState(NewPointState(golpe.NewRetornoFora())).
		AddState(NewPointState(golpe.NewQueimou()))

	return result
}

func AfterServeOut() *PointState {
	result := NewPointState(golpe.NewServicoFora())
	addFromPointStarting(result)
	return result
}

func AfterServeNet() *PointState {
	result := NewPointState(golpe.NewServicoNaRede())
	addFromPointStarting(result)

	return result
}

func AfterServeLet() *PointState {
	result := NewPointState(golpe.NewLET())
	addFromPointStarting(result)

	return result
}

func AfterReturnIn() *PointState {
	result := NewPointState(golpe.NewRetornoDentro()).
		AddState(NewPointState(golpe.NewDevolveuNaRede())).
		AddState(NewPointState(golpe.NewDevolveuDentro())).
		AddState(NewPointState(golpe.NewDevolveuFora())).
		AddState(NewPointState(golpe.NewQueimou()))

	return result
}

func AfterReturnOut() *PointState {
	result := NewPointState(golpe.NewRetornoFora())
	addFromPointStarting(result)

	return result
}

func AfterReturnNet() *PointState {
	result := NewPointState(golpe.NewRetornoNaRede())
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
	result := NewPointState(golpe.NewDevolveuDentro()).
		AddState(NewPointState(golpe.NewDevolveuNaRede())).
		AddState(NewPointState(golpe.NewDevolveuDentro())).
		AddState(NewPointState(golpe.NewDevolveuFora())).
		AddState(NewPointState(golpe.NewQueimou()))

	return result
}

func AfterHitBackOut() *PointState {
	result := NewPointState(golpe.NewDevolveuFora())
	addFromPointStarting(result)

	return result
}

func AfterHitBackNet() *PointState {
	result := NewPointState(golpe.NewDevolveuNaRede())
	addFromPointStarting(result)

	return result
}
