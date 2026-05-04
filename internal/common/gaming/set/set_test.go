package set

import (
	"testing"

	"github.com/cirobispo/sandbox/internal/common/pointing/point"
	"github.com/cirobispo/sandbox/internal/common/scoring"
	"github.com/cirobispo/sandbox/internal/common/turning"
	"github.com/cirobispo/sandbox/internal/common/turning/turn"
)

type set struct {
	ladoDoServico turning.TurningSide
	placar        scoring.EstadoResultadoEParametroPlacar
	pontos        []point.Point
}

func (s set) ServingSide() turning.TurningSide {
	return s.ladoDoServico
}

func (s set) Score() scoring.EstadoResultadoEParametroPlacar {
	return s.placar
}

func (s set) Points() []point.Point {
	return s.pontos
}

type placar struct {
	ladoDoServico    turning.TurningSide
	pontoDecisivo    bool
	placarA, placarB int
}

func (p placar) Terminado() bool {
	sA, sB := p.Resultado()
	diff := 2
	if p.pontoDecisivo {
		diff--
	}

	AWins := sA >= 4 && (sA-sB) >= diff
	BWins := sB >= 4 && (sB-sA) >= diff

	result := AWins || BWins

	return result
}

func (p placar) Resultado() (int, int) {
	return p.placarA, p.placarB
}

func (p placar) Lado() scoring.LadoDoPlacar {
	if p.placarB > p.placarA {
		return scoring.LPOposto
	}

	return scoring.LPServico
}

func (p placar) Tipo() scoring.TipoDoPlacar {
	return scoring.TPJogo
}

func (s *set) ajustaLadoServico(lado turning.TurningSide) {
	a, b := s.placar.Resultado()
	s.ladoDoServico = lado
	s.placar = novoPlacar(lado, a, b)
}

func novoSet(placarA, placarB int) set {
	result := set{placar: novoPlacar(turning.TSA, placarA, placarB), pontos: make([]point.Point, 0)}
	return result
}

func novoPlacar(servingSide turning.TurningSide, scoreA, scoreB int) placar {
	return placar{
		ladoDoServico: servingSide,
		placarA:       scoreA,
		placarB:       scoreB,
	}
}

func runTest(blocks []set, SideA, SideB int, t *testing.T) {
	myTurn := turn.New(turn.WithTurningSide(turning.TSA))
	mySet := New(WithDefaultSet(myTurn))

	sideToServe := myTurn.CurrentSide()
	mySet.AddOnAddingGameEvent(func(scoreA, scoreB int, done bool) {
		if done {
			t.Log("FINAL ")
		}

		t.Logf("%s -> Placar (%d x %d)\n", sideToServe, scoreA, scoreB)
	})

	mySet.AddOnPlayerChangeEvent(func() {
		t.Logf("Jogadores mudam de lado\n")
	})

	for j := range blocks {
		currentGame := mySet.NewGame()
		item := blocks[j]
		item.ajustaLadoServico(currentGame.ServingSide())
		sideToServe = item.ladoDoServico
		if err := mySet.AddGame(item); err != nil {
			t.Errorf("Erro ao adicionar %v. Mensagem: %s", item, err)
		}
	}
	sA, sB := mySet.Score().Resultado()
	if sA != SideA || sB != SideB {
		for i := range blocks {
			item := blocks[i]
			a, b := item.Score().Resultado()
			t.Logf("Placar do jogo #%d (%d x %d)\n", i+1, a, b)
		}
		t.Errorf("\n\nSet deveria ser (%d x %d) não (%d x %d)\n", SideA, SideB, sA, sB)
	}
}

func Test_6x4(t *testing.T) {
	data := []set{
		novoSet(6, 4), novoSet(2, 6),
		novoSet(6, 3), novoSet(1, 6),
		novoSet(6, 4), novoSet(4, 6),
		novoSet(6, 4), novoSet(4, 6),
		novoSet(6, 4), novoSet(6, 4),
		// newItem(6, 4), newItem(6, 4),
	}
	runTest(data, 6, 4, t)
}

func Test_4x6(t *testing.T) {
	data := []set{
		novoSet(6, 4), novoSet(2, 6),
		novoSet(6, 3), novoSet(1, 6),
		novoSet(6, 4), novoSet(4, 6),
		novoSet(6, 4), novoSet(5, 7),
		novoSet(5, 7), novoSet(4, 6),
		// newItem(4, 6), newItem(4, 6),
	}
	runTest(data, 4, 6, t)
}
