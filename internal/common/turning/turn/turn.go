package turn

import (
	"errors"

	"github.com/cirobispo/sandbox/internal/common/turning"
)

type mapData struct {
	value any
	reset func() any
}

func NewMapData[V any](value V, callBack func() any) mapData {
	return mapData{value: value, reset: callBack}
}

type ParamOption func(t *Turn)

type Turn struct {
	data map[string]mapData

	onBeforeChangeEvent []turning.OnChange
	onAfterChangeEvent  []turning.OnChange
}

func (t *Turn) AddOnBeforeChange(callback turning.OnChange) {
	t.onBeforeChangeEvent = append(t.onBeforeChangeEvent, callback)
}

func (t *Turn) AddOnAfterChange(callback turning.OnChange) {
	t.onAfterChangeEvent = append(t.onAfterChangeEvent, callback)
}

func (t *Turn) Execute() {
	t.executeOnChange(t.onBeforeChangeEvent)

	currentSide, _ := GetData[turning.TurningSide](t, "Turn_currentSide")

	if currentSide > turning.TSA {
		currentSide = -1
	}

	currentSide++
	UpdateData(t, "Turn_currentSide", currentSide)

	t.executeOnChange(t.onAfterChangeEvent)
}

func (t Turn) StartSide() turning.TurningSide {
	result, _ := GetData[turning.TurningSide](&t, "Turn_startSide")
	return result
}

func (t Turn) CurrentSide() turning.TurningSide {
	result, _ := GetData[turning.TurningSide](&t, "Turn_currentSide")
	return result
}

func (t Turn) Clone(start turning.TurningSide) *Turn {
	result := New(WithTurningSide(start))
	copy(result.onBeforeChangeEvent, t.onBeforeChangeEvent)
	copy(result.onAfterChangeEvent, t.onAfterChangeEvent)

	dataCount, dataAdded := len(t.data)-2, 0

	for k, v := range t.data {
		if f, _ := AddData(result, k, v); f {
			UpdateData(result, k, v.reset())
			dataAdded++
		}
	}

	if dataAdded != dataCount {
		panic(errors.New("data added to clone turn does not match."))
	}

	return result
}

func (t Turn) executeOnChange(list []turning.OnChange) {
	currentSide, _ := GetData[turning.TurningSide](&t, "Turn_currentSide")
	for i := range list {
		event := list[i]
		event(currentSide)
	}
}

func WithTurningSide(start turning.TurningSide) func(t *Turn) {
	return func(t *Turn) {
		startSide := NewMapData(start, func() any { return start })
		currentSide := NewMapData(start, func() any { return start })
		AddData(t, "Turn_startSide", startSide)
		AddData(t, "Turn_currentSide", currentSide)
	}
}

func New(param func(t *Turn)) *Turn {
	result := &Turn{
		data:                make(map[string]mapData),
		onBeforeChangeEvent: make([]turning.OnChange, 0),
		onAfterChangeEvent:  make([]turning.OnChange, 0),
	}

	if param != nil {
		param(result)
	}
	return result
}

func AddData(t *Turn, id string, data mapData) (bool, error) {
	_, f := t.data[id]
	if !f {
		t.data[id] = data
		return true, nil
	}

	return false, errors.New("id not found.")
}

func UpdateData[V any](t *Turn, id string, data V) (bool, error) {
	d, f := t.data[id]
	if f {
		t.data[id] = NewMapData[V](data, d.reset)
		return true, nil
	}

	return false, errors.New("id not found.")
}

func GetData[V any](t *Turn, id string) (V, bool) {
	data, found := t.data[id]

	var result V
	if found {
		result = data.value.(V)
	}

	return result, found
}
