package main

import "errors"

type Stack struct {
	stack []interface{}
}

func (this *Stack) Push(value interface{}) {
	this.stack = append(this.stack, value)
}

func (this *Stack) Pop() (interface{}, error) {
	result, err := this.Top()
	if err != nil {
		return 0, err
	}

	this.stack[len(this.stack)-1] = nil
	this.stack = this.stack[:len(this.stack)-1]
	return result, nil
}

func (this *Stack) Top() (interface{}, error) {
	if len(this.stack) == 0 {
		return 0, errors.New("Stack is empty")
	}
	result := this.stack[len(this.stack)-1]
	return result, nil
}

func (this *Stack) Empty() bool {
	return len(this.stack) == 0
}

func StackNew() *Stack {
	return &Stack{}
}
