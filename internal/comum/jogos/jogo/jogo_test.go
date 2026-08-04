package jogo

import (
	"testing"

	"github.com/cirobispo/sandbox/internal/comum/golpes/golpe"
	"github.com/cirobispo/sandbox/internal/comum/pontos/ponto"
	"github.com/cirobispo/sandbox/internal/comum/turnos"
	"github.com/cirobispo/sandbox/internal/comum/turnos/turno"
)

func ace(ladoInicial turnos.Lado) ponto.Ponto {
	result := ponto.New(turno.New(turno.DefinindoLado(ladoInicial)))
	result.AdicionarGolpe(golpe.NewAce())

	return result
}

func winnerDoSacador(ladoInicial turnos.Lado) ponto.Ponto {
	result := ponto.New(turno.New(turno.DefinindoLado(ladoInicial)))
	result.AdicionarGolpe(golpe.NewServicoDentro())
	result.AdicionarGolpe(golpe.NewRetornoDentro())
	result.AdicionarGolpe(golpe.NewWinner())

	return result
}

func winnerDoRecebedor(ladoInicial turnos.Lado) ponto.Ponto {
	result := ponto.New(turno.New(turno.DefinindoLado(ladoInicial)))
	result.AdicionarGolpe(golpe.NewServicoDentro())
	result.AdicionarGolpe(golpe.NewWinner())

	return result
}

func naoTocouDoSacador(ladoInicial turnos.Lado) ponto.Ponto {
	result := ponto.New(turno.New(turno.DefinindoLado(ladoInicial)))
	result.AdicionarGolpe(golpe.NewServicoDentro())
	result.AdicionarGolpe(golpe.NewRetornoDentro())
	result.AdicionarGolpe(golpe.NewNaoTocou())

	return result
}

func naoTocouDoRecebedor(ladoInicial turnos.Lado) ponto.Ponto {
	result := ponto.New(turno.New(turno.DefinindoLado(ladoInicial)))
	result.AdicionarGolpe(golpe.NewServicoDentro())
	result.AdicionarGolpe(golpe.NewRetornoDentro())
	result.AdicionarGolpe(golpe.NewDevolveuDentro())
	result.AdicionarGolpe(golpe.NewNaoTocou())

	return result
}

func foraDoSacador(ladoInicial turnos.Lado) ponto.Ponto {
	result := ponto.New(turno.New(turno.DefinindoLado(ladoInicial)))
	result.AdicionarGolpe(golpe.NewServicoDentro())
	result.AdicionarGolpe(golpe.NewRetornoDentro())
	result.AdicionarGolpe(golpe.NewDevolveuFora())

	return result
}

func foraDoRecebedor(ladoInicial turnos.Lado) ponto.Ponto {
	result := ponto.New(turno.New(turno.DefinindoLado(ladoInicial)))
	result.AdicionarGolpe(golpe.NewServicoDentro())
	result.AdicionarGolpe(golpe.NewRetornoDentro())
	result.AdicionarGolpe(golpe.NewDevolveuDentro())
	result.AdicionarGolpe(golpe.NewDevolveuFora())

	return result
}

func runTest(qualSacador turnos.Lado, j *Jogo, pontos []ponto.Ponto, A, B int, t *testing.T) {
	j.AdicionarEventoAoAdicionarPonto(func(placarA, placarB int, terminado bool) {
		//var descricao = []string{"love", "15", "30", "40", "ad", "game"}
		//tA, tB := placares.TraduzirPlacar(descricao, placarA, placarB)
		tA, tB := placarA, placarB
		if terminado {
			t.Logf("Game FINAL status: ( %v x %v )\n", tA, tB)
			t.Log()
			return
		}

		t.Logf("Game status: ( %v x %v )\n", tA, tB)
		t.Log()
	})

	for i, _ := range pontos {
		ponto := pontos[i]
		j.AdicionarPonto(&ponto)
	}

	if rA, rB := j.Placar().Resultado(); !j.Placar().Terminado() || (rA != A || rB != B) {
		t.Errorf("Resultado não esperado. Sacador: %v. Jogo encerrado: %v, resultado (%v x %v), aguardado (%v x %v)", qualSacador, j.Placar().Terminado(), rA, rB, A, B)
	}
}

// func TestTurnA_Game40(t *testing.T) {
// 	ladoInicial := turnos.LTA
// 	pontos := []ponto.Ponto{
// 		ace(ladoInicial),                 //1x0
// 		winnerDoRecebedor(ladoInicial),   //1x1
// 		winnerDoSacador(ladoInicial),     //2x1
// 		naoTocouDoRecebedor(ladoInicial), //3x1
// 		naoTocouDoSacador(ladoInicial),   //3x2
// 		foraDoSacador(ladoInicial),       //3x3
// 		naoTocouDoRecebedor(ladoInicial), //4x3
// 		foraDoRecebedor(ladoInicial),     //5x3
// 	}

// 	j := New(Regular(ladoInicial, false))

// 	runTest(ladoInicial, j, pontos, 5, 3, t)
// }

func TestTieBreak(t *testing.T) {
	ladoInicial := turnos.LadoA
	pontos := []ponto.Ponto{
		ace(ladoInicial),                 //1x0
		winnerDoRecebedor(ladoInicial),   //1x1
		winnerDoSacador(ladoInicial),     //2x1
		naoTocouDoRecebedor(ladoInicial), //3x1
		naoTocouDoSacador(ladoInicial),   //3x2
		foraDoSacador(ladoInicial),       //3x3
		naoTocouDoRecebedor(ladoInicial), //4x3
		foraDoRecebedor(ladoInicial),     //5x3
		ace(ladoInicial),                 //6x3
		ace(ladoInicial),                 //7x3
	}

	j := New(TieBreak(ladoInicial, false))

	runTest(ladoInicial, j, pontos, 7, 3, t)
}

func TestSuperTieBreak(t *testing.T) {
	ladoInicial := turnos.LadoA
	pontos := []ponto.Ponto{
		ace(ladoInicial),                 //1x0
		winnerDoRecebedor(ladoInicial),   //1x1
		winnerDoSacador(ladoInicial),     //2x1
		naoTocouDoRecebedor(ladoInicial), //3x1
		naoTocouDoSacador(ladoInicial),   //3x2
		foraDoSacador(ladoInicial),       //3x3
		naoTocouDoRecebedor(ladoInicial), //4x3
		foraDoRecebedor(ladoInicial),     //5x3
		ace(ladoInicial),                 //6x3
		ace(ladoInicial),                 //7x3
		winnerDoSacador(ladoInicial),     //8x3
		ace(ladoInicial),                 //9x3
		ace(ladoInicial),                 //10x3
	}

	j := New(SuperTieBreak(ladoInicial, false))

	runTest(ladoInicial, j, pontos, 10, 3, t)
}
