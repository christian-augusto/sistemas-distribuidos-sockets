package operations

import (
	"fmt"
	"log"
	"sistemas-distribuidos-sockets-client-golang/outputs"
	"sistemas-distribuidos-sockets-client-golang/utils"
)

func CreateClient(socketRequestPayload *outputs.SocketRequestPayload) error {
	var err error

	log.Print("Type the client name: ")
	socketRequestPayload.Name, err = utils.ReadFromKeyboard()

	if err != nil {
		return fmt.Errorf("Error reading client name: %v\n", err)
	}

	log.Print("Type the client cpf: ")
	socketRequestPayload.Cpf, err = utils.ReadFromKeyboard()

	if err != nil {
		return fmt.Errorf("Error reading client cpf: %v\n", err)
	}

	log.Print("Type the client address: ")
	socketRequestPayload.Address, err = utils.ReadFromKeyboard()

	if err != nil {
		return fmt.Errorf("Error reading client address: %v\n", err)
	}

	return err
}
