package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		name   string
		target int
		slice  []int
		want   int
	}{
		{
			name:   "EmptyCase",
			target: 3,
			slice:  []int{},
			want:   -1,
		},
		{
			name:   "MiddleCase",
			target: 3,
			slice:  []int{1, 2, 3, 4, 5, 6},
			want:   2,
		},
		{
			name:   "LeftCase",
			target: 2,
			slice:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			want:   1,
		},
		{
			name:   "RightCase",
			target: 10,
			slice:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			want:   9,
		},
		{
			name:   "OutOfSliceCase",
			target: 20,
			slice:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			want:   -1,
		},
		{
			name:   "MoreNumbersCase",
			target: 805,
			slice:  []int{10, 20, 30, 35, 50, 56, 75, 90, 100, 110, 200, 201, 305, 708, 805},
			want:   14,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := binarySearch(tc.slice, tc.target)

			assert.Equal(t, tc.want, result)
		})
	}
}
