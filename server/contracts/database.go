package contracts

import "sistemas-distribuidos-sockets-server-golang/models"

type Database interface {
	CreateClient(name string, cpf string, address string) *models.Client
	DeleteClientByName(name string) *models.Client
	GetClientByName(name string) *models.Client
}
