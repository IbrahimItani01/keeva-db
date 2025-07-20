package utils

import "bufio"

func ConsumeUnecessaryBytes (numberOfBytes int,reader *bufio.Reader) {
	for i := 0; i < numberOfBytes; i++ {
		reader.ReadByte()
	}
}