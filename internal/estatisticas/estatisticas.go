package estatisticas

import (
	"github.com/cirobispo/sandbox/internal/comum/golpes"
	"github.com/cirobispo/sandbox/internal/comum/jogos/partida"
)

type predicate func(golpes []golpes.Golpear, indice int) bool

func contarNoSet(set partida.Setting, f predicate) int {
	total := 0
	jogos := set.Games()
	for j := range jogos {
		ponto := jogos[j].Pontos()[j]
		golpes := make([]golpes.Golpear, len(ponto.Golpes()))
		for i := range ponto.Golpes() {
			golpes = append(golpes, ponto.Golpes()[i])
		}

		total += contar(golpes, f)
	}

	return total
}

func contar(golpes []golpes.Golpear, f predicate) int {
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
	EhAce := func(golpes_ []golpes.Golpear, indice int) bool {
		g := golpes_[indice]
		return g.Acao() == golpes.HTAce ||
			(g.Acao() == golpes.HTMiss && indice > 0 && golpes_[indice-1].Acao() == golpes.HTServeIn)
	}

	total := contarNoSet(set, EhAce)

	return total
}

func ContarWinners(set partida.Setting) int {
	EhWinner := func(golpes_ []golpes.Golpear, indice int) bool {
		g := golpes_[indice]
		return g.Acao() == golpes.HTAce ||
			(g.Acao() == golpes.HTMiss && indice > 0)
	}

	return contarNoSet(set, EhWinner)
}
