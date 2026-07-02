package estatisticas

import (
	"github.com/cirobispo/sandbox/internal/comum/golpes"
	"github.com/cirobispo/sandbox/internal/comum/jogos/partida"
)

type predicate func(golpes []golpes.Golpeando, indice int) bool

func contarNoSet(set partida.Setting, f predicate) int {
	total := 0
	jogos := set.Games()
	for j := range jogos {
		ponto := jogos[j].Pontos()[j]
		total += contar(ponto.Golpes(), f)
	}

	return total
}

func contar(golpes []golpes.Golpeando, f predicate) int {
	tamanho := len(golpes)
	total := 0

	for i := range tamanho {
		if f(golpes, i) {
			total++
		}
	}

	return total
}

func ContarAces(set partida.Setting) int {
	EhAce := func(golpes_ []golpes.Golpeando, indice int) bool {
		g := golpes_[indice]
		return g.Tipo() == golpes.HTAce ||
			(g.Tipo() == golpes.HTMiss && indice > 0 && golpes_[indice-1].Tipo() == golpes.HTServeIn)
	}

	total := contarNoSet(set, EhAce)

	return total
}

func ContarWinners(set partida.Setting) int {
	EhWinner := func(golpes_ []golpes.Golpeando, indice int) bool {
		g := golpes_[indice]
		return g.Tipo() == golpes.HTAce ||
			(g.Tipo() == golpes.HTMiss && indice > 0)
	}

	return contarNoSet(set, EhWinner)
}
