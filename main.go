package main

import (
	"fmt"

	"./lexem"
)

func main() {
	list := lexem.ConvertRegexToPostfix("(a.(b|d))*")
	for e := list.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value)
	}
	fmt.Println()
}
