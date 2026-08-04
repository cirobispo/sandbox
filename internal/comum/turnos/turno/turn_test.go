package turno

import (
	"testing"

	"github.com/cirobispo/sandbox/internal/comum/turnos"
)

type testItem struct {
	turns     int
	startSide turnos.Lado
}

func inverse(s turnos.Lado) turnos.Lado {
	if s == turnos.LadoA {
		return turnos.LadoB
	}
	return turnos.LadoA
}

func runTest(test testItem, t *testing.T) {
	obj := New(DefinindoLado(test.startSide))
	obj.AdicionarAntesDeMudarTurno(func(ldt turnos.Lado) {
		t.Logf("Lado antes: %v", ldt)
	})
	obj.AdicionarDepoisDeMudarTurno(func(ldt turnos.Lado) {
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
	runTest(testItem{10, turnos.LadoA}, t)
}
