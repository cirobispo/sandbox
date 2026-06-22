package turnos

type LadoDoTurno int

const (
	LTA LadoDoTurno = 0
	LTB LadoDoTurno = 1
)

type Turning interface {
	Execute()
	LadoInicial() LadoDoTurno
	LadoCorrente() LadoDoTurno
}

type AoMudarTurno func(LadoDoTurno)

func (s LadoDoTurno) String() string {
	result := "Lado A"
	if s == LTB {
		result = "Lado B"
	}

	return result
}
