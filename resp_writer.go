package main

import (
	"io"
	"strconv"
)

type Writer struct {
	writer io.Writer
}

func NewRespWriter(writer io.Writer) *Writer {
	return &Writer{writer: writer}
}

func (writer *Writer) Write(value Value) error {
	var bytes = value.Marshal()

	_, err := writer.writer.Write(bytes)
	if err != nil {
		return err
	}

	return nil
}

func (value Value) Marshal() []byte {
	switch value.dataType {
	case "array":
		return value.marshalArray()
	case "bulk":
		return value.marshalBulk()
	case "string":
		return value.marshalString()
	case "null":
		return value.marshallNull()
	case "error":
		return value.marshallError()
	default:
		return []byte{}
	}
}

func (value Value) marshalString() []byte {
	var bytes []byte
	bytes = append(bytes, STRING)
	bytes = append(bytes, value.str...)
	bytes = append(bytes, '\r', '\n')

	return bytes
}

func (value Value) marshalBulk() []byte {
	var bytes []byte
	bytes = append(bytes, BULK)
	bytes = append(bytes, strconv.Itoa(len(value.bulk))...)
	bytes = append(bytes, '\r', '\n')
	bytes = append(bytes, value.bulk...)
	bytes = append(bytes, '\r', '\n')

	return bytes
}

func (value Value) marshalArray() []byte {
	var bytes []byte
	length := len(value.array)
	bytes = append(bytes, ARRAY)
	bytes = append(bytes, strconv.Itoa(length)...)
	bytes = append(bytes, '\r', '\n')

	for i := range length {
		bytes = append(bytes, value.array[i].Marshal()...)
	}

	return bytes
}

func (value Value) marshallError() []byte {
	var bytes []byte
	bytes = append(bytes, ERROR)
	bytes = append(bytes, value.str...)
	bytes = append(bytes, '\r', '\n')

	return bytes
}

func (value Value) marshallNull() []byte {
	return []byte("$-1\r\n")
}
