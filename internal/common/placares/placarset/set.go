package placarset

import (
	"errors"

	"github.com/cirobispo/sandbox/internal/common/placares"
	"github.com/cirobispo/sandbox/internal/common/turning"
)

type AoMudarPlacar func(placarA, placarB int, tieBreak, terminado bool)
type ParamOption func(score *Set)

type Set struct {
	ladoInicio           turning.TurningSide
	maiorEmpate          int
	jogosConfirmaVitoria int
	placarA, placarB     int

	eventosAoMudarPlacar []AoMudarPlacar
}

func SetPadrao(ladoInicio turning.TurningSide) ParamOption {
	return func(score *Set) {
		score.ladoInicio = ladoInicio
		score.maiorEmpate = 6
		score.jogosConfirmaVitoria = 2
	}
}

func TamanhoETieBreak(ladoInicio turning.TurningSide, tamanho int, jogoDecisivo, tieBreakNoMaiorEmpate bool) ParamOption {
	return func(placar *Set) {
		placar.ladoInicio = ladoInicio
		placar.maiorEmpate = tamanho
		placar.jogosConfirmaVitoria = 2
		if jogoDecisivo {
			placar.jogosConfirmaVitoria--
		}
	}
}

func New(param ParamOption) Set {
	result := Set{
		placarA:              0,
		placarB:              0,
		eventosAoMudarPlacar: make([]AoMudarPlacar, 0),
	}

	if param != nil {
		param(&result)
	}

	return result
}

func (s Set) executarAoMudarPlacar(placarA, placarB int) {
	done, isTieBreak := s.Terminado(), s.IsTieBreak()
	for i := range s.eventosAoMudarPlacar {
		event := s.eventosAoMudarPlacar[i]
		event(placarA, placarB, isTieBreak, done)
	}
}

func (s *Set) placares() (*int, *int) {
	sA, sB := &s.placarA, &s.placarB
	if s.ladoInicio == turning.TSB {
		sA, sB = &s.placarB, &s.placarA
	}

	return sA, sB
}

func (s *Set) AdicionarAoMudarPlacar(event AoMudarPlacar) {
	s.eventosAoMudarPlacar = append(s.eventosAoMudarPlacar, event)
}

func (s *Set) AdicionarPlacar(placar placares.EstadoEParametroPlacar) error {
	if s.Terminado() { // am I acepting more points?
		return errors.New("Score completed already.")
	}

	if placar.Tipo() != placares.TPJogo {
		return errors.New("This is not a Game Score.")
	}

	if !placar.Terminado() { // am I acepting more points?
		return errors.New("Game is not completed.")
	}

	sA, sB := s.placares()
	sideToAdd := sA

	if placar.Lado() == placares.LPOposto {
		sideToAdd = sB
	}

	*sideToAdd += 1
	s.executarAoMudarPlacar(s.placarA, s.placarB)
	return nil
}

func (s Set) Resultado() (int, int) {
	return s.placarA, s.placarB
}

func (s Set) Lado() placares.LadoDoPlacar {
	return placares.Lado(s)
}

func (s Set) Tipo() placares.TipoDoPlacar {
	return placares.TPSet
}

func (s Set) Terminado() bool {
	sA, sB := s.Resultado()

	diff := s.jogosConfirmaVitoria
	if sA > s.maiorEmpate || sB > s.maiorEmpate {
		diff = 1
	}

	sideAWins := (sA >= s.maiorEmpate && sA-sB >= diff)
	sideBWins := (sB >= s.maiorEmpate && sB-sA >= diff)
	result := sideAWins || sideBWins

	return result
}

func (s Set) IsTieBreak() bool {
	sA, sB := s.Resultado()
	tie := s.maiorEmpate
	if s.jogosConfirmaVitoria == 1 {
		tie = s.maiorEmpate - 1
	}
	result := (sA >= tie && sB >= tie)

	return result
}
