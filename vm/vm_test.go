package vm

import (
	"bytes"
	"testing"
)

func TestManipulateMemoryPointer(t *testing.T) {
	type test struct {
		source                string
		expectedMemoryPointer uint16
		tag                   string
	}
	tests := []test{
		{source: ">>>", expectedMemoryPointer: 3, tag: "A"},
		{source: "<<<", expectedMemoryPointer: 29997, tag: "B"},
		{source: "<>>", expectedMemoryPointer: 1, tag: "C"},
	}
	for _, tc := range tests {
		vm := NewVM(tc.source, new(bytes.Buffer), new(bytes.Buffer), true)
		vm.Run()
		actualMemoryPointer := vm.memoryPointer
		if actualMemoryPointer != tc.expectedMemoryPointer {
			t.Errorf("expected:%d actual:%d (%s)", tc.expectedMemoryPointer, actualMemoryPointer, tc.tag)
		}
	}
}

func TestManipulateValue(t *testing.T) {
	type test struct {
		source        string
		expectedValue uint8
		tag           string
	}
	tests := []test{
		{source: "+++", expectedValue: 3, tag: "A"},
		{source: "---", expectedValue: 253, tag: "B"},
		{source: "-++", expectedValue: 1, tag: "C"},
	}
	for _, tc := range tests {
		vm := NewVM(tc.source, new(bytes.Buffer), new(bytes.Buffer), true)
		vm.Run()
		actualValue := vm.memory[0]
		if actualValue != tc.expectedValue {
			t.Errorf("expected:%d actual:%d (%s)", tc.expectedValue, actualValue, tc.tag)
		}
	}
}

func TestOutput(t *testing.T) {
	type test struct {
		source         string
		memory         [30000]uint8
		expectedOutput string
		tag            string
	}
	tests := []test{
		{source: ".", memory: [30000]uint8{65}, expectedOutput: "A", tag: "A"},
		{source: ".", memory: [30000]uint8{122}, expectedOutput: "z", tag: "B"},
		{source: ".>.>.>.>.>.>.>.>.>.", memory: [30000]uint8{72, 101, 108, 108, 111, 32, 84, 101, 115, 116}, expectedOutput: "Hello Test", tag: "C"},
	}
	for _, tc := range tests {
		output := new(bytes.Buffer)
		vm := NewVM(tc.source, new(bytes.Buffer), output, true)
		vm.memory = tc.memory
		vm.Run()
		actualOutput := output.String()
		if actualOutput != tc.expectedOutput {
			t.Errorf("expected:%s actual:%s (%s)", tc.expectedOutput, actualOutput, tc.tag)
		}
	}
}

func TestInput(t *testing.T) {
	type test struct {
		source                   string
		input                    string
		expectedMemoryFirstBytes []uint8
		tag                      string
	}
	tests := []test{
		{source: ",", input: "A", expectedMemoryFirstBytes: []uint8{65}, tag: "A"},
		{source: ",", input: "z", expectedMemoryFirstBytes: []uint8{122}, tag: "B"},
		{source: ",>,>,>,>,>,>,>,>,>,", input: "Hello Test", expectedMemoryFirstBytes: []uint8{72, 101, 108, 108, 111, 32, 84, 101, 115, 116}, tag: "C"},
	}
	for _, tc := range tests {
		input := bytes.NewBufferString(tc.input)
		vm := NewVM(tc.source, input, new(bytes.Buffer), true)
		vm.Run()
		actualMemoryFirstBytes := vm.memory[:len(tc.expectedMemoryFirstBytes)]
		if !bytes.Equal(actualMemoryFirstBytes, tc.expectedMemoryFirstBytes) {
			t.Errorf("expected:%d actual:%d (%s)", tc.expectedMemoryFirstBytes, actualMemoryFirstBytes, tc.tag)
		}
	}
}

func TestLooping(t *testing.T) {
	type test struct {
		source                   string
		expectedMemoryFirstBytes []uint8
		tag                      string
	}
	tests := []test{
		{source: "+>++[<+>-]", expectedMemoryFirstBytes: []uint8{3}, tag: "A"},      // 1 + 2 = 3
		{source: "++[>+++<-]", expectedMemoryFirstBytes: []uint8{0, 6}, tag: "B"},   // 2 * 3 = 6
		{source: "----[>+<-]", expectedMemoryFirstBytes: []uint8{0, 252}, tag: "C"}, // move value from memory[0] to memory[1]
		{source: ">+<[>[+]]+", expectedMemoryFirstBytes: []uint8{1, 1}, tag: "D"},   // required for full test coverage...
	}
	for _, tc := range tests {
		vm := NewVM(tc.source, new(bytes.Buffer), new(bytes.Buffer), true)
		vm.Run()
		actualMemoryFirstBytes := vm.memory[:len(tc.expectedMemoryFirstBytes)]
		if !bytes.Equal(actualMemoryFirstBytes, tc.expectedMemoryFirstBytes) {
			t.Errorf("expected:%d actual:%d (%s)", tc.expectedMemoryFirstBytes, actualMemoryFirstBytes, tc.tag)
		}
	}
}
