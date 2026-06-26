package golpes

/**/

type Item struct {
	// ponto  ponto.Pontuacao
	aces   int
	winner int
}

/**
func TestContasAces(tt *testing.T) {
	items := []Item{
		// _ServicoLETs(novoPonto( turnos.LTA)),
		// _AceUmGolpe(novoPonto( turnos.LTA)),
		// _AceDoisGolpes(novoPonto( turnos.LTA)),
		// _DuplaFaltaFora(novoPonto( turnos.LTA)),
		// _DuplaFaltaRede(novoPonto( turnos.LTA)),
		// _ServicoRetornoRede(novoPonto( turnos.LTA)),
		// _ServicoRetornoFora(novoPonto( turnos.LTA)),
		// _WinnerUmGolpeLC(novoPonto( turnos.LTA)),
		// _WinnerUmGolpeLO(novoPonto( turnos.LTA)),
		// _WinnerDoisGolpesLC(novoPonto( turnos.LTA)),
		// _WinnerDoisGolpesLO(novoPonto( turnos.LTA)),
	}

	GolpesAcesWinners := func(item Item) ([]TipoAcaoGolpe, int, int) {
		return nil, item.aces, item.winner
	}

	for i, _ := range items {
		it := items[i]
		golpes, aces, winners := GolpesAcesWinners(it)
		totalAces, totalWinners := ContarAces(golpes), ContarWinners(golpes)
		if totalAces != aces || totalWinners != winners {
			tt.Errorf("\nPonto deveria conter %v ace e %v winner e contem %v e %v. ", aces, winners, totalAces, totalWinners)
		}
	}
}

func novoPonto(side turnos.LadoDoTurno) *ponto.Ponto {
	ctt := turno.New(turnocontador.ComOutroTurno(turno.New((turnotemporizador.ComOutroTurno(turno.New(turno.DefinindoLado(side)))))))

	result := ponto.New(ctt)
	return &result
}

/**
func _AceUmGolpe(ponto *ponto.Ponto) Item {
	ponto.AdicionarGolpe(golpe.NewAce())
	return Item{ponto, 1, 1}
}

/**
func _AceDoisGolpes(ponto *ponto.Ponto) Item {
	ponto.AdicionarGolpe(golpe.NewServicoDentro())
	ponto.AdicionarGolpe(golpe.NewNaoTocou())
	return Item{ponto, 1, 1}
}

func _DuplaFaltaFora(ponto *ponto.Ponto) Item {
	ponto.AdicionarGolpe(golpe.NewServicoFora())
	ponto.AdicionarGolpe(golpe.NewServicoFora())
	return Item{ponto, 0, 0}
}

func _DuplaFaltaRede(ponto *ponto.Ponto) Item {
	ponto.AdicionarGolpe(golpe.NewServicoNaRede())
	ponto.AdicionarGolpe(golpe.NewServicoNaRede())
	return Item{ponto, 0, 0}
}

func _ServicoRetornoRede(ponto *ponto.Ponto) Item {
	ponto.AdicionarGolpe(golpe.NewServicoDentro())
	ponto.AdicionarGolpe(golpe.NewRetornoNaRede())
	return Item{ponto, 0, 0}
}

func _ServicoRetornoFora(ponto *ponto.Ponto) Item {
	ponto.AdicionarGolpe(golpe.NewServicoDentro())
	ponto.AdicionarGolpe(golpe.NewRetornoFora())
	return Item{ponto, 0, 0}
}

func _ServicoLETs(ponto *ponto.Ponto) Item {
	ponto.AdicionarGolpe(golpe.NewLET())
	ponto.AdicionarGolpe(golpe.NewLET())
	ponto.AdicionarGolpe(golpe.NewLET())
	return Item{ponto, 0, 0}
}

func _WinnerUmGolpeLC(ponto *ponto.Ponto) Item {
	ponto.AdicionarGolpe(golpe.NewServicoDentro())
	ponto.AdicionarGolpe(golpe.NewRetornoDentro())
	ponto.AdicionarGolpe(golpe.NewWinner())
	return Item{ponto, 0, 1}
}

func _WinnerUmGolpeLO(ponto *ponto.Ponto) Item {
	ponto.AdicionarGolpe(golpe.NewServicoDentro())
	ponto.AdicionarGolpe(golpe.NewRetornoDentro())
	ponto.AdicionarGolpe(golpe.NewDevolveuDentro())
	ponto.AdicionarGolpe(golpe.NewWinner())
	return Item{ponto, 0, 1}
}

func _WinnerDoisGolpesLC(ponto *ponto.Ponto) Item {
	ponto.AdicionarGolpe(golpe.NewServicoDentro())
	ponto.AdicionarGolpe(golpe.NewRetornoDentro())
	ponto.AdicionarGolpe(golpe.NewDevolveuDentro())
	ponto.AdicionarGolpe(golpe.NewNaoTocou())
	return Item{ponto, 0, 1}
}

func _WinnerDoisGolpesLO(ponto *ponto.Ponto) Item {
	ponto.AdicionarGolpe(golpe.NewServicoDentro())
	ponto.AdicionarGolpe(golpe.NewRetornoDentro())
	ponto.AdicionarGolpe(golpe.NewDevolveuDentro())
	ponto.AdicionarGolpe(golpe.NewDevolveuDentro())
	ponto.AdicionarGolpe(golpe.NewNaoTocou())
	return Item{ponto, 0, 1}
}

/**/
