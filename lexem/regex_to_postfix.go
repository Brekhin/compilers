package lexem

import (
	"container/list"

	"github.com/golang-collections/collections/stack"
)

func priority(ch string) int {
	switch ch {
	case "*":
		return 4
	case ".":
		return 3
	case "+":
		return 2
	case "|":
		return 1
	}
	return 0
}

func ConvertRegexToPostfix(infix string) list.List {

	opstack := stack.New()
	var result list.List

	for _, symbol := range infix {
		switch {
		case symbol == '(':
			opstack.Push(string(symbol))
		case symbol == ')':
			for opstack.Peek() != "(" {
				result.PushBack(opstack.Peek().(string))
				opstack.Pop()
			}
			opstack.Pop()
		case priority(string(symbol)) > 0:
			for opstack.Len() > 0 && priority(string(symbol)) <= priority(opstack.Peek().(string)) {
				result.PushBack(opstack.Peek().(string))
				opstack.Pop()
			}
			opstack.Push(string(symbol))
		default:
			result.PushBack(string(symbol))
		}
	}

	for opstack.Len() > 0 {
		result.PushBack(opstack.Peek().(string))
		opstack.Pop()
	}

	return result
}
