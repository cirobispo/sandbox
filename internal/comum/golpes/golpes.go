package golpes

type Acao int
type Reacao int

const (
	HTFootFault   Acao = 1
	HTServeNet    Acao = 2
	HTServeLet    Acao = 3
	HTServeIn     Acao = 4
	HTAce         Acao = 5
	HTServeOut    Acao = 6
	HTReturnOut   Acao = 7
	HTReturnNet   Acao = 8
	HTReturnIn    Acao = 9
	HTDoubleFault Acao = 10
	HTNet         Acao = 11
	HTIn          Acao = 12
	HTOut         Acao = 13
	HTWinner      Acao = 14
	HTToast       Acao = 15
	HTNetTouch    Acao = 16
	HTMiss        Acao = 17
)

const (
	TANulo        Reacao = 0
	TAProsseguir  Reacao = 1
	TAEncerrarPLC Reacao = 2
	TAEncerrarPLO Reacao = 3
	TACondicional Reacao = 4
)

type Golpear interface {
	Acao() Acao
	Reacao() Reacao
	ExecutaTurno() bool
}

func (t Acao) String() string {
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
