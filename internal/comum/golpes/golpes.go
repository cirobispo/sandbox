package golpes

type Acao int
type Reacao int

const (
	AServico         Acao = 20
	APeNaQuadra      Acao = 21
	ALet             Acao = 22
	AServicoNaRede   Acao = 23
	AServicoDentro   Acao = 24
	AServicoFora     Acao = 25
	AAce             Acao = 26
	ADuplaFalta      Acao = 27
	ARetorno         Acao = 30
	ARetornoNaRede   Acao = 31
	ARetornoDentro   Acao = 32
	ARetornoFora     Acao = 33
	ADevolucao       Acao = 40
	ADevolucaoNaRede Acao = 41
	ADevolucaoDentro Acao = 42
	ADevolucaoFora   Acao = 43
	AWinner          Acao = 140
	AToqueNaBola     Acao = 150
	AToqueNaRede     Acao = 160
	ANaoTocouNaBola  Acao = 170
)

const (
	RNulo           Reacao = 0
	RProsseguir     Reacao = 1
	REPLadoCorrente Reacao = 2
	REPLadoOposto   Reacao = 3
	RPCondicionado  Reacao = 4
)

type Golpear interface {
	Acao() Acao
	Reacao() Reacao
	ExecutaTurno() bool
}

func (t Acao) String() string {
	switch t {
	case APeNaQuadra:
		return "Foot fault"
	case AServicoNaRede:
		return "Serviço na rede"
	case ALet:
		return "Let"
	case AServicoDentro:
		return "Serviço dentro"
	case AAce:
		return "Ace!"
	case AServicoFora:
		return "Serviço fora"
	case ARetornoFora:
		return "Retorno fora"
	case ARetornoNaRede:
		return "Retorno rede"
	case ARetornoDentro:
		return "Retorno dentro"
	case ADuplaFalta:
		return "Dupla falta"
	case ADevolucaoNaRede:
		return "Devolveu na rede"
	case ADevolucaoDentro:
		return "Devolveu dentro"
	case ADevolucaoFora:
		return "Devolveu foras"
	case AWinner:
		return "Winner!"
	case AToqueNaBola:
		return "Bola no jogador"
	case AToqueNaRede:
		return "Tocou a rede"
	case ANaoTocouNaBola:
		return "Não tocou"
	default:
		return "Other"
	}
}
