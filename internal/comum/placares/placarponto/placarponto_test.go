package placarponto

import (
	"testing"

	"github.com/cirobispo/sandbox/internal/comum/golpes"
	"github.com/cirobispo/sandbox/internal/comum/golpes/golpe"
	"github.com/cirobispo/sandbox/internal/comum/pontos/ponto"
	"github.com/cirobispo/sandbox/internal/comum/turnos"
	"github.com/cirobispo/sandbox/internal/comum/turnos/turno"
)

var ladoInicial turnos.LadoDoTurno

func Test_PontoRecebedor(t *testing.T) {
	ladoInicial = turnos.LTA
	AvaliarResultado(t, NovoPlacarPonto(t, Ace(t, ladoInicial)))

	AvaliarResultado(t, NovoPlacarPonto(t, ServicoPlus1(t, ladoInicial)))

	AvaliarResultado(t, NovoPlacarPonto(t, ServicoEWinner(t, ladoInicial)))

	ladoInicial = turnos.LTB
	AvaliarResultado(t, NovoPlacarPonto(t, Ace(t, ladoInicial)))

	AvaliarResultado(t, NovoPlacarPonto(t, ServicoPlus1(t, ladoInicial)))

	AvaliarResultado(t, NovoPlacarPonto(t, ServicoEWinner(t, ladoInicial)))
}
func ConfiguraTurno(t *testing.T, tu *turno.Turno) {
	tu.AdicionarDepoisDeMudarTurno(func(ldt turnos.LadoDoTurno) {
		t.Log("Lado corrente da bola: ", ldt)
	})
}

func alertarGolpe(t *testing.T, ponto *ponto.Ponto) {
	ponto.AdicionarEventoAoAdicionarGolpe(func(tipoDoGolpe golpes.TipoDoGolpe, terminado bool) {
		t.Log("Golpe:", tipoDoGolpe, ", Encerrado:", terminado)
	})
}

func Ace(t *testing.T, lado turnos.LadoDoTurno) *ponto.Ponto {
	// Ace
	tu := turno.New(turno.DefinindoLado(lado))
	ConfiguraTurno(t, tu)
	p := ponto.New(tu)
	alertarGolpe(t, &p)
	p.AdicionarGolpe(golpe.NewAce())

	return &p
}

func ServicoPlus1(t *testing.T, lado turnos.LadoDoTurno) *ponto.Ponto {
	// Serviço + 1
	tu := turno.New(turno.DefinindoLado(lado))
	ConfiguraTurno(t, tu)
	p := ponto.New(tu)
	alertarGolpe(t, &p)

	p.AdicionarGolpe(golpe.NewServicoDentro())
	p.AdicionarGolpe(golpe.NewRetornoDentro())
	p.AdicionarGolpe(golpe.NewWinner())

	return &p
}

func ServicoEWinner(t *testing.T, lado turnos.LadoDoTurno) *ponto.Ponto {
	// Serviço e winner
	tu := turno.New(turno.DefinindoLado(lado))
	ConfiguraTurno(t, tu)
	p := ponto.New(tu)
	alertarGolpe(t, &p)

	p.AdicionarGolpe(golpe.NewServicoDentro())
	p.AdicionarGolpe(golpe.NewWinner())

	return &p
}

func NovoPlacarPonto(t *testing.T, ponto *ponto.Ponto) Ponto {
	t.Log("Lado inicial da bola:", ladoInicial, ", Lado do ponto:", ponto.LadoDoPonto(), ", Lado da bola:", ponto.LadoDaBola())

	placar := New(ponto)
	pA, pB := placar.Resultado()
	t.Log("Resultado: ", pA, pB)
	return placar
}

func AvaliarResultado(t *testing.T, placar Ponto) {
}
