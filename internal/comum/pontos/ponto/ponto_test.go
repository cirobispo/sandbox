package ponto

import (
	"testing"

	"github.com/cirobispo/sandbox/internal/comum/golpes"
	"github.com/cirobispo/sandbox/internal/comum/golpes/golpe"
	"github.com/cirobispo/sandbox/internal/comum/pontos"
	"github.com/cirobispo/sandbox/internal/comum/turnos"
	"github.com/cirobispo/sandbox/internal/comum/turnos/turno"
	"github.com/cirobispo/sandbox/internal/comum/turnos/turnocontador"
	"github.com/cirobispo/sandbox/internal/comum/turnos/turnotemporizador"
)

type Item struct {
	ponto       pontos.Pontuando
	ladoDoPonto pontos.LadoDoPonto
}

func TestTodosOsPontos(tt *testing.T) {
	items := []Item{
		_ServicoLETs(novoPonto(tt, turnos.LTA)),
		_AceUmGolpe(novoPonto(tt, turnos.LTA)),
		_AceDoisGolpes(novoPonto(tt, turnos.LTA)),
		_DuplaFaltaFora(novoPonto(tt, turnos.LTA)),
		_DuplaFaltaRede(novoPonto(tt, turnos.LTA)),
		_ServicoRetornoRede(novoPonto(tt, turnos.LTA)),
		_ServicoRetornoFora(novoPonto(tt, turnos.LTA)),
		_WinnerUmGolpeLC(novoPonto(tt, turnos.LTA)),
		_WinnerUmGolpeLO(novoPonto(tt, turnos.LTA)),
		_WinnerDoisGolpesLC(novoPonto(tt, turnos.LTA)),
		_WinnerDoisGolpesLO(novoPonto(tt, turnos.LTA)),
	}

	for i, _ := range items {
		it := items[i]
		if !it.ponto.Terminado() {
			tt.Logf("\nPonto não foi encerrado! ")
		}

		if it.ladoDoPonto != it.ponto.LadoDoPonto() {
			tt.Errorf("\nPonto deveria ser: %v mas resultou %v. ", it.ladoDoPonto, it.ponto.LadoDoPonto())
			// mostrarPontos(tt, it.ponto)
		}
		// tt.Log()
	}
}

func mostrarPontos(tt *testing.T, ponto *Ponto) {
	golpes := ponto.Golpes()
	tt.Logf("Golpes executados: %v, trocas feitas: %v", len(golpes), turnocontador.Contar(ponto.ladoDaBola))
	for i, _ := range golpes {
		golpe := golpes[i]
		tt.Logf("Golpe: %v ", golpe.Tipo())
	}
}

func novoPonto(tt *testing.T, side turnos.LadoDoTurno) *Ponto {
	ctt := turno.New(turnocontador.ComOutroTurno(turno.New((turnotemporizador.ComOutroTurno(turno.New(turno.DefinindoLado(side)))))))
	// ctt.AdicionarAntesDeMudarTurno(func(ldt turnos.LadoDoTurno) {
	// 	tt.Log("Lado antes: ", ldt)
	// })
	// ctt.AdicionarDepoisDeMudarTurno(func(ldt turnos.LadoDoTurno) {
	// 	tt.Log("Lado depois: ", ldt)
	// })
	result := New(ctt)
	result.AdicionarEventoAoAdicionarGolpe(func(tipoDoGolpe golpes.TipoDoGolpe, terminado bool) {
		if terminado {
			tt.Logf("Golpe: %v e encerrou o ponto com %v golpe(s) e %v troca(s)", tipoDoGolpe, len(result.Golpes()), turnocontador.Contar(result.ladoDaBola))
			tt.Log()
			return
		}
		tt.Logf("Golpe: %v", tipoDoGolpe)
	})
	return &result
}

func _AceUmGolpe(ponto *Ponto) Item {
	ponto.AdicionarGolpe(golpe.NewAce())
	return Item{ponto, pontos.LPCorrente}
}

func _AceDoisGolpes(ponto *Ponto) Item {
	ponto.AdicionarGolpe(golpe.NewServicoDentro())
	ponto.AdicionarGolpe(golpe.NewNaoTocou())
	return Item{ponto, pontos.LPOposto}
}

func _DuplaFaltaFora(ponto *Ponto) Item {
	ponto.AdicionarGolpe(golpe.NewServicoFora())
	ponto.AdicionarGolpe(golpe.NewServicoFora())
	return Item{ponto, pontos.LPOposto}
}

func _DuplaFaltaRede(ponto *Ponto) Item {
	ponto.AdicionarGolpe(golpe.NewServicoNaRede())
	ponto.AdicionarGolpe(golpe.NewServicoNaRede())
	return Item{ponto, pontos.LPOposto}
}

func _ServicoRetornoRede(ponto *Ponto) Item {
	ponto.AdicionarGolpe(golpe.NewServicoDentro())
	ponto.AdicionarGolpe(golpe.NewRetornoNaRede())
	return Item{ponto, pontos.LPOposto}
}

func _ServicoRetornoFora(ponto *Ponto) Item {
	ponto.AdicionarGolpe(golpe.NewServicoDentro())
	ponto.AdicionarGolpe(golpe.NewRetornoFora())
	return Item{ponto, pontos.LPOposto}
}

func _ServicoLETs(ponto *Ponto) Item {
	ponto.AdicionarGolpe(golpe.NewLET())
	ponto.AdicionarGolpe(golpe.NewLET())
	ponto.AdicionarGolpe(golpe.NewLET())
	return Item{ponto, pontos.LPNulo}
}

func _WinnerUmGolpeLC(ponto *Ponto) Item {
	ponto.AdicionarGolpe(golpe.NewServicoDentro())
	ponto.AdicionarGolpe(golpe.NewRetornoDentro())
	ponto.AdicionarGolpe(golpe.NewWinner())
	return Item{ponto, pontos.LPCorrente}
}

func _WinnerUmGolpeLO(ponto *Ponto) Item {
	ponto.AdicionarGolpe(golpe.NewServicoDentro())
	ponto.AdicionarGolpe(golpe.NewRetornoDentro())
	ponto.AdicionarGolpe(golpe.NewDevolveuDentro())
	ponto.AdicionarGolpe(golpe.NewWinner())
	return Item{ponto, pontos.LPCorrente}
}

func _WinnerDoisGolpesLC(ponto *Ponto) Item {
	ponto.AdicionarGolpe(golpe.NewServicoDentro())
	ponto.AdicionarGolpe(golpe.NewRetornoDentro())
	ponto.AdicionarGolpe(golpe.NewDevolveuDentro())
	ponto.AdicionarGolpe(golpe.NewNaoTocou())
	return Item{ponto, pontos.LPOposto}
}

func _WinnerDoisGolpesLO(ponto *Ponto) Item {
	ponto.AdicionarGolpe(golpe.NewServicoDentro())
	ponto.AdicionarGolpe(golpe.NewRetornoDentro())
	ponto.AdicionarGolpe(golpe.NewDevolveuDentro())
	ponto.AdicionarGolpe(golpe.NewDevolveuDentro())
	ponto.AdicionarGolpe(golpe.NewNaoTocou())
	return Item{ponto, pontos.LPOposto}
}
