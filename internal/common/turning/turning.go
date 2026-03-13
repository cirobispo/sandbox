package turning

type TurningSide int

const (
	STA TurningSide = 0
	STB TurningSide = 1
)

type Turning interface {
	Execute()
	GetStartSide() TurningSide
	GetCurrentSide() TurningSide
}

type OnChange func(TurningSide)

func (s TurningSide) String() string {
	result := "Side A"
	if s == STB {
		result = "Side B"
	}

	return result
}

// type TurnEventing interface {
// 	AddBeforeChangeEvent(event OnChange)
// 	AddAfterChangeEvent(event OnChange)
// }
