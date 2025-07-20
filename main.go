package main

import (
	"fmt"
	"keeva-db/services"
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
		resp := services.NewResp(connection)
		value, err := resp.Read()
		utils.ParseError(err,false)
		fmt.Println(value)
		connection.Write([]byte("+OK\r\n"))
	}

}