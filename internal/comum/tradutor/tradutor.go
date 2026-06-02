package tradutor

type MensagemAExibir struct {
	idioma int
	msgs   map[string]string
}

func New(idioma int) MensagemAExibir {
	return MensagemAExibir{idioma: idioma, msgs: make(map[string]string)}
}

func (m MensagemAExibir) RecuperarMensagem(nome string) any {
	return &[]string{}
}
