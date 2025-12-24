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

// type TurnEventing interface {
// 	AddBeforeChangeEvent(event OnChange)
// 	AddAfterChangeEvent(event OnChange)
// }
