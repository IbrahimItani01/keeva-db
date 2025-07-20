package services

import (
	"bufio"
	"fmt"
	"io"
	"keeva-db/utils"
	"os"
	"strconv"
	"strings"
)

func RESPReader (input string) {
	reader := bufio.NewReader(strings.NewReader(input))
	byte,_ := reader.ReadByte()
	if byte != '$'{
		fmt.Println("Invalid type, only bulk strings accepted")
		os.Exit(1)
	}
	size,_ := reader.ReadByte()
	strSize,_ := strconv.ParseInt(string(size),10,64)
	utils.ConsumeUnecessaryBytes(2,reader)
	bulkString := make([]byte, strSize)
	reader.Read(bulkString)
	fmt.Println(string(bulkString))
}

func NewResp(rd io.Reader) *utils.Resp {
	return &utils.Resp{reader: bufio.NewReader(rd)}
}