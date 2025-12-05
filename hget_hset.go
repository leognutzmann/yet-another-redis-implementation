package main

import "sync"

var HSets = map[string]map[string]string{}
var HSetsMutex = sync.RWMutex{}

func hset(args []Value) Value {
	if len(args) != 3 {
		return Value{dataType: "error", str: "ERR wrong number of arguments for 'hset' command"}
	}

	hash := args[0].bulk
	key := args[1].bulk
	value := args[2].bulk

	HSetsMutex.Lock()
	if _, ok := HSets[hash]; !ok {
		HSets[hash] = map[string]string{}
	}
	HSets[hash][key] = value
	HSetsMutex.Unlock()

	return Value{dataType: "string", str: "OK"}
}

func hget(args []Value) Value {
	if len(args) != 2 {
		return Value{dataType: "error", str: "ERR wrong number of arguments for 'hget' command"}
	}

	hash := args[0].bulk
	key := args[1].bulk

	HSetsMutex.RLock()
	value, ok := HSets[hash][key]
	HSetsMutex.RUnlock()

	if !ok {
		return Value{dataType: "null"}
	}

	return Value{dataType: "bulk", bulk: value}
}

func hgetall(args []Value) Value {
	if len(args) != 1 {
		return Value{dataType: "error", str: "ERR wrong number of arguments for 'hgetall' command"}
	}

	hash := args[0].bulk

	HSetsMutex.RLock()
	value, ok := HSets[hash]
	HSetsMutex.RUnlock()

	if !ok {
		return Value{dataType: "null"}
	}

	values := []Value{}
	for k, v := range value {
		values = append(values, Value{dataType: "bulk", bulk: k})
		values = append(values, Value{dataType: "bulk", bulk: v})
	}

	return Value{dataType: "array", array: values}
}
