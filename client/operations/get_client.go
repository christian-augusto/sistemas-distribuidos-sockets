package operations

import (
	"fmt"
	"log"
	"sistemas-distribuidos-sockets-client-golang/outputs"
	"sistemas-distribuidos-sockets-client-golang/utils"
)

func GetClient(socketRequestPayload *outputs.SocketRequestPayload) error {
	var err error

	log.Print("Type the client name: ")
	socketRequestPayload.Name, err = utils.ReadFromKeyboard()

	if err != nil {
		return fmt.Errorf("Error reading client name: %v\n", err)
	}

	return err
}
