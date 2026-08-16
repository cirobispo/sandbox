package turnotemporizador

import (
	"time"

	"github.com/cirobispo/sandbox/internal/comum/turnos"
	"github.com/cirobispo/sandbox/internal/comum/turnos/turno"
)

type TurnoTemporizador interface {
	turnos.Turning
	Duracao(t *turno.Turno) int
}

func Decorador() turno.ParamDecoratorOption {
	return func(t *turno.Turno) {
		turno.AdicionarDados(t, "Temporizador_Inicio", turno.NewMapData(time.Now(), func() any { return time.Now() }))
		turno.AdicionarDados(t, "Temporizador_Duracao", turno.NewMapData(time.Since(time.Now()), func() any { return time.Since(time.Now()) }))

		t.AdicionarDepoisDeMudarTurno(func(st turnos.Lado) {
			start, _ := turno.ObterDados[time.Time](t, "Temporizador_Inicio")
			duration := time.Since(start)

			turno.AtualizarDados(t, "Temporizador_Duracao", duration)
		})
	}
}

func Duracao(t *turno.Turno) time.Duration {
	duration, _ := turno.ObterDados[time.Duration](t, "timedTurning_duration")
	return duration
}
