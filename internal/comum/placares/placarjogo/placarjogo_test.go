package placarjogo

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

func executarTest(ladoInicial turnos.LadoDoTurno, pontoDecisivo bool, results []placar, PlacarA, PlacarB int, t *testing.T) (bool, Jogo) {
	jogo := New(ladoInicial, pontoDecisivo)
	for i, _ := range results {
		jogo.AdicionarPlacar(results[i])
	}

	if !jogo.Terminado() {
		t.Log("Jogo não foi encerrado!")
	}

	pA, pB := jogo.Resultado()
	return (pA == PlacarA && pB == PlacarB), *jogo
}

func checaPlacar(ladoInicial turnos.LadoDoTurno, pontoDecisivo bool, placares []placar, placarA int, placarB int, t *testing.T) {
	ok, tb := executarTest(ladoInicial, pontoDecisivo, placares, placarA, placarB, t)
	pA, pB := tb.Resultado()
	if !ok {
		t.Errorf("Saldo do placar não é o esperado. Resultado (%v, %v) ", pA, pB)
	}
	t.Logf("Saldo do placar (%v, %v) ", pA, pB)
}

func valorAleatorio(maximo int) int {
	var max big.Int = *big.NewInt(int64(maximo))
	valor, err := rand.Int(rand.Reader, &max)
	if err != nil {
		return 0
	}
	return int(valor.Int64()) + 1
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

func TestChegarCOMeSEMConfirmacao(t *testing.T) {
	ladoDoTurno := turnos.LTA
	if valorAleatorio(2) == 1 {
		ladoDoTurno = turnos.LTB
	}

	pontosParaJogo := 4
	if valorAleatorio(2) == 1 {
		pontosParaJogo = 5
	}

	placares := make([]placar, 0, 8)
	placares = append(placares, populaSlice(3, ladoDoTurno.Inverso())...)
	placares = append(placares, populaSlice(pontosParaJogo, ladoDoTurno)...)

	pA, pB := pontosParaJogo, 3

	t.Logf("Lado inicial: %v, Ponto decisivo: %v", ladoDoTurno, (pontosParaJogo == 4))
	checaPlacar(ladoDoTurno, (pontosParaJogo == 4), placares, pA, pB, t)
}
