package domain

import (
	"testing"
)

// test that PrependHttps works
func TestPrependHttps(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		in  string
		out string
	}{
		{"https://example.com", "https://example.com"},
		{"http://example.com", "https://example.com"},
		{"example.com", "https://example.com"},
		{"//example.com", "https://example.com"},
	}

	for _, tc := range testCases {
		if PrependHttps(tc.in) != tc.out {
			t.Errorf("PrependHttps(%q) = %q, want %q", tc.in, PrependHttps(tc.in), tc.out)
		}
	}
}

func TestRemoveProtocol(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		in  string
		out string
	}{
		{"http://example.com", "example.com"},
		{"http://example.com/", "example.com"},
		{"https://example.com", "example.com"},
		{"https://example.com/", "example.com"},
		{"example.com", "example.com"},
		{"example.com/", "example.com"},
		{"//example.com", "example.com"},
	}

	for _, tc := range testCases {
		if RemoveProtocol(tc.in) != tc.out {
			t.Errorf("RemoveProtocol(%q) = %q, want %q", tc.in, RemoveProtocol(tc.in), tc.out)
		}
	}
}

func TestStripProtocol(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		in  string
		out string
	}{
		{"http://example.com", "example.com"},
		{"http://example.com/", "example.com/"},
		{"https://example.com", "example.com"},
		{"https://example.com/", "example.com/"},
		{"example.com", "example.com"},
		{"example.com/", "example.com/"},
		{"//example.com", "//example.com"},
	}

	for _, tc := range testCases {
		if StripProtocol(tc.in) != tc.out {
			t.Errorf("StripProtocol(%q) = %q, want %q", tc.in, StripProtocol(tc.in), tc.out)
		}
	}
}

// test that ForceProtocol works
func TestForceProtocol(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		in       string
		out      string
		protocol string
	}{
		{"http://example.com", "http://example.com", "http"},
		{"https://example.com", "http://example.com", "http"},
		{"//example.com", "http://example.com", "http"},
		{"example.com", "http://example.com", "http"},
		{"https://example.com", "https://example.com", "https"},
		{"http://example.com", "https://example.com", "https"},
		{"//example.com", "https://example.com", "https"},
		{"example.com", "https://example.com", "https"},
	}

	for _, tc := range testCases {
		if ForceProtocol(tc.in, tc.protocol) != tc.out {
			t.Errorf("ForceProtocol(%q) = %q, want %q", tc.in, ForceProtocol(tc.in, tc.protocol), tc.out)
		}
	}
}
func TestHttpsToHttp(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		in  string
		out string
	}{
		{"https://example.com", "http://example.com"},
		{"http://example.com", "http://example.com"},
		{"example.com", "example.com"},
	}

	for _, tc := range testCases {
		if HttpsToHttp(tc.in) != tc.out {
			t.Errorf("HttpsToHttp(%q) = %q, want %q", tc.in, HttpsToHttp(tc.in), tc.out)
		}
	}
}
func TestHttpToHttps(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		in  string
		out string
	}{
		{"http://example.com", "https://example.com"},
		{"https://example.com", "https://example.com"},
		{"example.com", "example.com"},
	}

	for _, tc := range testCases {
		if HttpToHttps(tc.in) != tc.out {
			t.Errorf("HttpToHttps(%q) = %q, want %q", tc.in, HttpToHttps(tc.in), tc.out)
		}
	}
}

func TestPrependHttp(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		in  string
		out string
	}{
		{"http://example.com", "http://example.com"},
		{"https://example.com", "http://example.com"},
		{"example.com", "http://example.com"},
	}

	for _, tc := range testCases {
		if PrependHttp(tc.in) != tc.out {
			t.Errorf("PrependHttp(%q) = %q, want %q", tc.in, PrependHttp(tc.in), tc.out)
		}
	}
}

// NOTE: this function uses the builtin
// function url.ParseRequestURI(uri)
// so this is really testing a builtin function
// adding it here for completeness/coverage
func TestIsValidURI(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		in  string
		out bool
	}{
		{"http://example.com", true},
		{"https://example.com", true},
		{"example.com", false},
		{"//example.com", true},
		{"ftp://example.com", true},
		{"", false},
	}

	for _, tc := range testCases {
		if IsValidURI(tc.in) != tc.out {
			t.Errorf("IsValidURI(%q) = %v, want %v", tc.in, IsValidURI(tc.in), tc.out)
		}
	}
}
