package placares

import (
	"testing"

	"github.com/cirobispo/sandbox/internal/comum/turnos"
)

func TestQualLadoSacar(t *testing.T) {
	ladoInicialPartida := turnos.LadoA

	for s := range 1 {
		ladoInicialSet := QualLadoSacar(ladoInicialPartida, uint(s), false)
		for j := range 6 {
			ladoInicialJogo := QualLadoSacar(ladoInicialSet, uint(j), false)
			for p := range 4 {
				ladoSaque := QualLadoSacar(turnos.LadoA, uint(p), false)
				t.Logf("Partida: %v, Set (%d): %v, Jogo (%d) %v, Ponto: %d, Saque: %v",
					ladoInicialPartida,
					s+1, ladoInicialSet,
					j+1, ladoInicialJogo,
					p+1,
					ladoSaque,
				)
			}
			t.Log()
		}
	}

	for p := range 11 {
		ladoSaque := QualLadoSacar(turnos.LadoA, uint(p), true)
		t.Logf("Partida: %v, TieBreak (%d): %v, Jogo (%d) %v, Ponto: %02d, Saque: %v",
			ladoInicialPartida,
			1, ladoInicialPartida,
			1, ladoInicialPartida,
			p+1,
			ladoSaque,
		)
	}
}
