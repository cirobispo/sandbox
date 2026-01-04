package turn

import (
	"fmt"

	"github.com/cirobispo/sandbox/internal/common/turning"
)

type Turn struct {
	data map[string]any

	onBeforeChangeEvent []turning.OnChange
	onAfterChangeEvent  []turning.OnChange
}

func (t *Turn) Execute() {
	t.executeOnChange(t.onBeforeChangeEvent)

	currentSide, _ := GetData[turning.SideTurn](t, "currentSide")

	if currentSide > turning.STA {
		currentSide = -1
	}

	currentSide++
	AddData(t, "currentSide", currentSide)

	t.executeOnChange(t.onAfterChangeEvent)
}

func (t *Turn) GetAllData() map[string]any {
	return t.data
}

func (t Turn) StartSide() turning.SideTurn {
	result, _ := GetData[turning.SideTurn](&t, "startSide")
	return result
}

func (t Turn) CurrentSide() turning.SideTurn {
	result, _ := GetData[turning.SideTurn](&t, "currentSide")
	return result
}

func (t *Turn) AddBeforeChangeEvent(event turning.OnChange) {
	t.onBeforeChangeEvent = append(t.onBeforeChangeEvent, event)
}

func (t *Turn) AddAfterChangeEvent(event turning.OnChange) {
	t.onAfterChangeEvent = append(t.onAfterChangeEvent, event)
}

func (t *Turn) executeOnChange(list []turning.OnChange) {
	currentSide, _ := GetData[turning.SideTurn](t, "currentSide")
	for i := range list {
		event := list[i]
		event(currentSide)
	}
}

func New(start turning.SideTurn) *Turn {
	result := &Turn{
		data:                make(map[string]any),
		onBeforeChangeEvent: make([]turning.OnChange, 0),
		onAfterChangeEvent:  make([]turning.OnChange, 0),
	}

	AddData(result, "startSide", start)
	AddData(result, "currentSide", start)

	return result
}

func AddData[V any](t *Turn, id string, data V) (bool, error) {
	_, f := t.data[id]
	if !f {
		t.GetAllData()[id] = data
		return true, nil
	}

	return false, fmt.Errorf("%s exists.", id)
}

func UpdateData[V any](t *Turn, id string, data V) (bool, error) {
	_, f := t.data[id]
	if f {
		t.GetAllData()[id] = data
		return true, nil
	}

	return false, fmt.Errorf("%s not found.", id)
}

func GetData[V any](t *Turn, id string) (V, bool) {
	data, found := t.GetAllData()[id]

	var result V
	if found {
		result = data.(V)
	}

	return result, found
}
