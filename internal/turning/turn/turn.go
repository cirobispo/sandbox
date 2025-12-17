package turn

import (
	turning "github.com/cirobispo/sandbox/internal/turning"
)

type OnChange func(turning.SideTurn)

type Turn struct {
	currentSide turning.SideTurn
	startSide   turning.SideTurn
	data        map[string]any

	onBeforeChangeEvent []OnChange
	onAfterChangeEvent  []OnChange
}

func (t *Turn) Execute() {
	t.executeOnChange(t.onBeforeChangeEvent)

	if t.currentSide > turning.STA {
		t.currentSide = -1
	}

	t.currentSide++
	t.executeOnChange(t.onAfterChangeEvent)
}

func (t *Turn) AddBeforeChangeEvent(event OnChange) {
	t.onBeforeChangeEvent = append(t.onBeforeChangeEvent, event)
}

func (t *Turn) AddAfterChangeEvent(event OnChange) {
	t.onAfterChangeEvent = append(t.onAfterChangeEvent, event)
}

func (t *Turn) executeOnChange(list []OnChange) {
	for i := 0; i < len(list); i++ {
		event := list[i]
		event(t.currentSide)
	}
}

func AddData[V any](t *Turn, id string, data V) {
	t.data[id] = data
}

func GetData[V any](t *Turn, id string) V {
	data, found := t.data[id]

	var result V
	if found {
		result = data.(V)
	}

	return result
}

func New(start turning.SideTurn) Turn {
	return Turn{
		startSide:           start,
		currentSide:         start,
		data:                make(map[string]any),
		onBeforeChangeEvent: make([]OnChange, 0),
		onAfterChangeEvent:  make([]OnChange, 0),
	}
}
