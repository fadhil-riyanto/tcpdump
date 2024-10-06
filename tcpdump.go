package main

import (
	"log"
	"net"
	"os"
)

const (
	HOST = "localhost"
	PORT = "8080"
	TYPE = "tcp"
)

func handleRequest(conn net.Conn) {
	buff := make([]byte, 50)
	for {
		_, err := conn.Read(buff)

		if err != nil {
			println("Read data failed:", err.Error())
			// os.Exit(1)
		}

		log.Printf("[ 0x%01x 0x%01x 0x%01x 0x%01x 0x%01x 0x%01x 0x%01x 0x%01x ]", buff[0], buff[1], buff[2], buff[3], buff[4], buff[5], buff[6], buff[7])
	}

}

func main() {
	listen, err := net.Listen(TYPE, os.Args[1]+":"+os.Args[2])
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	// close listener
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		go handleRequest(conn)
	}
}
