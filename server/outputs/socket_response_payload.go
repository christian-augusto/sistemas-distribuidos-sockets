package outputs

import "sistemas-distribuidos-sockets-server-golang/models"

type SocketResponsePayload struct {
	Client *models.Client `json:"client"`
}

func NewSocketResponsePayload(client *models.Client) *SocketResponsePayload {
	return &SocketResponsePayload{
		Client: client,
	}
}
