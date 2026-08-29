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
		t.Errorf("Lado do turno é \"%s\" e deveria ser \"%s\"", obj.LadoCorrente(), endSide)
	}

	clonado := obj.Clonar(obj.LadoInicial())
	for k, v := range obj.dados {
		if mapdata, achou := clonado.dados[k]; !achou || (achou && mapdata.valor != v.valor) {
			t.Errorf("clonagem não funcionou adequadamente. mapData %v diferente", mapdata)
		}
	}
}

func Test10Crancks(t *testing.T) {
	runTest(testItem{10, turnos.LadoA}, t)
}
