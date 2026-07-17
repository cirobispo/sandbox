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

func executarTest(param ParamOption, results []placar, PlacarA, PlacarB int, t *testing.T) (bool, *TieBreak) {
	tieBreak := New(param)
	for i, _ := range results {
		err := tieBreak.AdicionarPlacar(results[i])
		if err != nil {
			t.Log("Tentando adicionar placar em um jogo já encerrado.")
		}
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
		t.Errorf("Saldo do placar não é o esperado. Resultado (%v vs %v). Esperado (%v vs %v) ", pA, pB, placarA, placarB)
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

func populaSlice(t *testing.T, tam int, ladoDoTurno turnos.LadoDoTurno, pontoDecisivo bool) []placar {
	placares := make([]placar, 0, 20)

	pA, pB := 1, 0
	if ladoDoTurno == turnos.LTB {
		pA, pB = 0, 1
	}

	qtdParaMudarLado := 1
	sacador := ladoDoTurno
	t.Logf("Sacará o 1 ponto no lado: %s", sacador)
	for range tam * 2 {
		placares = append(placares, placar{placarA: pA, placarB: pB})
		qtdParaMudarLado--

		if qtdParaMudarLado == 0 {
			sacador = sacador.Inverso()
			t.Logf("Sacará o %v ponto no lado: %s", len(placares)+1, sacador)
			qtdParaMudarLado = 2
		}
	}
	tie := New(ChegarEm10(ladoDoTurno, pontoDecisivo))
	if tam == 7 {
		tie = New(ChegarEm7(ladoDoTurno, pontoDecisivo))
	}

	for j := range 2 {
		ladoAPontuar := tie.QuemEstaSacando(uint(tam*2) + uint(j))
		if ladoAPontuar != ladoDoTurno {
			pA, pB = pB, pA
		}

		placares = append(placares, placar{placarA: pA, placarB: pB})
		pA, pB = pB, pA

		if pontoDecisivo {
			break
		}
	}

	return placares
}

func chegarEm7(t *testing.T, ladoDoTurno turnos.LadoDoTurno, pontoDecisivo bool) {
	placares := populaSlice(t, 7, ladoDoTurno, pontoDecisivo)

	pontoExtra := 2
	if pontoDecisivo {
		pontoExtra = 1
	}

	pA, pB := 5+(pontoExtra*2), 5+pontoExtra

	checaPlacar(ChegarEm7(ladoDoTurno, pontoDecisivo), placares, pA, pB, t)
}

// func TestChegarEm7ASV(t *testing.T) {
// 	chegarEm7(t, turnos.LTA, true)
// }

// func TestChegarEm7ACV(t *testing.T) {
// 	chegarEm7(t, turnos.LTA, false)
// }

// func TestChegarEm7BSV(t *testing.T) {
// 	chegarEm7(t, turnos.LTB, true)
// }

// func TestChegarEm7BCV(t *testing.T) {
// 	chegarEm7(t, turnos.LTB, false)
// }

func chegarEm10(t *testing.T, ladoDoTurno turnos.LadoDoTurno, pontoDecisivo bool) {
	placares := populaSlice(t, 10, ladoDoTurno, pontoDecisivo)

	pontoExtra := 2
	if pontoDecisivo {
		pontoExtra = 1
	}

	pA, pB := 8+(pontoExtra*2), 8+pontoExtra

	checaPlacar(ChegarEm10(ladoDoTurno, pontoDecisivo), placares, pA, pB, t)
}

func TestChegarEm10ASV(t *testing.T) {
	chegarEm10(t, turnos.LTA, true)
}

func TestChegarEm10ACV(t *testing.T) {
	chegarEm10(t, turnos.LTA, false)
}

func TestChegarEm10BSV(t *testing.T) {
	chegarEm10(t, turnos.LTB, true)
}

func TestChegarEm10BCV(t *testing.T) {
	chegarEm10(t, turnos.LTB, false)
}

// func TestAleatorio(t *testing.T) {
// 	quantosPontos := 0
// 	param, ladoDoTurno := gerarParamAleatorio(&quantosPontos)
// 	pontoDecisivo := true
// 	if valorAleatorio(2) == 1 {
// 		pontoDecisivo = false
// 	}
// 	placares := populaSlice(t, quantosPontos, ladoDoTurno, pontoDecisivo)
// 	t.Logf("Quantos pontos: %v, Lado inicial: %v, Ponto decisivo: %v", quantosPontos, ladoDoTurno, pontoDecisivo)

// 	pontoExtra := 1
// 	if pontoDecisivo {
// 		pontoExtra = 0
// 	}

// 	tam := quantosPontos + pontoExtra

// 	pA, pB := tam+pontoExtra, tam-1
// 	if ladoDoTurno == turnos.LTB {
// 		pA, pB = tam-1, tam+pontoExtra
// 	}

// 	checaPlacar(param, placares, pA, pB, t)
// }

// func TestQuemSacara(t *testing.T) {
// 	tie := New(ChegarEm7(turnos.LTA, true))

// 	for i := range 15 {
// 		tie.AdicionarPlacar(placar{placarA: i, placarB: 0})
// 		a, b := tie.Resultado()
// 		pontos := uint(a + b)
// 		t.Logf("Quem saca o proximo ponto %v, quem sacou o ultimo %v", tie.QuemEstaSacando(0), tie.QuemEstaSacando(pontos))
// 	}
// }
