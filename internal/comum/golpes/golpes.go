package golpes

type TipoDoGolpe int
type TipoAcao int

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
	TANulo        TipoAcao = 0
	TAProsseguir  TipoAcao = 1
	TAEncerrarPLC TipoAcao = 2
	TAEncerrarPLO TipoAcao = 3
	TACondicional TipoAcao = 4
)

type Golpeando interface {
	Tipo() TipoDoGolpe
	Acao() TipoAcao
}

func (t TipoDoGolpe) String() string {
	switch t {
	case HTFootFault:
		return "Foot fault"
	case HTServeNet:
		return "Serviço na rede"
	case HTServeLet:
		return "Let"
	case HTServeIn:
		return "Serviço dentro"
	case HTAce:
		return "Ace!"
	case HTServeOut:
		return "Serviço fora"
	case HTReturnOut:
		return "Retorno fora"
	case HTReturnNet:
		return "Retorno rede"
	case HTReturnIn:
		return "Retorno dentro"
	case HTDoubleFault:
		return "Dupla falta"
	case HTNet:
		return "Devolveu na rede"
	case HTIn:
		return "Devolveu dentro"
	case HTOut:
		return "Devolveu foras"
	case HTWinner:
		return "Winner!"
	case HTToast:
		return "Bola no jogador"
	case HTNetTouch:
		return "Tocou a rede"
	case HTMiss:
		return "Não tocou"
	default:
		return "Other"
	}
}
