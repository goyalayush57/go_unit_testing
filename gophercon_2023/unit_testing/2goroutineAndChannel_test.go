package unit_testing

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sync"
	"testing"
)

func Test_worker(t *testing.T) {
	originalStdout := os.Stdout
	var mockReader *os.File
	var mockWriter *os.File
	mockTasks := make(chan Task, 2)
	wg := &sync.WaitGroup{}

	type args struct {
		id    int
		tasks <-chan Task
		wg    *sync.WaitGroup
	}
	tests := []struct {
		name           string
		args           args
		preSetup       func()
		postSetup      func()
		expectedOutput string
	}{
		{
			name: "Checking for single task",
			args: args{
				id:    1,
				tasks: mockTasks,
				wg:    wg,
			},
			preSetup: func() {
				wg.Add(1)
				// Capture standard output
				mockReader, mockWriter, _ = os.Pipe()
				os.Stdout = mockWriter
				expectedTask := Task{ID: 1, Data: "Test Task 1"}
				mockTasks <- expectedTask
				close(mockTasks)
			},
			postSetup: func() {
				wg.Wait()
				// Restore stdout
				mockWriter.Close()
				os.Stdout = originalStdout
			},
			expectedOutput: fmt.Sprintf("Worker 1 processing task %d: %s\n", 1, "Test Task 1"),
		},
	}
	for _, tt := range tests {
		tt.preSetup()
		t.Run(tt.name, func(t *testing.T) {
			go worker(tt.args.id, tt.args.tasks, tt.args.wg)
			tt.postSetup()
			// Get the captured output
			var capturedOutput bytes.Buffer
			io.Copy(&capturedOutput, mockReader)
			// Compare the captured output with the expected output
			if capturedOutput.String() != tt.expectedOutput {
				t.Errorf("Expected output:\n%s\nActual output:\n%s", tt.expectedOutput, capturedOutput.String())
			}
		})

	}
}

/*
Post Setup

1. Use it in case of resources that needs closure or swap


*/
