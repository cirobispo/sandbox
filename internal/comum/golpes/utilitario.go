package golpes

func ExisteDuplaFalta(gs []TipoAcaoGolpe) bool {
	FoiFalta := func(hit TipoAcaoGolpe) bool {
		return hit.Tipo() == HTFootFault || hit.Tipo() == HTServeNet || hit.Tipo() == HTServeOut
	}

	tamanho := len(gs)
	if tamanho < 2 || !FoiFalta(gs[tamanho-1]) {
		return false
	}

	count := 1
	result := false
	for i := tamanho - 2; i >= 0; i-- {
		g := gs[i]
		if FoiFalta(g) {
			count++
		}
	}
	return result
}
