package placarponto

import (
	"testing"

	"github.com/cirobispo/sandbox/internal/comum/golpes/golpe"
	"github.com/cirobispo/sandbox/internal/comum/pontos"
	"github.com/cirobispo/sandbox/internal/comum/pontos/ponto"
	"github.com/cirobispo/sandbox/internal/comum/turnos"
	"github.com/cirobispo/sandbox/internal/comum/turnos/turno"
)

var ladoInicial turnos.LadoDoTurno

func Test_PontoRecebedor(t *testing.T) {
	ladoInicial = turnos.LTA
	AvaliarPlacarPonto(t, Ace(t, ladoInicial), pontos.LPCorrente)

	AvaliarPlacarPonto(t, ServicoPlus1(t, ladoInicial), pontos.LPCorrente)

	AvaliarPlacarPonto(t, ServicoEWinner(t, ladoInicial), pontos.LPCorrente)

	ladoInicial = turnos.LTB
	AvaliarPlacarPonto(t, Ace(t, ladoInicial), pontos.LPCorrente)

	AvaliarPlacarPonto(t, ServicoPlus1(t, ladoInicial), pontos.LPCorrente)

	AvaliarPlacarPonto(t, ServicoEWinner(t, ladoInicial), pontos.LPCorrente)
}

func ConfiguraTurno(t *testing.T, tu *turno.Turno) {
	// tu.AdicionarDepoisDeMudarTurno(func(ldt turnos.LadoDoTurno) {
	// 	t.Log("Lado corrente da bola: ", ldt)
	// })
}

func alertarGolpe(t *testing.T, ponto *ponto.Ponto) {
	// ponto.AdicionarEventoAoAdicionarGolpe(func(tipoDoGolpe golpes.TipoDoGolpe, terminado bool) {
	// 	t.Log("Golpe:", tipoDoGolpe, ", Encerrado:", terminado)
	// })
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

func AvaliarPlacarPonto(t *testing.T, p *ponto.Ponto, ladoDoPonto pontos.LadoDoPonto) PlacarPonto {
	t.Logf("Lado inicial da bola: %v , Lado do ponto: %v, Lado esperado: %v", ladoInicial, p.LadoDoPonto(), ladoDoPonto) //, ", Lado da bola:", p.LadoDaBola())

	placar := New(p, ladoInicial, 3)
	pA, pB := placar.Resultado()
	t.Log("Resultado: ", pA, pB)
	if p.LadoDoPonto() != ladoDoPonto {
		t.Errorf("Resultado não é o esperado. Resultado %v, esperado: %v ", p.LadoDoPonto(), ladoDoPonto)
	}
	return placar
}
