package turno

import (
	"errors"

	"github.com/cirobispo/sandbox/internal/comum/turnos"
)

type mapData struct {
	valor       any
	funcaoReset func() any
}

func NewMapData[V any](value V, callBack func() any) mapData {
	return mapData{valor: value, funcaoReset: callBack}
}

type ParamConstructorOption func(t *Turno)
type ParamDecoratorOption func(t *Turno)

type Turno struct {
	dados map[string]mapData

	eventosAntesDeMudarTurno  []turnos.AoMudarTurno
	eventosDepoisDeMudarTurno []turnos.AoMudarTurno
}

func DefinindoLado(ladoInicial turnos.Lado) func(t *Turno) {
	return func(t *Turno) {
		startSide := NewMapData(ladoInicial, func() any { return ladoInicial })
		currentSide := NewMapData(ladoInicial, func() any { return ladoInicial })
		AdicionarDados(t, "Turno_LadoInicial", startSide)
		AdicionarDados(t, "Turno_LadoCorrente", currentSide)
	}
}

func New(param ParamConstructorOption) *Turno {
	result := &Turno{
		dados:                     make(map[string]mapData),
		eventosAntesDeMudarTurno:  make([]turnos.AoMudarTurno, 0),
		eventosDepoisDeMudarTurno: make([]turnos.AoMudarTurno, 0),
	}

	if param != nil {
		param(result)
	}
	return result
}

func (t *Turno) Decorator(decorators ...ParamDecoratorOption) *Turno {
	for i := range decorators {
		decorators[i](t)
	}

	return t
}

func (t *Turno) Execute() {
	t.executeAoMudarTurno(t.eventosAntesDeMudarTurno)

	ladoCorrente, _ := ObterDados[turnos.Lado](t, "Turno_LadoCorrente")

	if ladoCorrente > turnos.LadoA {
		ladoCorrente = -1
	}

	ladoCorrente++
	AtualizarDados(t, "Turno_LadoCorrente", ladoCorrente)

	t.executeAoMudarTurno(t.eventosDepoisDeMudarTurno)
}

func (t Turno) LadoInicial() turnos.Lado {
	result, _ := ObterDados[turnos.Lado](&t, "Turno_LadoInicial")
	return result
}

func (t Turno) LadoCorrente() turnos.Lado {
	result, _ := ObterDados[turnos.Lado](&t, "Turno_LadoCorrente")
	return result
}

func (t Turno) Clonar(ladoInicial turnos.Lado) *Turno {
	result := New(DefinindoLado(ladoInicial))

	result.eventosAntesDeMudarTurno = make([]turnos.AoMudarTurno, len(t.eventosAntesDeMudarTurno))
	copy(result.eventosAntesDeMudarTurno, t.eventosAntesDeMudarTurno)

	result.eventosDepoisDeMudarTurno = make([]turnos.AoMudarTurno, len(t.eventosDepoisDeMudarTurno))
	copy(result.eventosDepoisDeMudarTurno, t.eventosDepoisDeMudarTurno)

	total, totalAdicionado := len(t.dados)-2, 0

	for k, v := range t.dados {
		if f, _ := AdicionarDados(result, k, v); f {
			AtualizarDados(result, k, v.funcaoReset())
			totalAdicionado++
		}
	}

	if totalAdicionado != total {
		panic(errors.New("dado adicionado para clonagem não confere."))
	}

	return result
}

func (t Turno) executeAoMudarTurno(list []turnos.AoMudarTurno) {
	if len(list) > 0 {
		currentSide, _ := ObterDados[turnos.Lado](&t, "Turno_LadoCorrente")
		for i := range list {
			event := list[i]
			event(currentSide)
		}
	}
}

func (t *Turno) AdicionarAntesDeMudarTurno(callback turnos.AoMudarTurno) {
	t.eventosAntesDeMudarTurno = append(t.eventosAntesDeMudarTurno, callback)
}

func (t *Turno) AdicionarDepoisDeMudarTurno(callback turnos.AoMudarTurno) {
	t.eventosDepoisDeMudarTurno = append(t.eventosDepoisDeMudarTurno, callback)
}

func AdicionarDados(t *Turno, id string, valor mapData) (bool, error) {
	_, achou := t.dados[id]
	if !achou {
		t.dados[id] = valor
		return true, nil
	}

	return false, errors.New("id não encontrado.")
}

func AtualizarDados[V any](t *Turno, id string, valor V) (bool, error) {
	dado, existe := t.dados[id]
	if existe {
		t.dados[id] = NewMapData(valor, dado.funcaoReset)
		return true, nil
	}

	return false, errors.New("id não encontrado.")
}

func ObterDados[V any](t *Turno, id string) (V, bool) {
	dado, existe := t.dados[id]

	var result V
	if existe {
		result = dado.valor.(V)
	}

	return result, existe
}
