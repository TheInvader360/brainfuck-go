package vm

import (
	"fmt"
	"io"
)

type VM struct {
	source             string       // bf source (e.g. "--[+++++++<---->>-->+>+>+<<<<]<.>++++[-<++++<++>>>->--<<]>>-.>--..>+.<<<.<<-.>>+>->>.+++[.<]<<++.")
	input              io.Reader    // input stream
	output             io.Writer    // output stream
	debug              bool         // print out debug info if true
	memory             [30000]uint8 // memory array - addresses 0 to 29999, each address stores an unsigned integer between 0 and 255
	memoryPointer      uint16       // the "current" memory address (0 to 29999)
	instructionPointer int          // the index of the "current" instruction in the source
}

func NewVM(source string, input io.Reader, output io.Writer, debug bool) *VM {
	return &VM{
		source: source,
		input:  input,
		output: output,
		debug:  debug,
	}
}

func (vm *VM) Run() {
	for vm.instructionPointer < len(vm.source) {
		instruction := vm.source[vm.instructionPointer]
		switch instruction {
		case '>': // increment the memoryPointer (wraps to 0)
			if vm.memoryPointer < 29999 {
				vm.memoryPointer++
			} else {
				vm.memoryPointer = 0
			}
		case '<': // decrement the memoryPointer (wraps to 29999)
			if vm.memoryPointer > 0 {
				vm.memoryPointer--
			} else {
				vm.memoryPointer = 29999
			}
		case '+': // increment the value stored at the "current" memory address (wraps to 0)
			if vm.memory[vm.memoryPointer] < 255 {
				vm.memory[vm.memoryPointer]++
			} else {
				vm.memory[vm.memoryPointer] = 0
			}
		case '-': // decrement the value stored at the "current" memory address (wraps to 255)
			if vm.memory[vm.memoryPointer] > 0 {
				vm.memory[vm.memoryPointer]--
			} else {
				vm.memory[vm.memoryPointer] = 255
			}
		case '.': // output the value stored at the "current" memory address (ignores errors)
			output := vm.memory[vm.memoryPointer]
			vm.output.Write([]byte{output})
			if vm.debug {
				fmt.Printf("\toutput: %v\n", output)
			}
		case ',': // store input value at the "current" memory address (ignores errors)
			input := make([]byte, 1)
			vm.input.Read(input)
			vm.memory[vm.memoryPointer] = uint8(input[0])
			if vm.debug {
				fmt.Printf("\tinput: %v\n", input)
			}
		case '[': // if the value stored at the "current" memory address is zero, jump to the instruction after the matching "]"
			if vm.memory[vm.memoryPointer] == 0 {
				vm.scan(1)
			}
		case ']': // if the value stored at the "current" memory address is not zero, jump to the instruction after the matching "["
			if vm.memory[vm.memoryPointer] != 0 {
				vm.scan(-1)
			}
		}

		if vm.debug {
			fmt.Printf("\t%d\t%c\tmemory[%d]=%d\n", vm.instructionPointer, vm.source[vm.instructionPointer], vm.memoryPointer, vm.memory[vm.memoryPointer])
		}
		vm.instructionPointer++
	}
}

func (vm *VM) scan(direction int) { // see https://github.com/golang/go/tree/master/test/turing.go
	for nest := direction; direction*nest > 0; vm.instructionPointer += direction {
		switch vm.source[vm.instructionPointer+direction] {
		case ']':
			nest--
		case '[':
			nest++
		}
	}
}
