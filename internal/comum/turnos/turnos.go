package turnos

type Lado int

const (
	LadoA Lado = 0
	LadoB Lado = 1
)

type Turning interface {
	Execute()
	LadoInicial() Lado
	LadoCorrente() Lado
}

type AoMudarTurno func(Lado)

func (l Lado) String() string {
	result := "A"
	if l == LadoB {
		result = "B"
	}

	return result
}

func (l Lado) Inverso() Lado {
	result := LadoB
	if l == LadoB {
		result = LadoA
	}
	return result
}
