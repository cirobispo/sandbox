package ponto

import (
	"github.com/cirobispo/sandbox/internal/comum/pontos"
	"github.com/cirobispo/sandbox/internal/comum/pontos/golpes"
	"github.com/cirobispo/sandbox/internal/comum/turnos"
	"github.com/cirobispo/sandbox/internal/comum/turnos/turno"
)

type Ponto struct {
	ladoDaBola               *turno.Turno
	golpes                   *[]golpes.Golpe
	duplasFaltas             int
	ladoDoPonto              pontos.LadoDoPonto
	terminado                bool
	eventosAoPontuarNoPlacar *[]pontos.AoPontuar
}

func New(sideControl *turno.Turno) Ponto {
	hit := make([]golpes.Golpe, 0, 3)
	events := make([]pontos.AoPontuar, 0)
	return Ponto{
		terminado:                false,
		ladoDoPonto:              pontos.LPNulo,
		ladoDaBola:               sideControl,
		duplasFaltas:             0,
		golpes:                   &hit,
		eventosAoPontuarNoPlacar: &events,
	}
}

func (p *Ponto) AdicionarEventoAoPontuar(ponteiroFnc pontos.AoPontuar) {
	*p.eventosAoPontuarNoPlacar = append(*p.eventosAoPontuarNoPlacar, ponteiroFnc)
}

func (p *Ponto) AdicionarGolpe(g golpes.Golpe) {
	if p.terminado {
		return
	}

	*p.golpes = append(*p.golpes, g)
	acao := g.Acao(*p.golpes)
	p.terminado = (acao == golpes.TAEncerrarPLC) || (acao == golpes.TAEncerrarPLO)
	if !p.terminado {
		p.ladoDaBola.Execute()
		return
	}

	if acao == golpes.TAEncerrarPLC {
		p.ladoDoPonto = pontos.LPCorrente
	}

	p.ladoDoPonto = pontos.LPOposto
}

func (p Ponto) Golpes() []golpes.Golpe {
	result := make([]golpes.Golpe, len(*p.golpes))
	copy(result, *p.golpes)

	return result
}

func (p Ponto) Tamanho() int {
	result := len(*p.golpes)
	return result
}

func (p Ponto) LadoDaBola() turnos.LadoDoTurno {
	return p.ladoDaBola.LadoCorrente()
}

func (p Ponto) LadoDoPonto() pontos.LadoDoPonto {
	return p.ladoDoPonto
}

func (p Ponto) Terminado() bool {
	return p.terminado
}

func (p Ponto) Clonar() Ponto {
	result := New(p.ladoDaBola.Clonar(p.ladoDaBola.LadoInicial()))
	*result.golpes = make([]golpes.Golpe, len(*p.golpes))
	copy(*result.golpes, *p.golpes)

	*result.eventosAoPontuarNoPlacar = make([]pontos.AoPontuar, len(*p.eventosAoPontuarNoPlacar))
	copy(*result.eventosAoPontuarNoPlacar, *p.eventosAoPontuarNoPlacar)
	result.terminado = p.terminado

	return result
}

func (p Ponto) executeEventosAoPontuar() {
	if len(*p.eventosAoPontuarNoPlacar) > 0 {
		golpe := (*p.golpes)[p.Tamanho()-1]
		tipo, terminado := golpe.Tipo(), p.terminado

		for i := range *p.eventosAoPontuarNoPlacar {
			event := (*p.eventosAoPontuarNoPlacar)[i]
			event(tipo, terminado)
		}
	}
}
