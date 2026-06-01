package turno

import (
	"errors"

	"github.com/cirobispo/sandbox/internal/common/turnos"
)

type mapData struct {
	valor       any
	funcaoReset func() any
}

func NewMapData[V any](value V, callBack func() any) mapData {
	return mapData{valor: value, funcaoReset: callBack}
}

type ParamOption func(t *Turno)

type Turno struct {
	dados map[string]mapData

	eventosAntesDeMudarTurno  []turnos.AoMudarTurno
	eventosDepoisDeMudarTurno []turnos.AoMudarTurno
}

func (t *Turno) AdicionarAntesDeMudarTurno(callback turnos.AoMudarTurno) {
	t.eventosAntesDeMudarTurno = append(t.eventosAntesDeMudarTurno, callback)
}

func (t *Turno) AdicionarDepoisDeMudarTurno(callback turnos.AoMudarTurno) {
	t.eventosDepoisDeMudarTurno = append(t.eventosDepoisDeMudarTurno, callback)
}

func (t *Turno) Execute() {
	t.executeAoMudarTurno(t.eventosAntesDeMudarTurno)

	ladoCorrente, _ := ObterDados[turnos.LadoDoTurno](t, "Turn_currentSide")

	if ladoCorrente > turnos.LTA {
		ladoCorrente = -1
	}

	ladoCorrente++
	AtualizarDados(t, "Turn_currentSide", ladoCorrente)

	t.executeAoMudarTurno(t.eventosDepoisDeMudarTurno)
}

func (t Turno) LadoInicial() turnos.LadoDoTurno {
	result, _ := ObterDados[turnos.LadoDoTurno](&t, "Turn_startSide")
	return result
}

func (t Turno) LadoCorrente() turnos.LadoDoTurno {
	result, _ := ObterDados[turnos.LadoDoTurno](&t, "Turn_currentSide")
	return result
}

func (t Turno) Clonar(start turnos.LadoDoTurno) *Turno {
	result := New(MudandoLado(start))
	copy(result.eventosAntesDeMudarTurno, t.eventosAntesDeMudarTurno)
	copy(result.eventosDepoisDeMudarTurno, t.eventosDepoisDeMudarTurno)

	dataCount, dataAdded := len(t.dados)-2, 0

	for k, v := range t.dados {
		if f, _ := AdicionarDados(result, k, v); f {
			AtualizarDados(result, k, v.funcaoReset())
			dataAdded++
		}
	}

	if dataAdded != dataCount {
		panic(errors.New("data added to clone turn does not match."))
	}

	return result
}

func (t Turno) executeAoMudarTurno(list []turnos.AoMudarTurno) {
	currentSide, _ := ObterDados[turnos.LadoDoTurno](&t, "Turn_currentSide")
	for i := range list {
		event := list[i]
		event(currentSide)
	}
}

func MudandoLado(start turnos.LadoDoTurno) func(t *Turno) {
	return func(t *Turno) {
		startSide := NewMapData(start, func() any { return start })
		currentSide := NewMapData(start, func() any { return start })
		AdicionarDados(t, "Turn_startSide", startSide)
		AdicionarDados(t, "Turn_currentSide", currentSide)
	}
}

func New(param func(t *Turno)) *Turno {
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

func AdicionarDados(t *Turno, id string, data mapData) (bool, error) {
	_, f := t.dados[id]
	if !f {
		t.dados[id] = data
		return true, nil
	}

	return false, errors.New("id not found.")
}

func AtualizarDados[V any](t *Turno, id string, data V) (bool, error) {
	d, f := t.dados[id]
	if f {
		t.dados[id] = NewMapData(data, d.funcaoReset)
		return true, nil
	}

	return false, errors.New("id not found.")
}

func ObterDados[V any](t *Turno, id string) (V, bool) {
	data, found := t.dados[id]

	var result V
	if found {
		result = data.valor.(V)
	}

	return result, found
}
