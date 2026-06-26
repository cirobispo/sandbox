package estatisticas

import "github.com/cirobispo/sandbox/internal/comum/golpes"

type predicate func(indice int) bool

func contar(hits []golpes.TipoAcaoGolpe, f predicate) int {
	tamanho := len(hits)
	total := 0

	for i := range tamanho {
		if f(i) {
			total++
		}
	}

	return total
}

func ContarAces(hits []golpes.TipoAcaoGolpe) int {
	EhAce := func(indice int) bool {
		g := hits[indice]
		return g.Tipo() == golpes.HTAce ||
			(g.Tipo() == golpes.HTMiss && indice > 0 && hits[indice-1].Tipo() == golpes.HTServeIn)
	}

	return contar(hits, EhAce)
}

func ContarWinners(hits []golpes.TipoAcaoGolpe) int {
	EhWinner := func(indice int) bool {
		g := hits[indice]
		return g.Tipo() == golpes.HTAce ||
			(g.Tipo() == golpes.HTMiss && indice > 0)
	}

	return contar(hits, EhWinner)
}
