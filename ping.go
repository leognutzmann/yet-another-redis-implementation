package main

func ping(args []Value) Value {
	if len(args) == 0 {
		return Value{dataType: "string", str: "PONG"}
	}

	return Value{dataType: "string", str: args[0].bulk}
}
