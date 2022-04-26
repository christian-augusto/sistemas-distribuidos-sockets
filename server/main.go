package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"sistemas-distribuidos-sockets-server-golang/config"
	"sistemas-distribuidos-sockets-server-golang/constants"
	"sistemas-distribuidos-sockets-server-golang/contracts"
	"sistemas-distribuidos-sockets-server-golang/infra/database"
	"sistemas-distribuidos-sockets-server-golang/inputs"
	"sistemas-distribuidos-sockets-server-golang/models"
	"sistemas-distribuidos-sockets-server-golang/outputs"

	"github.com/spf13/viper"
)

func main() {
	config.LoadEnvironmentVars(os.Getenv(constants.ENV_NAME))

	var err error
	var server net.Listener
	var clientConnection net.Conn
	server_socket_host := viper.GetString(constants.SERVER_SOCKET_HOST_NAME)
	server_socket_port := viper.GetString(constants.SERVER_SOCKET_PORT_NAME)
	server_socket_type := viper.GetString(constants.SERVER_SOCKET_TYPE_NAME)

	server, err = net.Listen(server_socket_type, fmt.Sprintf("%v:%v", server_socket_host, server_socket_port))

	if err != nil {
		log.Fatalln(err)
	}

	if viper.GetString(constants.LOGS_NAME) == constants.BOOLEAN_TRUE_ENV_VALUE {
		log.Printf("Server started at: %v:%v type %v\n", server_socket_host, server_socket_port, server_socket_type)
	}

	for {
		clientConnection, err = server.Accept()

		if err != nil {
			log.Printf("Error connecting: %v\n", err)
			continue
		}

		if viper.GetString(constants.LOGS_NAME) == constants.BOOLEAN_TRUE_ENV_VALUE {
			log.Println("Client connected")
		}

		go handleClientConnection(clientConnection)
	}
}

func handleClientConnection(clientConnection net.Conn) {
	for {
		var err error
		var requestPayloadBytes []byte
		var responsePayloadBytes []byte

		requestPayloadBytes, err = bufio.NewReader(clientConnection).ReadBytes('\n')

		if err != nil {
			log.Printf("Error in read client request: %v\n", err)
			return
		}

		if viper.GetString(constants.LOGS_NAME) == constants.BOOLEAN_TRUE_ENV_VALUE {
			log.Printf("Client requestPayloadBytes len: %v\n", len(requestPayloadBytes))
		}

		if viper.GetString(constants.LOGS_NAME) == constants.BOOLEAN_TRUE_ENV_VALUE {
			log.Printf("Client request: %v\n", string(requestPayloadBytes))
		}

		requestPayload := inputs.NewSocketRequestPayload()

		err = json.Unmarshal(requestPayloadBytes, requestPayload)

		if err != nil {
			log.Printf("Error in parse client request: %v\n", err)
			return
		}

		var db contracts.Database = database.NewMemoryDatabase()

		var client *models.Client

		switch requestPayload.Type {
		case constants.SOCKET_PAYLOAD_OPERATION_TYPE_CREATE:
			client = db.CreateClient(requestPayload.Name, requestPayload.Cpf, requestPayload.Address)
		case constants.SOCKET_PAYLOAD_OPERATION_TYPE_DELETE:
			client = db.DeleteClientByName(requestPayload.Name)
		case constants.SOCKET_PAYLOAD_OPERATION_TYPE_GET:
			client = db.GetClientByName(requestPayload.Name)
		}

		responsePayload := outputs.NewSocketResponsePayload(client)

		responsePayloadBytes, err = json.Marshal(responsePayload)
		responsePayloadBytes = append(responsePayloadBytes, []byte("\n")...)

		if err != nil {
			log.Printf("Error in creates json response: %v\n", err)
			return
		}

		if viper.GetString(constants.LOGS_NAME) == constants.BOOLEAN_TRUE_ENV_VALUE {
			log.Printf("responsePayload: %v\n", string(responsePayloadBytes))
		}

		clientConnection.Write(responsePayloadBytes)
	}
}
