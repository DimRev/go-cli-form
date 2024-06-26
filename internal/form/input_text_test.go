package form

import (
	"bytes"
	"os"
	"testing"
)

func TestTextInput(t *testing.T) {
	cases := []struct {
		name          string
		question      string
		input         string
		expected      string
		expectedError string
	}{
		// TODO: Add tests for ANSI input checking colors and themes.
		{
			name:          "normal input",
			question:      "Test question",
			input:         "Test answer\n",
			expected:      "Test answer",
			expectedError: "",
		},
		{
			name:          "empty input",
			question:      "Test question",
			input:         "\n",
			expected:      "",
			expectedError: "",
		},
		{
			name:          "empty question",
			question:      "",
			input:         "",
			expected:      "",
			expectedError: "question cannot be empty",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			// Redirect stdin and stdout
			oldStdin := os.Stdin
			r, w, _ := os.Pipe()
			os.Stdin = r
			go func() {
				w.Write([]byte(tc.input))
				w.Close()
			}()

			oldStdout := os.Stdout
			out, outW, _ := os.Pipe()
			os.Stdout = outW

			actualErrStr := ""
			Form := Start("default")
			actual, err := Form.TextInput(tc.question)
			if err != nil {
				actualErrStr = err.Error()
			}

			// Close and restore stdout
			outW.Close()
			os.Stdout = oldStdout

			var buf bytes.Buffer
			buf.ReadFrom(out)

			// Restore stdin
			os.Stdin = oldStdin

			if actual != tc.expected {
				t.Errorf("\n❌ Wrong Value ❌\n\nActual:\t\t%q\nExpected:\t\t%q\n", actual, tc.expected)
			}
			if actualErrStr != tc.expectedError {
				t.Errorf("\n❌ Wrong Value Error ❌\n\nActual Error:\t\t%q\nExpected Error:\t\t%q\n", actualErrStr, tc.expectedError)
			}
		})
	}
}
