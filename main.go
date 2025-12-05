package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", ":6379")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Listening on port :6379")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("Client connected:", conn.RemoteAddr())

	for {
		reader := NewRespReader(conn)
		writer := NewRespWriter(conn)
		value, err := reader.Read()
		if err != nil {
			if err.Error() != "EOF" {
				fmt.Println("Error reading:", err)
			}
			fmt.Println("Client disconnected:", conn.RemoteAddr())
			return
		}

		if value.dataType != "array" {
			fmt.Println("Invalid request, expected array")
			continue
		}

		if len(value.array) == 0 {
			fmt.Println("Invalid request, expected array length > 0")
			continue
		}

		command := strings.ToUpper(value.array[0].bulk)
		args := value.array[1:]

		handler, ok := Handlers[command]
		if !ok {
			fmt.Println("Invalid command: ", command)
			writer.Write(Value{dataType: "string", str: ""})
			continue
		}

		result := handler(args)
		writer.Write(result)
	}
}
