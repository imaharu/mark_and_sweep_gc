package main

import (
	"fmt"
	"strconv"
)

type Object struct {
	marked          bool
	object_type     ObjectType
	ptr             []string // TODO: ObjectTypeがIntの時は、intにしたい
	size            int
	next_free_index int
}

type ObjectType string

var roots []int

var free_list int

var heap [HEAP_SIZE]Object

func newObjectType(object_type string) ObjectType {
	if object_type == "Array" || object_type == "Int" {
		return ObjectType(object_type)
	} else {
		panic("error type")
	}
}

func mark_phase() {
	for i := range roots {
		var heap_index = roots[i]
		mark(&heap[heap_index])
	}
}

func mark(o *Object) {
	o.marked = true

	if o.object_type == "Array" {
		for i := range o.ptr {
			index, _ := strconv.Atoi(o.ptr[i])
			mark(&heap[index])
		}
	}
}

func sweep_phase() {
	free_list = -1
	for i := range heap {
		if heap[i].marked == true {
			heap[i].marked = false
		} else {
			free_obj(&heap[i])
			heap[i].next_free_index = free_list
			free_list = i
		}
	}
}

func free_obj(o *Object) {
	o.ptr = []string{""}
}

const (
	HEAP_SIZE = 10
)

func init_global_vars() {
	h := Object{marked: false, object_type: "Null", ptr: []string{""}, size: 0}
	for i := range heap {
		heap[i] = h
	}

	var array_type = newObjectType("Array")
	var int_type = newObjectType("Int")

	// TODO: あとで絵を描く

	// rootsから辿れる
	heap[0] = Object{marked: false, object_type: array_type, ptr: []string{"5", "6", "7"}, size: 3}
	heap[5] = Object{marked: false, object_type: int_type, ptr: []string{"55555"}, size: 5}
	heap[6] = Object{marked: false, object_type: int_type, ptr: []string{"66666"}, size: 5}
	heap[7] = Object{marked: false, object_type: int_type, ptr: []string{"77777"}, size: 5}

	heap[8] = Object{marked: false, object_type: int_type, ptr: []string{"88888"}, size: 5}

	heap[4] = Object{marked: false, object_type: int_type, ptr: []string{"44444"}, size: 5}

	// rootsから辿れない
	heap[2] = Object{marked: false, object_type: array_type, ptr: []string{"3", "9"}, size: 2}
	heap[3] = Object{marked: false, object_type: int_type, ptr: []string{"33333"}, size: 5}
	heap[9] = Object{marked: false, object_type: int_type, ptr: []string{"99999"}, size: 5}

	heap[1] = Object{marked: false, object_type: int_type, ptr: []string{"11111"}, size: 5}

	roots = []int{0, 4, 8}
}

func print_global_vars() {
	fmt.Println("### heap ###")
	for i := range heap {
		fmt.Printf("--- heap %d ---\n", i)
		fmt.Println(heap[i])
	}
}

func main() {
	init_global_vars()
	print_global_vars()

	mark_phase()
	fmt.Println("### mark phase done ###")
	print_global_vars()

	sweep_phase()
	fmt.Println("### sweep phase done ###")
	print_global_vars()
}
