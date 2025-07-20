package main

import (
	"fmt"
	"keeva-db/utils"
	"net"
)

func main () {
	port := ":6379"
	fmt.Println("Listening on port",port)
	listener, err:= net.Listen("tcp",port)
	utils.ParseError(err,false)
	connection, err := listener.Accept()
	utils.ParseError(err,false)
	defer connection.Close()

	for {
		buffer := make([]byte,1024)
		_,err := connection.Read(buffer)
		utils.ParseError(err,true)
		connection.Write([]byte("+OK\r\n"))
	}

}