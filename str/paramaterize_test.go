package str

import (
	"testing"
)

func TestParameterize(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Hello World", "hello-world"},
		{"Hello, World!", "hello-world"},
		{"This is a test", "this-is-a-test"},
		{"12345", "12345"},
		{"!@#$%^&*()", ""},
		{"", ""},
	}

	for _, test := range tests {
		result := Parameterize(test.input)
		if result != test.expected {
			t.Errorf("Parameterize failed for input '%s'. Expected '%s', but got '%s'", test.input, test.expected, result)
		}
	}
}
