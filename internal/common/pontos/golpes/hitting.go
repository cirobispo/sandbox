package golpes

type TipoDoGolpe int
type LadoDoGolpe int

const (
	HTFootFault   TipoDoGolpe = 1
	HTServeNet    TipoDoGolpe = 2
	HTServeLet    TipoDoGolpe = 3
	HTServeIn     TipoDoGolpe = 4
	HTAce         TipoDoGolpe = 5
	HTServeOut    TipoDoGolpe = 6
	HTReturnOut   TipoDoGolpe = 7
	HTReturnNet   TipoDoGolpe = 8
	HTReturnIn    TipoDoGolpe = 9
	HTDoubleFault TipoDoGolpe = 10
	HTNet         TipoDoGolpe = 11
	HTIn          TipoDoGolpe = 12
	HTOut         TipoDoGolpe = 13
	HTWinner      TipoDoGolpe = 14
	HTToast       TipoDoGolpe = 15
	HTNetTouch    TipoDoGolpe = 16
	HTMiss        TipoDoGolpe = 17
)

const (
	HTDNone         LadoDoGolpe = 0
	HTDChangeSide   LadoDoGolpe = 1
	HTDSameSide     LadoDoGolpe = 2
	HTDOppositeSide LadoDoGolpe = 3
	HTDConditional  LadoDoGolpe = 4
)

type Hitting interface {
	Tipo() TipoDoGolpe
	Lado() LadoDoGolpe
}

func (h LadoDoGolpe) Inverso() LadoDoGolpe {
	switch h {
	case HTDSameSide:
		return HTDOppositeSide
	case HTDOppositeSide:
		return HTDSameSide
	default:
		return h
	}
}

func (h TipoDoGolpe) String() string {
	switch h {
	case HTFootFault:
		return "Foot fault"
	case HTServeNet:
		return "Serve on net"
	case HTServeLet:
		return "Let"
	case HTServeIn:
		return "Serve in"
	case HTAce:
		return "Ace!"
	case HTServeOut:
		return "Serve out"
	case HTReturnOut:
		return "Return out"
	case HTReturnNet:
		return "Return net"
	case HTReturnIn:
		return "Return in"
	case HTDoubleFault:
		return "Double fault"
	case HTNet:
		return "Hit net"
	case HTIn:
		return "Hit in"
	case HTOut:
		return "Hit out"
	case HTWinner:
		return "Winner!"
	case HTToast:
		return "Toast!"
	case HTNetTouch:
		return "Touch net"
	case HTMiss:
		return "Miss"
	default:
		return "Other"
	}
}

func (h LadoDoGolpe) String() string {
	switch h {
	case HTDNone:
		return "None"
	case HTDChangeSide:
		return "Change Side"
	case HTDSameSide:
		return "Same side"
	case HTDOppositeSide:
		return "Opposite side"
	case HTDConditional:
		return "Conditional"
	default:
		return "other"
	}
}
