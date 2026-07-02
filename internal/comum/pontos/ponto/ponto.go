package ponto

import (
	"errors"

	"github.com/cirobispo/sandbox/internal/comum/golpes"
	"github.com/cirobispo/sandbox/internal/comum/pontos"
	"github.com/cirobispo/sandbox/internal/comum/turnos"
	"github.com/cirobispo/sandbox/internal/comum/turnos/turno"
)

type Ponto struct {
	ladoDaBola              *turno.Turno
	golpes                  []golpes.Golpeando
	ladoDoPonto             pontos.LadoDoPonto
	eventosAoAdicionarGolpe *[]pontos.AoAdicionarGolpe
}

func New(sideControl *turno.Turno) Ponto {
	hit := make([]golpes.Golpeando, 0, 3)
	events := make([]pontos.AoAdicionarGolpe, 0)
	return Ponto{
		ladoDoPonto:             pontos.LPNulo,
		ladoDaBola:              sideControl,
		golpes:                  hit,
		eventosAoAdicionarGolpe: &events,
	}
}

func (p *Ponto) AdicionarEventoAoAdicionarGolpe(ponteiroFnc pontos.AoAdicionarGolpe) {
	*p.eventosAoAdicionarGolpe = append(*p.eventosAoAdicionarGolpe, ponteiroFnc)
}

func (p *Ponto) AdicionarGolpe(g golpes.Golpeando) error {
	if p.ladoDoPonto != pontos.LPNulo {
		return errors.New("Não é possivel adicionar outro golpe com o ponto encerrado.")
	}

	p.golpes = append(p.golpes, g)
	acao := g.Acao()
	if acao == golpes.TACondicional && existeDuplaFalta(p.golpes) {
		acao = golpes.TAEncerrarPLO
	}

	terminado := (acao == golpes.TAEncerrarPLC) || (acao == golpes.TAEncerrarPLO)

	if terminado {
		p.ladoDoPonto = pontos.LPCorrente
		if acao == golpes.TAEncerrarPLO {
			p.ladoDoPonto = pontos.LPOposto
		}
	}

	if (acao == golpes.TAProsseguir) || (acao == golpes.TAEncerrarPLC) {
		p.ladoDaBola.Execute()
	}
	p.executeEventosAoAdicionarGolpe()

	return nil
}

func (p Ponto) Golpes() []golpes.Golpeando {
	result := make([]golpes.Golpeando, len(p.golpes))
	copy(result, p.golpes)

	return result
}

func (p Ponto) Tamanho() int {
	result := len(p.golpes)
	return result
}

func (p Ponto) LadoDaBola() turnos.LadoDoTurno {
	return p.ladoDaBola.LadoCorrente()
}

func (p Ponto) LadoDoPonto() pontos.LadoDoPonto {
	return p.ladoDoPonto
}

func (p Ponto) Terminado() bool {
	return p.ladoDoPonto != pontos.LPNulo
}

func (p Ponto) Clonar() Ponto {
	result := New(p.ladoDaBola.Clonar(p.ladoDaBola.LadoInicial()))
	result.golpes = make([]golpes.Golpeando, len(p.golpes))
	copy(result.golpes, p.golpes)

	*result.eventosAoAdicionarGolpe = make([]pontos.AoAdicionarGolpe, len(*p.eventosAoAdicionarGolpe))
	copy(*result.eventosAoAdicionarGolpe, *p.eventosAoAdicionarGolpe)
	result.ladoDoPonto = p.ladoDoPonto

	return result
}

func (p Ponto) executeEventosAoAdicionarGolpe() {
	if len(*p.eventosAoAdicionarGolpe) > 0 {
		golpe := p.golpes[p.Tamanho()-1]
		tipo, terminado := golpe.Tipo(), (p.ladoDoPonto != pontos.LPNulo)

		for i := range *p.eventosAoAdicionarGolpe {
			event := (*p.eventosAoAdicionarGolpe)[i]
			event(tipo, terminado)
		}
	}
}

func existeDuplaFalta(gs []golpes.Golpeando) bool {
	FoiFalta := func(hit golpes.Golpeando) bool {
		return hit.Tipo() == golpes.HTFootFault || hit.Tipo() == golpes.HTServeNet || hit.Tipo() == golpes.HTServeOut
	}

	tamanho := len(gs)
	if tamanho < 2 || !FoiFalta(gs[tamanho-1]) {
		return false
	}

	count := 1
	for i := tamanho - 2; i >= 0; i-- {
		g := gs[i]
		if FoiFalta(g) {
			count++
		}
	}
	return (count == 2)
}
