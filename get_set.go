package main

import "sync"

var Sets = map[string]string{}
var SetsMutex = sync.RWMutex{}

func set(args []Value) Value {
	if len(args) != 2 {
		return Value{dataType: "error", str: "ERR wrong number of arguments for 'set' command"}
	}

	key := args[0].bulk
	value := args[1].bulk

	SetsMutex.Lock()
	Sets[key] = value
	SetsMutex.Unlock()

	return Value{dataType: "string", str: "OK"}
}

func get(args []Value) Value {
	if len(args) != 1 {
		return Value{dataType: "error", str: "ERR wrong number of arguments for 'get' command"}
	}

	key := args[0].bulk

	SetsMutex.RLock()
	value, ok := Sets[key]
	SetsMutex.RUnlock()

	if !ok {
		return Value{dataType: "null"}
	}

	return Value{dataType: "bulk", bulk: value}
}
