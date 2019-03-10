package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDisjointSet_Merge(t *testing.T) {
	type fields struct {
		tables []int
	}
	type args struct {
		destination int
		source      int
	}

	tables0 := []int{10, 0, 5, 0, 3, 3}
	args0 := []args{{destination: 6, source: 6}, {destination: 6, source: 5}, {destination: 5, source: 4}, {destination: 4, source: 3}}
	expected0 := []int{10, 10, 10, 11}

	tables1 := []int{1, 1, 1, 1, 1}
	args1 := []args{{destination: 3, source: 5}, {destination: 2, source: 4}, {destination: 1, source: 4}, {destination: 5, source: 4}, {destination: 5, source: 3}}
	expected1 := []int{2, 2, 3, 5, 5}

	tests := []struct {
		name   string
		fields fields
		args   []args
		want   []int
	}{
		{"6 tables, 4 merges", fields{tables: tables0}, args0, expected0},
		{"5 tables, 5 merges", fields{tables: tables1}, args1, expected1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := NewDisjointSet(tt.fields.tables)
			for i, arg := range tt.args {
				// inputs are indexed from 1 but implementation is indexed from 0
				maxTableSize := ds.Merge(arg.destination-1, arg.source-1)
				assert.Equal(t, tt.want[i], maxTableSize)
			}
		})
	}
}
