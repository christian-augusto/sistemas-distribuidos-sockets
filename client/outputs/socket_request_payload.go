package outputs

type SocketRequestPayload struct {
	Type    string `json:"type"`
	Name    string `json:"name"`
	Cpf     string `json:"cpf"`
	Address string `json:"address"`
}

func NewSocketRequestPayload() *SocketRequestPayload {
	return &SocketRequestPayload{}
}
