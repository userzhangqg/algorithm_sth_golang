package main

import (
	"reflect"
)

func GetValue() int8 {
	return 1
}

func main() {
	i := GetValue()

	switch reflect.TypeOf(i).Name() {
	case "int":
		println("int")
	default:
		println(reflect.TypeOf(i).Name())
	}

}
