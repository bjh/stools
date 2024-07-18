package cnvrt

import "testing"

func TestToI(t *testing.T) {
	testCases := []struct {
		in  string
		out int
	}{
		{"123", 123},
		{"-456", -456},
		{"0", 0},
		{"999999999", 999999999},
	}

	for _, tc := range testCases {
		result := ToI(tc.in)
		if result != tc.out {
			t.Errorf("ToI(%q) = %d, want %d", tc.in, result, tc.out)
		}
	}
}
func TestToF(t *testing.T) {
	testCases := []struct {
		in  string
		out float32
	}{
		{"3.14", 3.14},
		{"-2.5", -2.5},
		{"0", 0},
		{"123.456", 123.456},
	}

	for _, tc := range testCases {
		result := ToF(tc.in)
		if result != tc.out {
			t.Errorf("ToF(%q) = %f, want %f", tc.in, result, tc.out)
		}
	}
}
