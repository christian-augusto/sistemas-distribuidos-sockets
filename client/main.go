package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"sistemas-distribuidos-sockets-client-golang/config"
	"sistemas-distribuidos-sockets-client-golang/constants"
	"sistemas-distribuidos-sockets-client-golang/operations"
	"sistemas-distribuidos-sockets-client-golang/outputs"
	"sistemas-distribuidos-sockets-client-golang/utils"

	"github.com/spf13/viper"
)

func main() {
	config.LoadEnvironmentVars(os.Getenv(constants.ENV_NAME))

	var err error
	var serverConnection net.Conn
	var requestPayloadBytes []byte
	var responsePayloadBytes []byte
	socketRequestPayload := outputs.NewSocketRequestPayload()
	server_socket_host := viper.GetString(constants.SERVER_SOCKET_HOST_NAME)
	server_socket_port := viper.GetString(constants.SERVER_SOCKET_PORT_NAME)
	server_socket_type := viper.GetString(constants.SERVER_SOCKET_TYPE_NAME)

	serverConnection, err = net.Dial(server_socket_type, fmt.Sprintf("%v:%v", server_socket_host, server_socket_port))

	if err != nil {
		log.Printf("Error connecting: %v\n", err)
	}

	defer serverConnection.Close()

	for {
		log.Print("Type the operation type (create, delete or get): ")
		socketRequestPayload.Type, err = utils.ReadFromKeyboard()

		if err != nil {
			log.Printf("Error reading operation type: %v\n", err)
			continue
		}

		switch socketRequestPayload.Type {
		case constants.SOCKET_PAYLOAD_OPERATION_TYPE_CREATE:
			err = operations.CreateClient(socketRequestPayload)
		case constants.SOCKET_PAYLOAD_OPERATION_TYPE_DELETE:
			err = operations.DeleteClient(socketRequestPayload)
		case constants.SOCKET_PAYLOAD_OPERATION_TYPE_GET:
			err = operations.GetClient(socketRequestPayload)
		default:
			log.Printf("Invalid operation type, try again")
			continue
		}

		if err != nil {
			log.Printf("Error in get values to operation: %v\n", err)
			continue
		}

		requestPayloadBytes, err = json.Marshal(socketRequestPayload)
		requestPayloadBytes = append(requestPayloadBytes, []byte("\n")...)

		if err != nil {
			log.Printf("Error in creates json request: %v\n", err)
			continue
		}

		_, err = serverConnection.Write(requestPayloadBytes)

		if err != nil {
			log.Printf("Error in write response: %v\n", err)
			continue
		}

		responsePayloadBytes, err = bufio.NewReader(serverConnection).ReadBytes('\n')

		if err != nil {
			log.Printf("Error in read server response: %v\n", err)
			return
		}

		if viper.GetString(constants.LOGS_NAME) == constants.BOOLEAN_TRUE_ENV_VALUE {
			log.Printf("Server responsePayloadBytes len: %v\n", len(responsePayloadBytes))
		}

		if viper.GetString(constants.LOGS_NAME) == constants.BOOLEAN_TRUE_ENV_VALUE {
			log.Printf("Server response: %v\n", string(responsePayloadBytes))
		}
	}
}
