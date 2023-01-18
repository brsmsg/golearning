package main

import (
	"log"
	"net"
	"time"
)

func main() {
	service := ":13000" // port

	listener, err := net.Listen("tcp", service)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			// prevent crash
			continue
		}

		daytime := time.Now().String()

		_, err = conn.Write([]byte(daytime))
		if err != nil {
			conn.Close()
			continue
		}
		conn.Close()
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatalf(err.Error())
	}
}
