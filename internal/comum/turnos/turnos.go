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

func (l LadoDoTurno) String() string {
	result := "Lado A"
	if l == LTB {
		result = "Lado B"
	}

	return result
}

func (l LadoDoTurno) Inverso() LadoDoTurno {
	result := LTB
	if l == LTB {
		result = LTA
	}
	return result
}
