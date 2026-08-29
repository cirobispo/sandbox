package turno

import "errors"

type mapData struct {
	valor       any
	funcaoReset func() any
}

func NewMapData[V any](value V, callBack func() any) mapData {
	return mapData{valor: value, funcaoReset: callBack}
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

func Chaves(t *Turno) []string {
	result := make([]string, 0, len(t.dados))
	for k, _ := range t.dados {
		result = append(result, k)
	}

	return result
}
