package turnocontador

import (
	"github.com/cirobispo/sandbox/internal/comum/turnos"
	"github.com/cirobispo/sandbox/internal/comum/turnos/turno"
)

type TurnoContador interface {
	turnos.Turning
	Contar(t *turno.Turno) int
}

func ComOutroTurno(t *turno.Turno) func(t *turno.Turno) {
	return func(t *turno.Turno) {
		turno.AdicionarDados(t, "CountedTurning_count", turno.NewMapData(0, func() any { return 0 }))

		t.AdicionarDepoisDeMudarTurno(func(st turnos.LadoDoTurno) {
			value, _ := turno.ObterDados[int](t, "CountedTurning_count")
			value++
			turno.AtualizarDados(t, "CountedTurning_count", value)
		})
	}
}

func Contar(t *turno.Turno) int {
	result, _ := turno.ObterDados[int](t, "CountedTurning_count")
	return result

}
