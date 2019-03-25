package main

import (
	"fmt"

	"./lexem"
	"./nfa"
)

func main() {
	list := lexem.ConvertRegexToPostfix("a(bb)+a")
	for e := list.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value)
	}
	pp := nfa.Post2nfa(list)
	fmt.Println()
}
