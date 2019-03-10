package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_buildHeap(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
		want [][2]int
	}{
		{"5 4 3 2 1", args{data: []int{5, 4, 3, 2, 1}}, [][2]int{{1, 4}, {0, 1}, {1, 3}}},
		{"1 2 3 4 5", args{data: []int{1, 2, 3, 4, 5}}, [][2]int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := buildHeap(tt.args.data)
			assert.Equal(t, tt.want, got)
		})
	}
}
