package turning

type SideTurn int

const (
	STA SideTurn = 0
	STB SideTurn = 1
)

type Turning interface {
	Execute()
	GetStartSide() SideTurn
	GetCurrentSide() SideTurn
}

type OnChange func(SideTurn)

func (s SideTurn) String() string {
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
