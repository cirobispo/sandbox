package placares

type LadoDoPlacar int

const (
	LPNulo    LadoDoPlacar = 0
	LPServico LadoDoPlacar = 1
	LPOposto  LadoDoPlacar = 2
)

type TipoDoPlacar int

const (
	TPPonto   TipoDoPlacar = 0
	TPJogo    TipoDoPlacar = 1
	TPSet     TipoDoPlacar = 2
	TPPartida TipoDoPlacar = 3
)

type ParametroDoPlacar interface {
	Lado() LadoDoPlacar
	Tipo() TipoDoPlacar
}

type EstadoDoPlacar interface {
	Terminado() bool
}

type ResultadoDoPlacar interface {
	Resultado() (int, int)
}

type EstadoEResultadoPlacar interface {
	EstadoDoPlacar
	ResultadoDoPlacar
}

type EstadoEParametroPlacar interface {
	EstadoDoPlacar
	ParametroDoPlacar
}

type AdicionadorDePlacar interface {
	AdicionaPlacar(epp EstadoEParametroPlacar) error
}

type EstadoResultadoEParametroPlacar interface {
	EstadoDoPlacar
	ResultadoDoPlacar
	ParametroDoPlacar
}

type EstadoResultadoParametroEAdicionadorPlacar interface {
	EstadoDoPlacar
	ResultadoDoPlacar
	ParametroDoPlacar
	AdicionadorDePlacar
}

func Lado(erp EstadoEResultadoPlacar) LadoDoPlacar {
	if !erp.Terminado() {
		return LPNulo
	}

	if a, b := erp.Resultado(); b > a {
		return LPOposto
	}

	return LPServico
}

func TraduzirPlacar(valores []string, placarA, placarB int) (string, string) {
	if len(valores) != 6 {
		panic("Descrição dos pontos incorreta (0, 15, 30, 40, vantagem e jogo)")
	}

	getText := func(placar, b int) string {
		if placar == 5 || placar == 4 && b < 3 {
			return valores[len(valores)-1]
		}
		return valores[placar]
	}

	return getText(placarA, placarB), getText(placarB, placarA)
}
