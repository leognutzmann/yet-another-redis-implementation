package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

type RespReader struct {
	reader *bufio.Reader
}

func NewRespReader(rd io.Reader) *RespReader {
	return &RespReader{reader: bufio.NewReader(rd)}
}

func (resp *RespReader) Read() (Value, error) {
	dataType, err := resp.reader.ReadByte()

	if err != nil {
		return Value{}, err
	}

	switch dataType {
	case ARRAY:
		return resp.readArray()
	case BULK:
		return resp.readBulk()
	default:
		fmt.Printf("Unknown type: %v", string(dataType))
		return Value{}, nil
	}
}

func (resp *RespReader) readArray() (Value, error) {
	value := Value{}
	value.dataType = "array"

	length, _, err := resp.readInteger()
	if err != nil {
		return value, err
	}

	value.array = make([]Value, 0)
	for i := 0; i < length; i++ {
		val, err := resp.Read()
		if err != nil {
			return value, err
		}
		value.array = append(value.array, val)
	}

	return value, nil
}

func (resp *RespReader) readBulk() (Value, error) {
	value := Value{}
	value.dataType = "bulk"

	len, _, err := resp.readInteger()
	if err != nil {
		return value, err
	}

	bulk := make([]byte, len)
	resp.reader.Read(bulk)
	value.bulk = string(bulk)

	resp.readLine()

	return value, nil
}

func (resp *RespReader) readLine() (line []byte, numberOfBytesRead int, err error) {
	for {
		b, err := resp.reader.ReadByte()
		if err != nil {
			return nil, 0, err
		}
		numberOfBytesRead += 1
		line = append(line, b)
		if len(line) >= 2 && line[len(line)-2] == '\r' && line[len(line)-1] == '\n' {
			break
		}
	}
	return line[:len(line)-2], numberOfBytesRead, nil
}

func (resp *RespReader) readInteger() (value int, numberOfBytesRead int, err error) {
	line, numberOfBytesRead, err := resp.readLine()
	if err != nil {
		return 0, 0, err
	}
	i64, err := strconv.ParseInt(string(line), 10, 64)
	if err != nil {
		return 0, numberOfBytesRead, err
	}
	return int(i64), numberOfBytesRead, nil
}
