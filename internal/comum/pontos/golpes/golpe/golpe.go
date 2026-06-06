package golpe

import (
	"github.com/cirobispo/sandbox/internal/comum/pontos/golpes"
)

type Golpe struct {
	tipo golpes.TipoDoGolpe
	lado golpes.LadoDoGolpe
}

func (g Golpe) Tipo() golpes.TipoDoGolpe {
	return g.tipo
}

func (g Golpe) Lado() golpes.LadoDoGolpe {
	return g.lado
}

func New(t golpes.TipoDoGolpe, s golpes.LadoDoGolpe) Golpe {
	return Golpe{
		tipo: t,
		lado: s,
	}
}

func NewFootFault() Golpe {
	return Golpe{
		tipo: golpes.HTFootFault,
		lado: golpes.HTDConditional,
	}
}

func NewAce() Golpe {
	return Golpe{
		tipo: golpes.HTAce,
		lado: golpes.HTDSameSide,
	}
}

func NewServeOut() Golpe {
	return Golpe{
		tipo: golpes.HTServeOut,
		lado: golpes.HTDConditional,
	}
}

func NewServeIn() Golpe {
	return Golpe{
		tipo: golpes.HTServeIn,
		lado: golpes.HTDChangeSide,
	}
}

func NewServeLet() Golpe {
	return Golpe{
		tipo: golpes.HTServeLet,
		lado: golpes.HTDNone,
	}
}

func NewServeNet() Golpe {
	return Golpe{
		tipo: golpes.HTServeNet,
		lado: golpes.HTDConditional,
	}
}

func NewReturnNet() Golpe {
	return Golpe{
		tipo: golpes.HTReturnNet,
		lado: golpes.HTDOppositeSide,
	}
}

func NewReturnIn() Golpe {
	return Golpe{
		tipo: golpes.HTReturnIn,
		lado: golpes.HTDChangeSide,
	}
}

func NewReturnOut() Golpe {
	return Golpe{
		tipo: golpes.HTReturnOut,
		lado: golpes.HTDOppositeSide,
	}
}

func NewDoubleFault() Golpe {
	return Golpe{
		tipo: golpes.HTDoubleFault,
		lado: golpes.HTDOppositeSide,
	}
}

func NewHitNet() Golpe {
	return Golpe{
		tipo: golpes.HTNet,
		lado: golpes.HTDOppositeSide,
	}
}

func NewHitBackIn() Golpe {
	return Golpe{
		tipo: golpes.HTIn,
		lado: golpes.HTDChangeSide,
	}
}

func NewWinner() Golpe {
	return Golpe{
		tipo: golpes.HTWinner,
		lado: golpes.HTDSameSide,
	}
}

func NewHitOut() Golpe {
	return Golpe{
		tipo: golpes.HTOut,
		lado: golpes.HTDOppositeSide,
	}
}

func NewMiss() Golpe {
	return Golpe{
		tipo: golpes.HTOut,
		lado: golpes.HTDOppositeSide,
	}
}

func NewToast() Golpe {
	return Golpe{
		tipo: golpes.HTToast,
		lado: golpes.HTDSameSide,
	}
}

func NewNetTouch() Golpe {
	return Golpe{
		tipo: golpes.HTNetTouch,
		lado: golpes.HTDOppositeSide,
	}
}
