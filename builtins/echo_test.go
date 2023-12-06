package builtins_test

import (
	"bytes"
	"testing"

	"github.com/SamFisher0208/Shell_Built_Ins/builtins"
)

func TestEcho(t *testing.T) {
	type args struct {
		writer    *bytes.Buffer
		arguments []string
	}
	tests := []struct {
		name           string
		args           args
		expectedOutput string
	}{
		{
			name: "echo 'Hello, World!'",
			args: args{
				writer:    &bytes.Buffer{},
				arguments: []string{"Hello,", "World!"},
			},
			expectedOutput: "Hello, World!\n",
		},
		{
			name: "echo empty arguments",
			args: args{
				writer:    &bytes.Buffer{},
				arguments: []string{},
			},
			expectedOutput: "\n",
		},
		{
			name: "echo special characters",
			args: args{
				writer:    &bytes.Buffer{},
				arguments: []string{"Hello,", "Wörld!"},
			},
			expectedOutput: "Hello, Wörld!\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Execute the Echo function
			err := builtins.Echo(tt.args.writer, tt.args.arguments...)
			if err != nil {
				t.Errorf("Echo failed with error: %v", err)
			}

			// Check the output
			actualOutput := tt.args.writer.String()
			if actualOutput != tt.expectedOutput {
				t.Errorf("Expected output: %s, but got: %s", tt.expectedOutput, actualOutput)
			}
		})
	}
}
