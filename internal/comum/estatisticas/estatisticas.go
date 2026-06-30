package estatisticas

import (
	"github.com/cirobispo/sandbox/internal/comum/golpes"
	"github.com/cirobispo/sandbox/internal/comum/jogos/partida"
)

type predicate func(golpes []golpes.TipoAcaoGolpe, indice int) bool

func contarNoSet(set partida.Setting, f predicate) int {
	total := 0
	jogos := set.Games()
	for j := range jogos {
		ponto := jogos[j].Pontos()[j]
		total += contar(ponto.Golpes(), f)
	}

	return total
}

func contar(golpes []golpes.TipoAcaoGolpe, f predicate) int {
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
	EhAce := func(golpes_ []golpes.TipoAcaoGolpe, indice int) bool {
		g := golpes_[indice]
		return g.Tipo() == golpes.HTAce ||
			(g.Tipo() == golpes.HTMiss && indice > 0 && golpes_[indice-1].Tipo() == golpes.HTServeIn)
	}

	total := contarNoSet(set, EhAce)

	// Jogos := set.Games()
	// for j := range Jogos {
	// 	for p := range Jogos[j].Pontos() {
	// 		ponto := Jogos[j].Pontos()[p]
	// 		golpes_ := ponto.Golpes()
	// 		total += contar(golpes_, EhAce)
	// 	}
	// }

	return total
}

func ContarAcess(golpes_ []golpes.TipoAcaoGolpe) int {
	EhAce := func(golpes_ []golpes.TipoAcaoGolpe, indice int) bool {
		g := golpes_[indice]
		return g.Tipo() == golpes.HTAce ||
			(g.Tipo() == golpes.HTMiss && indice > 0 && golpes_[indice-1].Tipo() == golpes.HTServeIn)
	}

	return contar(golpes_, EhAce)
}

func ContarWinners(golpes_ []golpes.TipoAcaoGolpe) int {
	EhWinner := func(golpes_ []golpes.TipoAcaoGolpe, indice int) bool {
		g := golpes_[indice]
		return g.Tipo() == golpes.HTAce ||
			(g.Tipo() == golpes.HTMiss && indice > 0)
	}

	return contar(golpes_, EhWinner)
}
