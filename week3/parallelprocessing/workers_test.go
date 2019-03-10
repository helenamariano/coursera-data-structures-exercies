package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	type args struct {
		numWorkers int
		jobs       []int
	}
	tests := []struct {
		name       string
		args       args
		wantResult [][2]int
	}{
		{"", args{numWorkers: 2, jobs: []int{1, 2, 3, 4, 5}}, [][2]int{{0, 0}, {1, 0}, {0, 1}, {1, 2}, {0, 4}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult := Solve(tt.args.numWorkers, tt.args.jobs)
			assert.Equal(t, tt.wantResult, gotResult)
		})
	}
}
