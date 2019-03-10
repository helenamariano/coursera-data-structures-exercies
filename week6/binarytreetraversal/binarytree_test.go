package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBT_InOrder(t *testing.T) {
	type fields struct {
		keys  []int
		left  []int
		right []int
	}
	tests := []struct {
		name   string
		fields fields
		want   []int
	}{
		{"5 nodes", fields{keys: []int{4, 2, 5, 1, 3}, left: []int{1, 3, -1, -1, -1}, right: []int{2, 4, -1, -1, -1}}, []int{1, 2, 3, 4, 5}},
		{"10 nodes", fields{keys: []int{0, 10, 20, 30, 40, 50, 60, 70, 80, 90}, left: []int{7, -1, -1, 8, 3, -1, 1, 5, -1, -1}, right: []int{2, -1, 6, 9, -1, -1, -1, 4, -1, -1}}, []int{50, 70, 80, 30, 90, 40, 0, 20, 10, 60}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bt := NewBinaryTree(tt.fields.keys, tt.fields.left, tt.fields.right)
			result := bt.InOrder()
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestBT_PreOrder(t *testing.T) {
	type fields struct {
		keys  []int
		left  []int
		right []int
	}
	tests := []struct {
		name   string
		fields fields
		want   []int
	}{
		{"5 nodes", fields{keys: []int{4, 2, 5, 1, 3}, left: []int{1, 3, -1, -1, -1}, right: []int{2, 4, -1, -1, -1}}, []int{4, 2, 1, 3, 5}},
		{"10 nodes", fields{keys: []int{0, 10, 20, 30, 40, 50, 60, 70, 80, 90}, left: []int{7, -1, -1, 8, 3, -1, 1, 5, -1, -1}, right: []int{2, -1, 6, 9, -1, -1, -1, 4, -1, -1}}, []int{0, 70, 50, 40, 30, 80, 90, 20, 60, 10}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bt := NewBinaryTree(tt.fields.keys, tt.fields.left, tt.fields.right)
			result := bt.PreOrder()
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestBT_PostOrder(t *testing.T) {
	type fields struct {
		keys  []int
		left  []int
		right []int
	}
	tests := []struct {
		name   string
		fields fields
		want   []int
	}{
		{"5 nodes", fields{keys: []int{4, 2, 5, 1, 3}, left: []int{1, 3, -1, -1, -1}, right: []int{2, 4, -1, -1, -1}}, []int{1, 3, 2, 5, 4}},
		{"10 nodes", fields{keys: []int{0, 10, 20, 30, 40, 50, 60, 70, 80, 90}, left: []int{7, -1, -1, 8, 3, -1, 1, 5, -1, -1}, right: []int{2, -1, 6, 9, -1, -1, -1, 4, -1, -1}}, []int{50, 80, 90, 30, 40, 70, 10, 60, 20, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bt := NewBinaryTree(tt.fields.keys, tt.fields.left, tt.fields.right)
			result := bt.PostOrder()
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestBT_IsBst(t *testing.T) {
	type fields struct {
		keys  []int
		left  []int
		right []int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"empty tree", fields{keys: []int{}, left: []int{}, right: []int{}}, true},
		{"3 nodes", fields{keys: []int{2, 1, 3}, left: []int{1, -1, -1}, right: []int{2, -1, -1}}, true},
		{"3 nodes", fields{keys: []int{1, 2, 3}, left: []int{1, -1, -1}, right: []int{2, -1, -1}}, false},
		{"5 nodes", fields{keys: []int{1, 2, 3, 4, 5}, left: []int{-1, -1, -1, -1, -1}, right: []int{1, 2, 3, 4, -1}}, true},
		{"7 nodes", fields{keys: []int{4, 2, 6, 1, 3, 5, 7}, left: []int{1, 3, 5, -1, -1, -1, -1}, right: []int{2, 4, 6, -1, -1, -1, -1}}, true},
		{"4 nodes", fields{keys: []int{4, 2, 1, 5}, left: []int{1, 2, -1, -1}, right: []int{-1, 3, -1, -1}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bt := NewBinaryTree(tt.fields.keys, tt.fields.left, tt.fields.right)
			result := bt.IsBst()
			assert.Equal(t, tt.want, result)
		})
	}
}
