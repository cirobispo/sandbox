package turnocontador

import (
	"testing"

	"github.com/cirobispo/sandbox/internal/comum/turnos"
	"github.com/cirobispo/sandbox/internal/comum/turnos/turno"
)

type testItem struct {
	turns     int
	startSide turnos.LadoDoTurno
}

func runTest(test testItem, t *testing.T) {
	obj := turno.New(ComOutroTurno(turno.New(turno.DefinindoLado(test.startSide))))
	obj.AdicionarAntesDeMudarTurno(func(ldt turnos.LadoDoTurno) {
		t.Logf("Lado antes: %v", ldt)
	})
	obj.AdicionarDepoisDeMudarTurno(func(ldt turnos.LadoDoTurno) {
		t.Logf("Lado depois: %v", ldt)
	})

	for a := test.turns; a > 0; a-- {
		obj.Execute()
	}

	if value := Contar(obj); value != test.turns {
		t.Errorf("turn is \"%v\" turns and should be \"%v\"", value, test.turns)
	}
}

func Test10CrancksSideA(t *testing.T) {
	runTest(testItem{10, turnos.LTA}, t)
}

func Test10CrancksSideB(t *testing.T) {
	runTest(testItem{10, turnos.LTB}, t)
}
