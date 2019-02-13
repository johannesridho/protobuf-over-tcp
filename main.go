package main

import (
	"github.com/golang/protobuf/proto"
	"github.com/johannesridho/protobuf-over-tcp/message"
	"log"
	"net"
	"os"
)

func main() {
	connType := os.Getenv("TYPE")

	if connType == "server" {
		startServer()
	} else if connType == "client" {
		startClient()
	} else {
		log.Println("please provide the TYPE arg (TYPE=server or TYPE=client)")
	}
}

func startServer() {
	log.Println("starting tcp server")
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	checkError(err)

	for {
		if conn, err := listener.Accept(); err == nil {
			handleConn(conn)
		}
	}
}

func handleConn(conn net.Conn) {
	log.Println("client connected")

	defer conn.Close()

	messageProto := message.Message{Text: "Hello World"}
	data, err := proto.Marshal(&messageProto)
	checkError(err)

	length, err := conn.Write(data)
	checkError(err)

	log.Printf("Hello world sent, length %d bytes", length)
}

func startClient() {
	log.Println("starting tcp client")

	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	checkError(err)

	defer conn.Close()

	data := make([]byte, 4096)
	length, err := conn.Read(data)
	checkError(err)

	messageProto := message.Message{}
	err = proto.Unmarshal(data[:length], &messageProto)
	checkError(err)

	log.Printf("received message : %s", messageProto.Text)
}

func checkError(err error) {
	if err != nil {
		log.Println(err.Error())
	}
}
