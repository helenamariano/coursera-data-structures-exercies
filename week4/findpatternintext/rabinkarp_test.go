package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindPattern(t *testing.T) {
	type args struct {
		pattern string
		text    string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"pattern=aba, text=abacaba", args{pattern: "aba", text: "abacaba"}, []int{0, 4}},
		{"pattern=Test, text=testTesttesT", args{pattern: "Test", text: "testTesttesT"}, []int{4}},
		{"pattern=aaaaa, text=baaaaaaa", args{pattern: "aaaaa", text: "baaaaaaa"}, []int{1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindPattern(tt.args.pattern, tt.args.text)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_polyHash(t *testing.T) {
	type args struct {
		s string
		x int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"Test", args{s: "Test", x: 908732146}, 196834490},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := polyHash(tt.args.s, tt.args.x)
			assert.Equal(t, tt.want, got)
		})
	}
}
