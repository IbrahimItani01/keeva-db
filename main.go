package main

import (
	"fmt"
	"keeva-db/util"
	"net"
)

func main () {
	port := ":6379"
	fmt.Println("Listening on port",port)
	listener, err:= net.Listen("tcp",port)
	util.ParseError(err,false)
	connection, err := listener.Accept()
	util.ParseError(err,false)
	defer connection.Close()

	for {
		buffer := make([]byte,1024)
		_,err := connection.Read(buffer)
		util.ParseError(err,true)
		connection.Write([]byte("+OK\r\n"))
	}

}