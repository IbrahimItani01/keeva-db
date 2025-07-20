package utils

import "bufio"

type Resp struct {
	reader *bufio.Reader
}

type Value struct {
	typ   string
	str   string
	num   int
	bulk  string
	array []Value
}