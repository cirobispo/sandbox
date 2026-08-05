package turnotemporizador

import (
	"testing"
	"time"

	"github.com/cirobispo/sandbox/internal/comum/turnos"
	"github.com/cirobispo/sandbox/internal/comum/turnos/turno"
)

type testItem struct {
	turns     int
	startSide turnos.Lado
}

func runTest(test testItem, t *testing.T) {
	obj := turno.New(turno.DefinindoLado(test.startSide)).Decorator(ComOutroTurno())
	begin := time.Now()
	for a := test.turns; a > 0; a-- {
		obj.Execute()
	}
	diff := time.Since(begin).Round(time.Millisecond)

	if value := Duracao(obj).Round(time.Millisecond); value != diff {
		t.Errorf("turn is \"%v\" turns and should be \"%v\"", value, diff)
	}
}

func Test10CrancksSideA(t *testing.T) {
	runTest(testItem{100, turnos.LadoA}, t)
}

func Test10CrancksSideB(t *testing.T) {
	runTest(testItem{100, turnos.LadoB}, t)
}
