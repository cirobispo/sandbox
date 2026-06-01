package set

import (
	"testing"

	"github.com/cirobispo/sandbox/internal/common/placares"
	"github.com/cirobispo/sandbox/internal/common/pointing/point"
	"github.com/cirobispo/sandbox/internal/common/turnos"
	"github.com/cirobispo/sandbox/internal/common/turnos/turno"
)

type set struct {
	ladoDoServico turnos.LadoDoTurno
	placar        placares.EstadoResultadoEParametroPlacar
	pontos        []point.Ponto
}

func (s set) ServingSide() turnos.LadoDoTurno {
	return s.ladoDoServico
}

func (s set) Score() placares.EstadoResultadoEParametroPlacar {
	return s.placar
}

func (s set) Points() []point.Ponto {
	return s.pontos
}

type placar struct {
	ladoDoServico    turnos.LadoDoTurno
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

func (p placar) Lado() placares.LadoDoPlacar {
	if p.placarB > p.placarA {
		return placares.LPOposto
	}

	return placares.LPServico
}

func (p placar) Tipo() placares.TipoDoPlacar {
	return placares.TPJogo
}

func (s *set) ajustaLadoServico(lado turnos.LadoDoTurno) {
	a, b := s.placar.Resultado()
	s.ladoDoServico = lado
	s.placar = novoPlacar(lado, a, b)
}

func novoSet(placarA, placarB int) set {
	result := set{placar: novoPlacar(turnos.LTA, placarA, placarB), pontos: make([]point.Ponto, 0)}
	return result
}

func novoPlacar(servingSide turnos.LadoDoTurno, scoreA, scoreB int) placar {
	return placar{
		ladoDoServico: servingSide,
		placarA:       scoreA,
		placarB:       scoreB,
	}
}

func runTest(blocks []set, SideA, SideB int, t *testing.T) {
	myTurn := turno.New(turno.MudandoLado(turnos.LTA))
	mySet := New(SetPadrao(myTurn))

	sideToServe := myTurn.LadoCorrente()
	mySet.AdicionarAoAdicionarJogo(func(scoreA, scoreB int, done bool) {
		if done {
			t.Log("FINAL ")
		}

		t.Logf("%s -> Placar (%d x %d)\n", sideToServe, scoreA, scoreB)
	})

	mySet.AdicionarAoMudarLadoJogador(func() {
		t.Logf("Jogadores mudam de lado\n")
	})

	for j := range blocks {
		currentGame := mySet.NovoJogo()
		item := blocks[j]
		item.ajustaLadoServico(currentGame.ServingSide())
		sideToServe = item.ladoDoServico
		if err := mySet.AdicionarJogo(item); err != nil {
			t.Errorf("Erro ao adicionar %v. Mensagem: %s", item, err)
		}
	}
	sA, sB := mySet.Placar().Resultado()
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
