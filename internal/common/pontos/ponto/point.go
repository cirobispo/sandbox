package ponto

import (
	"errors"

	"github.com/cirobispo/sandbox/internal/common/pontos"
	"github.com/cirobispo/sandbox/internal/common/pontos/golpes"
	"github.com/cirobispo/sandbox/internal/common/pontos/golpes/golpe"
	"github.com/cirobispo/sandbox/internal/common/turnos/turno"
)

type Ponto struct {
	ladoDaBola               *turno.Turno
	golpes                   *[]golpes.Hitting
	terminado                bool
	eventosAoPontuarNoPlacar *[]pontos.AoPontuarNoPlacar
}

func New(sideControl *turno.Turno) Ponto {
	hit := make([]golpes.Hitting, 0, 3)
	events := make([]pontos.AoPontuarNoPlacar, 0)
	return Ponto{
		terminado:                false,
		ladoDaBola:               sideControl,
		golpes:                   &hit,
		eventosAoPontuarNoPlacar: &events,
	}
}

func temDuplaFalta(hits *[]golpes.Hitting) bool {
	lastHit := (*hits)[len(*hits)-1]
	fault := lastHit.Tipo() == golpes.HTFootFault || lastHit.Tipo() == golpes.HTServeNet || lastHit.Tipo() == golpes.HTServeOut
	if lastHit.Lado() != golpes.HTDConditional && !fault {
		return false
	}

	count := 0
	result := false
	for i := range *hits {
		hit := (*hits)[i]
		if tp := hit.Tipo(); tp == golpes.HTServeOut || tp == golpes.HTServeNet || tp == golpes.HTFootFault {
			count++
			if count > 1 {
				result = true
				break
			}
		}
	}
	return result
}

func (p *Ponto) AdicionarEventoAoPontuarNoPlacar(callback pontos.AoPontuarNoPlacar) {
	*p.eventosAoPontuarNoPlacar = append(*p.eventosAoPontuarNoPlacar, callback)
}

func (p *Ponto) AdicionaGolpe(h golpes.Hitting) {
	if p.terminado {
		return
	}

	*p.golpes = append(*p.golpes, h)

	ballInPlay := (h.Lado() == golpes.HTDNone || h.Lado() == golpes.HTDChangeSide)
	if ballInPlay {
		p.ladoDaBola.Execute()
		return
	}

	if h.Lado() == golpes.HTDConditional && !temDuplaFalta(p.golpes) {
		return
	}

	p.terminado = true
	p.executeEventosAoPontuarNoPlacar(h.Tipo(), h.Lado(), p.terminado)
}

func (p Ponto) Golpes() []golpe.Hit {
	result := make([]golpe.Hit, 0, len(*p.golpes))
	for j := range *p.golpes {
		item := (*p.golpes)[j]
		result = append(result, golpe.New(item.Tipo(), item.Lado()))
	}

	return result
}

func (p Ponto) Tamanho() int {
	result := len(*p.golpes)
	return result
}

func (p Ponto) UltimoGolpe() (golpes.TipoDoGolpe, error) {
	hitCount := p.Tamanho()
	if hitCount == 0 {
		return golpes.HTDoubleFault, errors.New("no hit found.")
	}

	return (*p.golpes)[hitCount-1].Tipo(), nil
}

func (p Ponto) LadoDaBola() turno.Turno {
	return *p.ladoDaBola
}

func (p Ponto) LadoDoPonto() pontos.LadoDoPonto {
	if !p.Terminado() {
		return pontos.LPNulo
	}

	lastHit := (*p.golpes)[len(*p.golpes)-1]
	isDoubleFault := (lastHit.Lado() == golpes.HTDConditional && temDuplaFalta(p.golpes))
	isOrdinaryPoint := (lastHit.Lado() == golpes.HTDSameSide || lastHit.Lado() == golpes.HTDOppositeSide)
	if !isOrdinaryPoint && !isDoubleFault {
		return pontos.LPNulo
	}

	if isDoubleFault {
		lastHit = golpe.NewDoubleFault()
	}

	if p.ladoDaBola.LadoCorrente() == p.ladoDaBola.LadoInicial() {
		return LadoDoGolpeParaLadoDoPonto(lastHit.Lado())
	} else {
		return LadoDoGolpeParaLadoDoPonto(lastHit.Lado()).Inverso()
	}
}

func (p Ponto) Terminado() bool {
	return p.terminado
}

func (p Ponto) Clonar() Ponto {
	result := New(p.ladoDaBola.Clonar(p.ladoDaBola.LadoInicial()))
	copy(*result.golpes, *p.golpes)
	copy(*result.eventosAoPontuarNoPlacar, *p.eventosAoPontuarNoPlacar)
	result.terminado = p.terminado

	return result
}

func (p Ponto) executeEventosAoPontuarNoPlacar(hitType golpes.TipoDoGolpe, side golpes.LadoDoGolpe, done bool) {
	for i := range *p.eventosAoPontuarNoPlacar {
		event := (*p.eventosAoPontuarNoPlacar)[i]
		event(hitType, side, done)
	}
}

func LadoDoGolpeParaLadoDoPonto(s golpes.LadoDoGolpe) pontos.LadoDoPonto {
	switch s {
	case golpes.HTDSameSide:
		return pontos.LPServico
	case golpes.HTDOppositeSide:
		return pontos.LPOposto
	default:
		return pontos.LPNulo
	}
}
