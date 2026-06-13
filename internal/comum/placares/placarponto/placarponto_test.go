package placarponto

import (
	"testing"

	"github.com/cirobispo/sandbox/internal/comum/placares"
)

func adicionarPontoNoPlacar(placar *PlacarPonto, lado placares.LadoDoPlacar) {
	if lado != placares.LPNulo {
		placar.Zerar()
		if lado == placar.Lado() {
			placar.PontuarA()
		}
	}
}

func Test_PontoServico(t *testing.T) {
	placar := New()
	adicionarPontoNoPlacar(&placar, placares.LPServico)
	if placar.Terminado() {
	}
}

func Test_PontoRecebedor(t *testing.T) {

}
