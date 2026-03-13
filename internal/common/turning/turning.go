package turning

type TurningSide int

const (
	TSA TurningSide = 0
	TSB TurningSide = 1
)

type Turning interface {
	Execute()
	GetStartSide() TurningSide
	GetCurrentSide() TurningSide
}

type OnChange func(TurningSide)

func (s TurningSide) String() string {
	result := "Side A"
	if s == TSB {
		result = "Side B"
	}

	return result
}

// type TurnEventing interface {
// 	AddBeforeChangeEvent(event OnChange)
// 	AddAfterChangeEvent(event OnChange)
// }
