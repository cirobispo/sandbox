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

func ComOutroTurno(t *turno.Turno) func(t *turno.Turno) {
	return func(t *turno.Turno) {
		turno.AdicionarDados(t, "timedTurning_start", turno.NewMapData(time.Now(), func() any { return time.Now() }))
		turno.AdicionarDados(t, "timedTurning_duration", turno.NewMapData(time.Since(time.Now()), func() any { return time.Since(time.Now()) }))

		t.AdicionarDepoisDeMudarTurno(func(st turnos.LadoDoTurno) {
			start, _ := turno.ObterDados[time.Time](t, "timedTurning_start")
			duration := time.Since(start)

			turno.AtualizarDados(t, "timedTurning_duration", duration)
		})
	}
}

func Duracao(t *turno.Turno) time.Duration {
	duration, _ := turno.ObterDados[time.Duration](t, "timedTurning_duration")
	return duration
}
