package turnocontador

import (
	"github.com/cirobispo/sandbox/internal/comum/turnos"
	"github.com/cirobispo/sandbox/internal/comum/turnos/turno"
)

type TurnoContador interface {
	turnos.Turning
	Contar(t *turno.Turno) int
}

func Decorador() turno.ParamDecoratorOption {
	return func(t *turno.Turno) {
		turno.AdicionarDados(t, "Contador_Valor", turno.NewMapData(0, func() any { return 0 }))

		t.AdicionarDepoisDeMudarTurno(func(st turnos.Lado) {
			value, _ := turno.ObterDados[int](t, "Contador_Valor")
			value++
			turno.AtualizarDados(t, "Contador_Valor", value)
		})
	}
}

func Contar(t *turno.Turno) int {
	result, _ := turno.ObterDados[int](t, "Contador_Valor")
	return result
}
