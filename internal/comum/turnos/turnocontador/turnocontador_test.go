package turnocontador

import (
	"testing"

	"github.com/cirobispo/sandbox/internal/comum/turnos"
	"github.com/cirobispo/sandbox/internal/comum/turnos/turno"
)

type testItem struct {
	turns     int
	startSide turnos.Lado
}

func runTest(test testItem, t *testing.T) {
	obj := turno.New(turno.DefinindoLado(test.startSide)).Decorar(Decorador())
	obj.AdicionarAntesDeMudarTurno(func(ldt turnos.Lado) {
		t.Logf("Lado antes: %v", ldt)
	})
	obj.AdicionarDepoisDeMudarTurno(func(ldt turnos.Lado) {
		t.Logf("Lado depois: %v", ldt)
	})

	for a := test.turns; a > 0; a-- {
		obj.Execute()
	}

	if value, achou := Contar(obj); !achou {
		t.Errorf("turn is \"%v\" turns and should be \"%v\"", value, test.turns)
	}

	clonado := obj.Clonar(obj.LadoInicial())
	chaves := turno.Chaves(obj)
	for idx := range chaves {
		_, achou_lado := obtemLado(clonado, chaves[idx])
		_, achou_contador := obtemContador(clonado)
		if !achou_lado && !achou_contador {
			t.Errorf("clonagem não funcionou adequadamente.")
		}
	}
}

func obtemLado(t *turno.Turno, chave string) (turnos.Lado, bool) {
	if chave != "Turno_LadoInicial" && chave != "Turno_LadoCorrente" {
		return -1, false
	}

	if chave == "Turno_LadoCorrente" {
		return t.LadoCorrente(), true
	}
	return t.LadoInicial(), true
}

func obtemContador(t *turno.Turno) (int, bool) {
	return Contar(t)
}

func Test10CrancksSideA(t *testing.T) {
	runTest(testItem{10, turnos.LadoA}, t)
}

func Test10CrancksSideB(t *testing.T) {
	runTest(testItem{10, turnos.LadoB}, t)
}
