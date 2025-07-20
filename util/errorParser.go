package util

import (
	"fmt"
	"io"
	"os"
)

func ParseError (err error, fromClient bool){
	if err != nil {
		if err == io.EOF && fromClient {
			return 
		}
		if fromClient{
			fmt.Println("Error reading from client",err.Error())
		}else {
			fmt.Println(err)
		}
		os.Exit(1)
	}
}