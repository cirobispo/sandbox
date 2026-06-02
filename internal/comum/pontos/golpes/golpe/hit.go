package golpe

import (
	"github.com/cirobispo/sandbox/internal/comum/pontos/golpes"
)

type Hit struct {
	tipo golpes.TipoDoGolpe
	lado golpes.LadoDoGolpe
}

func (h Hit) Tipo() golpes.TipoDoGolpe {
	return h.tipo
}

func (h Hit) Lado() golpes.LadoDoGolpe {
	return h.lado
}

func New(t golpes.TipoDoGolpe, s golpes.LadoDoGolpe) Hit {
	return Hit{
		tipo: t,
		lado: s,
	}
}

func NewFootFault() Hit {
	return Hit{
		tipo: golpes.HTFootFault,
		lado: golpes.HTDConditional,
	}
}

func NewAce() Hit {
	return Hit{
		tipo: golpes.HTAce,
		lado: golpes.HTDSameSide,
	}
}

func NewServeOut() Hit {
	return Hit{
		tipo: golpes.HTServeOut,
		lado: golpes.HTDConditional,
	}
}

func NewServeIn() Hit {
	return Hit{
		tipo: golpes.HTServeIn,
		lado: golpes.HTDChangeSide,
	}
}

func NewServeLet() Hit {
	return Hit{
		tipo: golpes.HTServeLet,
		lado: golpes.HTDNone,
	}
}

func NewServeNet() Hit {
	return Hit{
		tipo: golpes.HTServeNet,
		lado: golpes.HTDConditional,
	}
}

func NewReturnNet() Hit {
	return Hit{
		tipo: golpes.HTReturnNet,
		lado: golpes.HTDOppositeSide,
	}
}

func NewReturnIn() Hit {
	return Hit{
		tipo: golpes.HTReturnIn,
		lado: golpes.HTDChangeSide,
	}
}

func NewReturnOut() Hit {
	return Hit{
		tipo: golpes.HTReturnOut,
		lado: golpes.HTDOppositeSide,
	}
}

func NewDoubleFault() Hit {
	return Hit{
		tipo: golpes.HTDoubleFault,
		lado: golpes.HTDOppositeSide,
	}
}

func NewHitNet() Hit {
	return Hit{
		tipo: golpes.HTNet,
		lado: golpes.HTDOppositeSide,
	}
}

func NewHitBackIn() Hit {
	return Hit{
		tipo: golpes.HTIn,
		lado: golpes.HTDChangeSide,
	}
}

func NewWinner() Hit {
	return Hit{
		tipo: golpes.HTWinner,
		lado: golpes.HTDSameSide,
	}
}

func NewHitOut() Hit {
	return Hit{
		tipo: golpes.HTOut,
		lado: golpes.HTDOppositeSide,
	}
}

func NewMiss() Hit {
	return Hit{
		tipo: golpes.HTOut,
		lado: golpes.HTDOppositeSide,
	}
}

func NewToast() Hit {
	return Hit{
		tipo: golpes.HTToast,
		lado: golpes.HTDSameSide,
	}
}

func NewNetTouch() Hit {
	return Hit{
		tipo: golpes.HTNetTouch,
		lado: golpes.HTDOppositeSide,
	}
}
