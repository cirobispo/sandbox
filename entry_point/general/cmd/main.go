package main

import (
	"fmt"

	"github.com/cirobispo/sandbox/internal/state_machine/ponto"
)

type VerificarResultado func(valores ...int) bool

type Placar struct {
	placarA, placarB int
	confirmar        bool
	verificar        VerificarResultado
}

func boolParaInt(value bool) int {
	if value {
		return 1
	}
	return 0
}

func (p Placar) Terminou() bool {
	if p.verificar == nil {
		return false
	}

	return p.verificar(p.placarA, p.placarB, boolParaInt(p.confirmar))
}

func readCommand(s *ponto.PointState) int {
	fmt.Printf("\t0 Sair")
	subStates := s.SubStates()
	// for j := range subStates {
	// item := subStates[j]
	// fmt.Printf("\t%d %s", j+1, item.Hit().Tipo())
	// }
	size := len(subStates)
	result := 0
	if size > 0 {
		for {
			fmt.Printf("\nDigite entre 1 e %d\n", size)
			num := 0
			if n, err := fmt.Scanln(&num); (err != nil) || (n < 1) {
				if err != nil {
					fmt.Println(err)
				}
				continue
			}

			if num >= 0 && num <= size {
				result = num
				break
			}

			fmt.Println("numero inválido!")
		}
	}

	return result
}

func main() {
	/**
		point := point.New(turn.New(turn.WithTurningSide(turning.TSA)))
		c := sm_point.NewPointStateControl(&point)

		var next int
		for {
			if next = readCommand(c.CurrentState()); next == 0 {
				break
			}

			state := c.FindState(c.CurrentState().StatesToChoose()[next-1])
			if state != nil {
				c.UpdateState(state)
				if !c.BallInPlay() {
					break
				}
			}
		}
	/**/
	p := Placar{placarA: 3, placarB: 5, confirmar: true}
	p.verificar = testaSemVantagem

	if p.Terminou() {
		fmt.Println("Acabou o sem vantagem!")
	}

	p.verificar = testaComVantagem

	if p.Terminou() {
		fmt.Println("Acabou o com vantagem!")
	}
}

func testaSemVantagem(valores ...int) bool {
	sA, sB := valores[0], valores[1]
	if sA > 3 || sB > 3 {
		return true
	}

	return false
}

func testaComVantagem(valores ...int) bool {
	if len(valores) < 3 {
		panic("Não existem parametros suficientes!")
	}

	sA, sB, dp := valores[0], valores[1], valores[2]
	if sA > 3 || sB > 3 {
		return dp == 0 || (dp == 1 && sA-sB >= 2 || sB-sA >= 2)
	}

	return false
}
