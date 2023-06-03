package main

import "testing"

func TestSum(t *testing.T) {
	tests := []struct {
		a    int
		b    int
		want int
	}{
		{1, 1, 2},
		{10, 5, 15},
	}

	for _, tt := range tests {
		tt := tt
		t.Run("test", func(t *testing.T) {
			t.Parallel()
			got := Sum(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Sum(%d,%d) == %d, wabt %d\n", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
