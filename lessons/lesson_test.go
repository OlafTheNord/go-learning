package lessons

import "testing"

func TestIsEven(t *testing.T) {
	tests := []struct {
		input    int
		expected bool
	}{
		{3, false},
		{4, true},
		{5, false},
		{6, true},
	}
	for _, tt := range tests {
		got := IsEven(tt.input)
		if got != tt.expected {
			t.Errorf("IsEvev(%v)=%v; expected: %v", tt.input, got, tt.expected)
		}
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		a, b, exp int
	}{
		{1, 2, 2},
		{1, 3, 3},
		{1, 4, 4},
		{1, 5, 5},
		{1, 6, 6},
		{1, 7, 7},
	}

	for _, tt := range tests {
		got := Max(tt.a, tt.b)
		if got != tt.exp {
			t.Errorf("Max of %v, %v is %v; expexted %v", tt.a, tt.b, got, tt.exp)
		}
	}
}
