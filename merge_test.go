package main

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		name        string
		collection1 []int
		collection2 []int
		collection3 []int
		want        []int
	}{
		{
			name:        "all three collections have values",
			collection1: []int{9, 5, 1},
			collection2: []int{0, 2, 6},
			collection3: []int{3, 4, 8},
			want:        []int{0, 1, 2, 3, 4, 5, 6, 8, 9},
		},
		{
			name:        "collection1 descending mixed with two ascending",
			collection1: []int{9, 7, 4, 1},
			collection2: []int{0, 2, 5, 8},
			collection3: []int{0, 3, 6, 10},
			want:        []int{0, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name:        "all collections empty",
			collection1: []int{},
			collection2: []int{},
			collection3: []int{},
			want:        []int{},
		},
		{
			name:        "nil collections treated as empty",
			collection1: nil,
			collection2: nil,
			collection3: nil,
			want:        []int{},
		},
		{
			name:        "only collection1 has values",
			collection1: []int{5, 3, 1},
			collection2: []int{},
			collection3: []int{},
			want:        []int{1, 3, 5},
		},
		{
			name:        "only collection2 has values",
			collection1: []int{},
			collection2: []int{0, 2, 4},
			collection3: []int{},
			want:        []int{0, 2, 4},
		},
		{
			name:        "only collection3 has values",
			collection1: []int{},
			collection2: []int{},
			collection3: []int{1, 3, 7},
			want:        []int{1, 3, 7},
		},
		{
			name:        "one collection empty",
			collection1: []int{8, 2},
			collection2: []int{},
			collection3: []int{1, 4, 9},
			want:        []int{1, 2, 4, 8, 9},
		},
		{
			name:        "duplicate values across collections",
			collection1: []int{4, 4, 2},
			collection2: []int{2, 3, 4},
			collection3: []int{1, 4},
			want:        []int{1, 2, 2, 3, 4, 4, 4, 4},
		},
		{
			name:        "single element in each collection",
			collection1: []int{5},
			collection2: []int{1},
			collection3: []int{3},
			want:        []int{1, 3, 5},
		},
		{
			name:        "already overlapping ranges",
			collection1: []int{10, 6, 2},
			collection2: []int{0, 6, 11},
			collection3: []int{2, 7},
			want:        []int{0, 2, 2, 6, 6, 7, 10, 11},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Merge(tt.collection1, tt.collection2, tt.collection3)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
