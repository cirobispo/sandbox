package estatisticas

type predicate func(indice int) bool

// func contar(hits []TipoAcaoGolpe, f predicate) int {
// 	tamanho := len(hits)
// 	total := 0

// 	for i := range tamanho {
// 		if f(i) {
// 			total++
// 		}
// 	}

// 	return total
// }

// func ContarAces(hits []TipoAcaoGolpe) int {
// 	EhAce := func(indice int) bool {
// 		g := hits[indice]
// 		return g.Tipo() == HTAce ||
// 			(g.Tipo() == HTMiss && indice > 0 && hits[indice-1].Tipo() == HTServeIn)
// 	}

// 	return contar(hits, EhAce)
// }

// func ContarWinners(hits []TipoAcaoGolpe) int {
// 	EhWinner := func(indice int) bool {
// 		g := hits[indice]
// 		return g.Tipo() == HTAce ||
// 			(g.Tipo() == HTMiss && indice > 0)
// 	}

// 	return contar(hits, EhWinner)
// }
