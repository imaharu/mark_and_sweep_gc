package main

import (
	"fmt"
	"reflect"
)

type Heap struct {
	marked      bool
	object_type ObjectType
	ptr         []string
	size        int
}

type ObjectType string

func newObjectType(object_type string) ObjectType {
	if object_type == "Array" || object_type == "Int" {
		return ObjectType("object_type")
	} else {
		panic("error type")
	}
}

func (h Heap) heap_size() int {
	return reflect.ValueOf(h).NumField()
}

func mark_phase() {}

func sweep_phase() {}

func main() {
	var h = Heap{marked: true}
	fmt.Println(h.marked)
	fmt.Println(h.heap_size())
}
