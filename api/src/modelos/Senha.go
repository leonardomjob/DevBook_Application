package modelos

// Senha representa o formatado da requisição de alterção de senha
type Senha struct {
	Nova  string `json:"nova"`
	Atual string `json:"atual"`
}
