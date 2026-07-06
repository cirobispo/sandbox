package placartiebreak

import (
	"crypto/rand"
	"math/big"
	"testing"

	"github.com/cirobispo/sandbox/internal/comum/placares"
	"github.com/cirobispo/sandbox/internal/comum/turnos"
)

type placar struct {
	placarA, placarB int
}

func (p placar) Terminado() bool {
	return true
}

func (p placar) Resultado() (int, int) {
	return p.placarA, p.placarB
}

func (p placar) Lado() placares.LadoDoPlacar {
	if p.placarB > p.placarA {
		return placares.LPOposto
	}

	return placares.LPServico
}

func (p placar) Tipo() placares.TipoDoPlacar {
	return placares.TPPonto
}

func executarTest(param ParamOption, results []placar, PlacarA, PlacarB int, t *testing.T) (bool, TieBreak) {
	tieBreak := New(param)
	for i, _ := range results {
		tieBreak.AdicionaPlacar(results[i])
	}

	if !tieBreak.Terminado() {
		t.Log("Tiebreak não foi encerrado!")
	}

	pA, pB := tieBreak.Resultado()
	return (pA == PlacarA && pB == PlacarB), tieBreak
}

func checaPlacar(param ParamOption, placares []placar, placarA int, placarB int, t *testing.T) {
	ok, tb := executarTest(param, placares, placarA, placarB, t)
	if !ok {
		pA, pB := tb.Resultado()
		t.Errorf("Saldo do placar não é o esperado. Resultado (%v, %v) ", pA, pB)
	}
}

func valorAleatorio(maximo int) int {
	var max big.Int = *big.NewInt(int64(maximo))
	valor, err := rand.Int(rand.Reader, &max)
	if err != nil {
		return 0
	}
	return int(valor.Int64()) + 1
}

func gerarParamAleatorio(seteOuDez *int) (ParamOption, turnos.LadoDoTurno) {
	pontoDecisivo := valorAleatorio(2) == 1
	ladoDoTurno := turnos.LTA
	if valorAleatorio(2) == 1 {
		ladoDoTurno = turnos.LTB
	}

	*seteOuDez = 7
	param := ChegarEm7(ladoDoTurno, pontoDecisivo)
	if valorAleatorio(2) == 1 {
		param = ChegarEm10(ladoDoTurno, pontoDecisivo)
		*seteOuDez = 10
	}

	return param, ladoDoTurno
}

func populaSlice(tam int, ladoDoTurno turnos.LadoDoTurno) []placar {
	placares := make([]placar, 0, 10)

	placarA, placarB := tam, 0
	if ladoDoTurno == turnos.LTB {
		placarA, placarB = 0, tam
	}

	for range tam {
		placares = append(placares, placar{placarA: placarA, placarB: placarB})
	}

	return placares
}

func TestChegarEm7(t *testing.T) {
	ladoDoTurno := turnos.LTA
	placares := populaSlice(7, ladoDoTurno)

	pA, pB := 7, 0
	if ladoDoTurno == turnos.LTB {
		pA, pB = 0, 7
	}

	checaPlacar(ChegarEm7(turnos.LTA, false), placares, pA, pB, t)
}

func TestChegarEm10(t *testing.T) {
	ladoDoTurno := turnos.LTA
	placares := populaSlice(10, ladoDoTurno)

	pA, pB := 10, 0
	if ladoDoTurno == turnos.LTB {
		pA, pB = 0, 10
	}

	checaPlacar(ChegarEm10(turnos.LTA, false), placares, pA, pB, t)
}

func TestChegarEm7ou10(t *testing.T) {
	quantosPontos := 0
	param, ladoDoTurno := gerarParamAleatorio(&quantosPontos)
	placares := populaSlice(quantosPontos, ladoDoTurno)
	t.Logf("Quantos pontos: %v, Lado inicial: %v", quantosPontos, ladoDoTurno)
	pA, pB := quantosPontos, 0
	if ladoDoTurno == turnos.LTB {
		pB, pA = 0, quantosPontos
	}
	checaPlacar(param, placares, pA, pB, t)
}
