package ponto

import (
	"errors"

	"github.com/cirobispo/sandbox/internal/comum/pontos"
	"github.com/cirobispo/sandbox/internal/comum/pontos/golpes"
	"github.com/cirobispo/sandbox/internal/comum/turnos"
	"github.com/cirobispo/sandbox/internal/comum/turnos/turno"
)

type Ponto struct {
	ladoDaBola              *turno.Turno
	golpes                  *[]golpes.Golpe
	ladoDoPonto             pontos.LadoDoPonto
	terminado               bool
	eventosAoAdicionarGolpe *[]pontos.AoAdicionarGolpe
}

func New(sideControl *turno.Turno) Ponto {
	hit := make([]golpes.Golpe, 0, 3)
	events := make([]pontos.AoAdicionarGolpe, 0)
	return Ponto{
		terminado:               false,
		ladoDoPonto:             pontos.LPNulo,
		ladoDaBola:              sideControl,
		golpes:                  &hit,
		eventosAoAdicionarGolpe: &events,
	}
}

func (p *Ponto) AdicionarEventoAoAdicionarGolpe(ponteiroFnc pontos.AoAdicionarGolpe) {
	*p.eventosAoAdicionarGolpe = append(*p.eventosAoAdicionarGolpe, ponteiroFnc)
}

func (p *Ponto) AdicionarGolpe(g golpes.Golpe) error {
	if p.terminado {
		return errors.New("Não é possivel adicionar outro golpe com o ponto encerrado.")
	}

	*p.golpes = append(*p.golpes, g)
	acao := g.Acao(*p.golpes)
	if acao == golpes.TANulo {
		return nil
	}

	p.terminado = (acao == golpes.TAEncerrarPLC) || (acao == golpes.TAEncerrarPLO)
	if !p.terminado {
		p.ladoDaBola.Execute()
		p.executeEventosAoAdicionarGolpe()
		return nil
	}

	p.ladoDoPonto = pontos.LPCorrente
	if acao == golpes.TAEncerrarPLO {
		p.ladoDoPonto = pontos.LPOposto
	}

	if p.ladoDaBola.LadoCorrente() != p.ladoDaBola.LadoInicial() {
		p.ladoDoPonto = p.ladoDoPonto.Inverso()
	}

	p.executeEventosAoAdicionarGolpe()

	return nil
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

	*result.eventosAoAdicionarGolpe = make([]pontos.AoAdicionarGolpe, len(*p.eventosAoAdicionarGolpe))
	copy(*result.eventosAoAdicionarGolpe, *p.eventosAoAdicionarGolpe)
	result.ladoDoPonto = p.ladoDoPonto
	result.terminado = p.terminado

	return result
}

func (p Ponto) executeEventosAoAdicionarGolpe() {
	if len(*p.eventosAoAdicionarGolpe) > 0 {
		golpe := (*p.golpes)[p.Tamanho()-1]
		tipo, terminado := golpe.Tipo(), p.terminado

		for i := range *p.eventosAoAdicionarGolpe {
			event := (*p.eventosAoAdicionarGolpe)[i]
			event(tipo, terminado)
		}
	}
}
