package database

import (
	"container/list"
	"sistemas-distribuidos-sockets-server-golang/models"
	"strings"
)

type MemoryDatabase struct{}

func NewMemoryDatabase() *MemoryDatabase {
	return &MemoryDatabase{}
}

var clients *list.List

func (md *MemoryDatabase) CreateClient(name string, cpf string, address string) *models.Client {
	startDatabase()

	client := models.NewClient(name, cpf, address)

	clients.PushBack(client)

	return client
}

func (md *MemoryDatabase) DeleteClientByName(name string) *models.Client {
	startDatabase()

	if clients.Len() > 0 {
		for elem := clients.Front(); elem != nil; elem = elem.Next() {
			client := elem.Value.(*models.Client)

			if strings.Contains(client.Name, name) {
				clients.Remove(elem)
				return client
			}
		}
	}

	return nil
}

func (md *MemoryDatabase) GetClientByName(name string) *models.Client {
	startDatabase()

	if clients.Len() > 0 {
		for elem := clients.Front(); elem != nil; elem = elem.Next() {
			client := elem.Value.(*models.Client)

			if strings.Contains(client.Name, name) {
				return client
			}
		}
	}

	return nil
}

func startDatabase() {
	if clients == nil {
		clients = list.New()
	}
}
