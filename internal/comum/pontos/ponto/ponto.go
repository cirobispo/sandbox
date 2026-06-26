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
	golpes                  []golpes.TipoAcaoGolpe
	ladoDoPonto             pontos.LadoDoPonto
	terminado               bool
	eventosAoAdicionarGolpe *[]pontos.AoAdicionarGolpe
}

type Pontuacao interface {
	Golpes() []golpes.TipoAcaoGolpe
	Tamanho() int
	Terminado() bool
}

type PontuacaoLados interface {
	LadoDaBola() turnos.LadoDoTurno
	LadoDoPonto() pontos.LadoDoPonto
}

func New(sideControl *turno.Turno) Ponto {
	hit := make([]golpes.TipoAcaoGolpe, 0, 3)
	events := make([]pontos.AoAdicionarGolpe, 0)
	return Ponto{
		terminado:               false,
		ladoDoPonto:             pontos.LPNulo,
		ladoDaBola:              sideControl,
		golpes:                  hit,
		eventosAoAdicionarGolpe: &events,
	}
}

func (p *Ponto) AdicionarEventoAoAdicionarGolpe(ponteiroFnc pontos.AoAdicionarGolpe) {
	*p.eventosAoAdicionarGolpe = append(*p.eventosAoAdicionarGolpe, ponteiroFnc)
}

func (p *Ponto) AdicionarGolpe(g golpes.TipoAcaoGolpe) error {
	if p.terminado {
		return errors.New("Não é possivel adicionar outro golpe com o ponto encerrado.")
	}

	p.golpes = append(p.golpes, g)
	acao := g.Acao()
	if acao == golpes.TACondicional && golpes.ExisteDuplaFalta(p.golpes) {
		acao = golpes.TAEncerrarPLO
	}
	p.terminado = (acao == golpes.TAEncerrarPLC) || (acao == golpes.TAEncerrarPLO)

	if p.terminado {
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

func (p Ponto) Golpes() []golpes.TipoAcaoGolpe {
	result := make([]golpes.TipoAcaoGolpe, len(p.golpes))
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
	return p.terminado
}

func (p Ponto) Clonar() Ponto {
	result := New(p.ladoDaBola.Clonar(p.ladoDaBola.LadoInicial()))
	result.golpes = make([]golpes.TipoAcaoGolpe, len(p.golpes))
	copy(result.golpes, p.golpes)

	*result.eventosAoAdicionarGolpe = make([]pontos.AoAdicionarGolpe, len(*p.eventosAoAdicionarGolpe))
	copy(*result.eventosAoAdicionarGolpe, *p.eventosAoAdicionarGolpe)
	result.ladoDoPonto = p.ladoDoPonto
	result.terminado = p.terminado

	return result
}

func (p Ponto) executeEventosAoAdicionarGolpe() {
	if len(*p.eventosAoAdicionarGolpe) > 0 {
		golpe := p.golpes[p.Tamanho()-1]
		tipo, terminado := golpe.Tipo(), p.terminado

		for i := range *p.eventosAoAdicionarGolpe {
			event := (*p.eventosAoAdicionarGolpe)[i]
			event(tipo, terminado)
		}
	}
}
