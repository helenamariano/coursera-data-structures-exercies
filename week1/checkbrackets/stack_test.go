package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack_Push(t *testing.T) {
	tests := []struct {
		name  string
		items []int
	}{
		{"Push 3 items to stack", []int{3, 2, 1}},
		{"Push 5 items to stack", []int{3, 5, 8, 9, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := StackNew()

			for _, item := range tt.items {
				s.Push(item)
			}
			assert.ElementsMatch(t, tt.items, s.stack)
		})
	}
}

func TestStack_Pop(t *testing.T) {
	tests := []struct {
		name  string
		items []int
	}{
		{"Pop 3 items to stack", []int{3, 2, 1}},
		{"Pop 5 items to stack", []int{3, 5, 8, 9, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := StackNew()

			for _, item := range tt.items {
				s.Push(item)
			}

			numItems := len(tt.items)
			for i := range tt.items {
				val, err := s.Pop()
				assert.NoError(t, err)
				assert.Equal(t, tt.items[numItems-1-i], val)
			}

		})
	}
}

func TestNew(t *testing.T) {
	t.Run("New Stack", func(t *testing.T) {
		s := StackNew()
		assert.NotNil(t, s)
	})
}

func TestStack_TopStackEmpty(t *testing.T) {
	t.Run("Top with Stack Empty", func(t *testing.T) {
		s := StackNew()
		_, err := s.Top()
		assert.EqualError(t, err, "Stack is empty")
	})
}

func TestStack_PopStackEmpty(t *testing.T) {
	t.Run("Pop with Stack Empty", func(t *testing.T) {
		s := StackNew()
		_, err := s.Pop()
		assert.EqualError(t, err, "Stack is empty")
	})
}
