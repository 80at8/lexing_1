package main

import (
	"fmt"
	"strconv"
)

type node struct {
	value  int
	opType lexeme
	child  []node
}

func (n *node) addChildren(children ...node) {
	for _, j := range children {
		fmt.Printf("-- j is opType %v, opVal %v\n", j.opType, j.value)
		n.child = append(n.child, j)
	}
}

func (n *node) getChildren() ([]node, error) {
	if len(n.child) > 0 {
		return n.child, nil
	}
	return nil, nil
}

type ast struct {
	root           node
	numberOfLevels int
	currentLevel   int
}

//var stack ast

func parse() {

	//Now For constructing expression tree we use a stack. We loop through input expression and do following for every character.
	// 1) If character is operand push that into stack
	// 2) If character is operator pop two values from stack make them its child and push current node again.
	// At the end only element of stack will be root of expression tree

	var stack = make(map[int]node)
	var a, b node

	for n := 0; n < len(tokenListType.tokenList); n++ {

		if tokenListType.tokenTypeList[n] == LITERAL {
			a.value, _ = strconv.Atoi(tokenListType.tokenList[n])
			stack[len(stack)] = a
		}

		if tokenListType.tokenTypeList[n] == OPERATOR {

			if len(stack) > 0 && n+1 < len(tokenListType.tokenList) {
				//pop single
				a.value, _ = strconv.Atoi(tokenListType.tokenList[n])
				stack[len(stack)-1] = a
			} else {
				// pop double
				b.addChildren(stack[n-2], stack[n-1])
				stack[len(stack)] = b

			}
		}

	}

	fmt.Println("Stack Debug")
	for n := 0; n < len(stack); n++ {
		fmt.Printf("n[%d] is %v, value is %d\n", n, stack[n].opType, stack[n].value)
	}
}
