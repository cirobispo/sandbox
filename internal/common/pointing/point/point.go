package point

import (
	"errors"

	"github.com/cirobispo/sandbox/internal/common/pointing"
	"github.com/cirobispo/sandbox/internal/common/pointing/hitting"
	"github.com/cirobispo/sandbox/internal/common/pointing/hitting/hit"
	"github.com/cirobispo/sandbox/internal/common/turnos/turno"
)

type Ponto struct {
	ladoDaBola               *turno.Turno
	golpes                   *[]hitting.Hitting
	terminado                bool
	eventosAoPontuarNoPlacar *[]pointing.AoPontuarNoPlacar
}

func New(sideControl *turno.Turno) Ponto {
	hit := make([]hitting.Hitting, 0, 3)
	events := make([]pointing.AoPontuarNoPlacar, 0)
	return Ponto{
		terminado:                false,
		ladoDaBola:               sideControl,
		golpes:                   &hit,
		eventosAoPontuarNoPlacar: &events,
	}
}

func temDuplaFalta(hits *[]hitting.Hitting) bool {
	lastHit := (*hits)[len(*hits)-1]
	fault := lastHit.Type() == hitting.HTFootFault || lastHit.Type() == hitting.HTServeNet || lastHit.Type() == hitting.HTServeOut
	if lastHit.Side() != hitting.HTDConditional && !fault {
		return false
	}

	count := 0
	result := false
	for i := range *hits {
		hit := (*hits)[i]
		if tp := hit.Type(); tp == hitting.HTServeOut || tp == hitting.HTServeNet || tp == hitting.HTFootFault {
			count++
			if count > 1 {
				result = true
				break
			}
		}
	}
	return result
}

func (p *Ponto) AdicionarEventoAoPontuarNoPlacar(callback pointing.AoPontuarNoPlacar) {
	*p.eventosAoPontuarNoPlacar = append(*p.eventosAoPontuarNoPlacar, callback)
}

func (p *Ponto) AdicionaGolpe(h hitting.Hitting) {
	if p.terminado {
		return
	}

	*p.golpes = append(*p.golpes, h)

	ballInPlay := (h.Side() == hitting.HTDNone || h.Side() == hitting.HTDChangeSide)
	if ballInPlay {
		p.ladoDaBola.Execute()
		return
	}

	if h.Side() == hitting.HTDConditional && !temDuplaFalta(p.golpes) {
		return
	}

	p.terminado = true
	p.executeEventosAoPontuarNoPlacar(h.Type(), h.Side(), p.terminado)
}

func (p Ponto) Golpes() []hit.Hit {
	result := make([]hit.Hit, 0, len(*p.golpes))
	for j := range *p.golpes {
		item := (*p.golpes)[j]
		result = append(result, hit.New(item.Type(), item.Side()))
	}

	return result
}

func (p Ponto) Tamanho() int {
	result := len(*p.golpes)
	return result
}

func (p Ponto) UltimoGolpe() (hitting.HitType, error) {
	hitCount := p.Tamanho()
	if hitCount == 0 {
		return hitting.HTDoubleFault, errors.New("no hit found.")
	}

	return (*p.golpes)[hitCount-1].Type(), nil
}

func (p Ponto) LadoDaBola() turno.Turno {
	return *p.ladoDaBola
}

func (p Ponto) LadoDoPonto() pointing.LadoDoPonto {
	if !p.Terminado() {
		return pointing.LPNulo
	}

	lastHit := (*p.golpes)[len(*p.golpes)-1]
	isDoubleFault := (lastHit.Side() == hitting.HTDConditional && temDuplaFalta(p.golpes))
	isOrdinaryPoint := (lastHit.Side() == hitting.HTDSameSide || lastHit.Side() == hitting.HTDOppositeSide)
	if !isOrdinaryPoint && !isDoubleFault {
		return pointing.LPNulo
	}

	if isDoubleFault {
		lastHit = hit.NewDoubleFault()
	}

	if p.ladoDaBola.LadoCorrente() == p.ladoDaBola.LadoInicial() {
		return LadoDoGolpeParaLadoDoPonto(lastHit.Side())
	} else {
		return LadoDoGolpeParaLadoDoPonto(lastHit.Side()).Inverso()
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

func (p Ponto) executeEventosAoPontuarNoPlacar(hitType hitting.HitType, side hitting.HitSide, done bool) {
	for i := range *p.eventosAoPontuarNoPlacar {
		event := (*p.eventosAoPontuarNoPlacar)[i]
		event(hitType, side, done)
	}
}

func LadoDoGolpeParaLadoDoPonto(s hitting.HitSide) pointing.LadoDoPonto {
	switch s {
	case hitting.HTDSameSide:
		return pointing.LPServico
	case hitting.HTDOppositeSide:
		return pointing.LPOposto
	default:
		return pointing.LPNulo
	}
}
