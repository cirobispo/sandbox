package utilitario

import (
	"github.com/cirobispo/sandbox/internal/comum/pontos"
	"github.com/cirobispo/sandbox/internal/comum/pontos/golpes"
)

func ExisteDuplaFalta(hits *[]golpes.Golpe) bool {
	tamanho := len(*hits)
	if tamanho < 2 {
		return false
	}

	ultimoGolpe := (*hits)[tamanho-1]
	FoiFalta := func(hit golpes.Golpe) bool {
		return hit.Tipo() == golpes.HTFootFault ||
			hit.Tipo() == golpes.HTServeNet ||
			hit.Tipo() == golpes.HTServeOut
	}
	if !FoiFalta(ultimoGolpe) {
		return false
	}

	count := 1
	result := false
	for i := tamanho - 2; i >= 0; i-- {
		hit := (*hits)[i]
		if FoiFalta(hit) {
			count++
			if count > 1 {
				result = true
				break
			}
		}
	}
	return result
}

func LadoDoGolpeParaLadoDoPonto(s golpes.LadoDoGolpe) pontos.LadoDoPonto {
	switch s {
	case golpes.HTDSameSide:
		return pontos.LPCorrente
	case golpes.HTDOppositeSide:
		return pontos.LPOposto
	default:
		return pontos.LPNulo
	}
}
