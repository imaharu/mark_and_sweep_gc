package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type Heap struct {
	marked      bool
	object_type ObjectType
	ptr         []string // TODO: ObjectTypeがIntの時は、intにしたい
	size        int
}

type ObjectType string

var roots []int

var heaps [HEAP_SIZE]Heap

func newObjectType(object_type string) ObjectType {
	if object_type == "Array" || object_type == "Int" {
		return ObjectType(object_type)
	} else {
		panic("error type")
	}
}

func (h Heap) heap_size() int {
	return reflect.ValueOf(h).NumField()
}

func mark_phase() {
	for i := range roots {
		var heap_index = roots[i]
		mark(&heaps[heap_index])
	}
}

func mark(h *Heap) {
	h.marked = true

	if h.object_type == "Array" {
		for i := range h.ptr {
			index, _ := strconv.Atoi(h.ptr[i])
			mark(&heaps[index])
		}
	}
}

func sweep_phase() {}

const (
	HEAP_SIZE = 11
)

func init_global_vars() {
	h := Heap{marked: false, object_type: "Null", ptr: []string{""}, size: 0}
	for i := range heaps {
		heaps[i] = h
	}

	var array_type = newObjectType("Array")
	var int_type = newObjectType("Int")

	// TODO: あとで絵を描く

	// rootsから辿れる
	heaps[0] = Heap{marked: false, object_type: array_type, ptr: []string{"5", "6", "7"}, size: 3}
	heaps[5] = Heap{marked: false, object_type: int_type, ptr: []string{"55555"}, size: 5}
	heaps[6] = Heap{marked: false, object_type: int_type, ptr: []string{"66666"}, size: 5}
	heaps[7] = Heap{marked: false, object_type: int_type, ptr: []string{"77777"}, size: 5}

	heaps[8] = Heap{marked: false, object_type: int_type, ptr: []string{"88888"}, size: 5}

	heaps[4] = Heap{marked: false, object_type: int_type, ptr: []string{"44444"}, size: 5}

	// rootsから辿れない
	heaps[2] = Heap{marked: false, object_type: array_type, ptr: []string{"3", "9"}, size: 2}
	heaps[3] = Heap{marked: false, object_type: int_type, ptr: []string{"33333"}, size: 5}
	heaps[9] = Heap{marked: false, object_type: int_type, ptr: []string{"99999"}, size: 5}

	heaps[1] = Heap{marked: false, object_type: int_type, ptr: []string{"11111"}, size: 5}

	roots = []int{0, 4, 8, 10}
}

func print_global_vars() {
	fmt.Println("### heaps ###")
	for i := range heaps {
		fmt.Printf("--- heap %d ---\n", i)
		fmt.Println(heaps[i])
	}
}

func main() {
	init_global_vars()
	print_global_vars()
	mark_phase()

	fmt.Println("### mark phase done ###")
	print_global_vars()
}
