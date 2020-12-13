package binarySearch

import "testing"

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		name string
		args int
		want bool
	}{
		{name: "", args: 5, want: true},
		{name: "", args: 100, want: false},
		{name: "", args: 0, want: false},
		{name: "", args: 8, want: true},
		{name: "", args: 34, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearch(tt.args); got != tt.want {
				t.Errorf("BinarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
