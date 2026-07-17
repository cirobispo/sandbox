package placartiebreak

import (
	"errors"

	"github.com/cirobispo/sandbox/internal/comum/placares"
	"github.com/cirobispo/sandbox/internal/comum/turnos"
)

type AoMudarPlacar func(placarA, placarB int, terminado bool)
type AoMudarSacador func(lado turnos.LadoDoTurno)
type TestaEncerramento func(valores ...int) bool

type ParamOption func(t *TieBreak)

type TieBreak struct {
	ladoInicial       turnos.LadoDoTurno
	pontoDecisivo     bool
	placarA, placarB  int
	testaEncerramento TestaEncerramento

	eventosAoMudarPlacar []AoMudarPlacar
}

func testaEncerrarEm7(valores ...int) bool {
	sA, sB, pontoDecisivo := valores[0], valores[1], valores[2] == 1
	diff := 2
	if pontoDecisivo {
		diff--
	}

	AWins := sA >= 7 && (sA-sB) >= diff
	BWins := sB >= 7 && (sB-sA) >= diff

	result := AWins || BWins

	return result
}

func testaEncerrarEm10(valores ...int) bool {
	sA, sB, pontoDecisivo := valores[0], valores[1], valores[2] == 1
	diff := 2
	if pontoDecisivo {
		diff--
	}

	AWins := sA >= 10 && (sA-sB) >= diff
	BWins := sB >= 10 && (sB-sA) >= diff

	result := AWins || BWins

	return result
}

func ChegarEm7(ladoInicial turnos.LadoDoTurno, pontoDecisivo bool) ParamOption {
	return func(t *TieBreak) {
		t.ladoInicial = ladoInicial
		t.pontoDecisivo = pontoDecisivo
		t.DefinirTestaEncerramento(testaEncerrarEm7)
	}
}

func ChegarEm10(ladoInicial turnos.LadoDoTurno, pontoDecisivo bool) ParamOption {
	return func(t *TieBreak) {
		t.ladoInicial = ladoInicial
		t.pontoDecisivo = pontoDecisivo
		t.DefinirTestaEncerramento(testaEncerrarEm10)
	}
}

func New(param ParamOption) *TieBreak {
	result := &TieBreak{
		placarA:              0,
		placarB:              0,
		eventosAoMudarPlacar: make([]AoMudarPlacar, 0),
	}

	param(result)
	return result
}

func (j TieBreak) QuemEstaSacando(qualPonto uint) turnos.LadoDoTurno {
	trocasDeSacador := uint(((j.placarA+j.placarB)%2 + (j.placarA+j.placarB)/2))
	if qualPonto != 0 {
		trocasDeSacador = qualPonto - 1
	}

	if trocasDeSacador%2 == 1 {
		return j.ladoInicial.Inverso()
	}
	return j.ladoInicial
}

func (j *TieBreak) placares() (*int, *int) {
	sA, sB := &j.placarA, &j.placarB
	if j.ladoInicial == turnos.LTB {
		sA, sB = &j.placarB, &j.placarA
	}

	return sA, sB
}

func (j *TieBreak) placarInvertido(side *int) *int {
	sA, sB := j.placares()
	if side == sA {
		return sB
	}
	return sA
}

func (j TieBreak) executarEventosAoMudarPlacar() {
	if len(j.eventosAoMudarPlacar) > 0 {
		placarA, placarB := j.Resultado()
		terminado := j.Terminado()
		for i := range j.eventosAoMudarPlacar {
			event := j.eventosAoMudarPlacar[i]
			event(placarA, placarB, terminado)
		}
	}
}

func (j TieBreak) verificarEstado(placar placares.EstadoEParametroPlacar) error {
	if j.Terminado() {
		return errors.New("Tie Break já encerrado.")
	}

	if placar.Tipo() != placares.TPPonto {
		return errors.New("Este placar não é de ponto.")
	}

	if !placar.Terminado() {
		return errors.New("Este ponto não foi encerrado.")
	}

	return nil
}

func (j *TieBreak) AdicionarPlacar(placarPonto placares.EstadoEParametroPlacar) error {
	if err := j.verificarEstado(placarPonto); err != nil {
		return err
	}

	sA, sB := j.placares()
	// Padronizando que o lado A sempre fica com o sacador.
	// Então quando o inicio for o lado B, o sA será sB e vice-versa.
	if j.QuemEstaSacando(0) == turnos.LTB {
		sA, sB = sB, sA
	}

	adicionarEm := sA
	if who := placarPonto.Lado(); who == placares.LPOposto {
		adicionarEm = sB
	}

	*adicionarEm += 1
	j.executarEventosAoMudarPlacar()

	return nil
}

func (j *TieBreak) AdicionaAoMudarPlacar(event AoMudarPlacar) {
	j.eventosAoMudarPlacar = append(j.eventosAoMudarPlacar, event)
}

func (j *TieBreak) DefinirTestaEncerramento(teste TestaEncerramento) {
	j.testaEncerramento = teste
}

func (j TieBreak) Resultado() (int, int) {
	A, B := j.placares()
	return *A, *B
}

func (j TieBreak) Lado() placares.LadoDoPlacar {
	return placares.Lado(j)
}

func (j TieBreak) Tipo() placares.TipoDoPlacar {
	return placares.TPJogo
}

func (j TieBreak) Terminado() bool {
	placarA, placaB, pontoDecisivo := j.placarA, j.placarB, 0
	if j.pontoDecisivo {
		pontoDecisivo = 1
	}

	if j.testaEncerramento != nil {
		return j.testaEncerramento(placarA, placaB, pontoDecisivo)
	}

	return false
}
