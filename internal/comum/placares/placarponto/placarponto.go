package placarponto

import (
	"fmt"

	"github.com/cirobispo/sandbox/internal/comum/placares"
	"github.com/cirobispo/sandbox/internal/comum/pontos"
	"github.com/cirobispo/sandbox/internal/comum/pontos/ponto"
	"github.com/cirobispo/sandbox/internal/comum/turnos"
)

type PlacarPonto struct {
	ponto            pontos.Pontuando
	ladoDoPlacar     placares.LadoDoPlacar
	placarA, placarB int
}

func New(ponto *ponto.Ponto, ladoInicial, qualSacador turnos.LadoDoTurno) PlacarPonto {
	result := PlacarPonto{ponto: ponto, placarA: 0, placarB: 0}

	fmt.Printf("Quem está sacando: %v ", qualSacador)

	if ponto.LadoDoPonto() == pontos.LPCorrente {
		result.placarA, result.placarB = 1, 0
		result.ladoDoPlacar = placares.LPServico
		if qualSacador == turnos.LTB {
			result.placarA, result.placarB = 0, 1
			result.ladoDoPlacar = placares.LPOposto
		}
	} else {
		result.placarA, result.placarB = 0, 1
		result.ladoDoPlacar = placares.LPOposto
		if qualSacador == turnos.LTA {
			result.placarA, result.placarB = 1, 0
			result.ladoDoPlacar = placares.LPServico
		}
	}

	return result
}

func (p PlacarPonto) Lado() placares.LadoDoPlacar {
	return p.ladoDoPlacar
}

func (p PlacarPonto) Tipo() placares.TipoDoPlacar {
	return placares.TPPonto
}

func (p PlacarPonto) Terminado() bool {
	return p.ponto.Terminado()
}

func (p PlacarPonto) Resultado() (int, int) {
	return p.placarA, p.placarB
}
