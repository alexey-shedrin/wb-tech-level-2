package main

import (
	"testing"
)

func TestUnpackString(t *testing.T) {
	type testCase struct {
		name        string
		input       string
		expected    string
		expectError bool
	}

	testCases := []testCase{
		{
			name:        "Standard case",
			input:       "a4bc2d5e",
			expected:    "aaaabccddddde",
			expectError: false,
		},
		{
			name:        "No numbers",
			input:       "abcd",
			expected:    "abcd",
			expectError: false,
		},
		{
			name:        "Invalid string (only numbers)",
			input:       "45",
			expected:    "",
			expectError: true,
		},
		{
			name:        "Invalid string (starts with a number)",
			input:       "3abc",
			expected:    "",
			expectError: true,
		},
		{
			name:        "Empty string",
			input:       "",
			expected:    "",
			expectError: false,
		},
		{
			name:        "Escaped numbers",
			input:       "qwe\\4\\5",
			expected:    "qwe45",
			expectError: false,
		},
		{
			name:        "Escaped number with repetition",
			input:       "qwe\\45",
			expected:    "qwe44444",
			expectError: false,
		},
		{
			name:        "Escaped backslash",
			input:       "qwe\\\\5",
			expected:    "qwe\\\\\\\\\\",
			expectError: false,
		},
		{
			name:        "Single character",
			input:       "a",
			expected:    "a",
			expectError: false,
		},
		{
			name:        "Ending with a number",
			input:       "abc3",
			expected:    "abccc",
			expectError: false,
		},
		{
			name:        "Ending with an escaped backslash",
			input:       "abc\\",
			expected:    "",
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := UnpackString(tc.input)

			if tc.expectError {
				if err == nil {
					t.Errorf("expected an error, but got none for input: %q", tc.input)
				}
				return
			}

			if err != nil {
				t.Errorf("did not expect an error, but got: %v for input: %q", err, tc.input)
			}

			if result != tc.expected {
				t.Errorf("for input %q, expected %q, but got %q", tc.input, tc.expected, result)
			}
		})
	}
}
