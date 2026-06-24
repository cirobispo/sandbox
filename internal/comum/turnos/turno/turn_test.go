package turno

import (
	"testing"

	"github.com/cirobispo/sandbox/internal/comum/turnos"
)

type testItem struct {
	turns     int
	startSide turnos.LadoDoTurno
}

func inverse(s turnos.LadoDoTurno) turnos.LadoDoTurno {
	if s == turnos.LTA {
		return turnos.LTB
	}
	return turnos.LTA
}

func runTest(test testItem, t *testing.T) {
	obj := New(DefinindoLado(test.startSide))
	obj.AdicionarAntesDeMudarTurno(func(ldt turnos.LadoDoTurno) {
		t.Logf("Lado antes: %v", ldt)
	})
	obj.AdicionarDepoisDeMudarTurno(func(ldt turnos.LadoDoTurno) {
		t.Logf("Lado depois: %v", ldt)
	})

	endSide := test.startSide
	if test.turns%2 != 0 {
		endSide = inverse(endSide)
	}

	for test.turns > 0 {
		obj.Execute()
		test.turns--
	}

	if obj.LadoCorrente() != endSide {
		t.Errorf("turn is \"%s\" and should be \"%s\"", obj.LadoCorrente(), endSide)
	}
}

func Test10Crancks(t *testing.T) {
	runTest(testItem{10, turnos.LTA}, t)
}
