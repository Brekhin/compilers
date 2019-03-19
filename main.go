package main

import (
	"container/list"
	"fmt"

	"./utils"
)

func main() {
	opstack := make(utils.Stack, 0)
	var x list.List

	x.PushBack(1)
	x.PushBack(2)
	for e := x.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	opstack = opstack.Push("1")
	opstack = opstack.Push("2")
	opstack = opstack.Push("3")

	opstack, p := opstack.Pop()
	fmt.Println(opstack, p)
	opstack, e := opstack.Pop()

	fmt.Println(opstack, e)
}
