package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main () {
	port := ":6379"
	fmt.Println("Listening on port",port)
	listener, err:= net.Listen("tcp",port)
	if err != nil {
		fmt.Println(err)
		return
	}
	connection, err := listener.Accept()
	if err!=nil{
		fmt.Println(err)
		return
	}
	defer connection.Close()

	for {
		buffer := make([]byte,1024)
		_,err := connection.Read(buffer)
		if err !=nil{
			if err == io.EOF {
				break
			}
			fmt.Println("Error reading from client",err.Error())
			os.Exit(1)
		}
		connection.Write([]byte("+OK\r\n"))
	}

}