package main

import (
	"fmt"
	"os"

	"github.com/RednibCoding/tinvm"
)

func main() {
	// args := []string{"gobot.exe", "test.tin"}
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Usage: gobot <script-file>")
		os.Exit(1)
	}
	source, err := os.ReadFile(args[1])
	if err != nil {
		fmt.Printf("ERROR: Can't find source file '%s'.\n", args[1])
		os.Exit(1)
	}

	vm := tinvm.New()

	vm.AddFunction("move", customFunction_Move)
	vm.AddFunction("mouseclick", customFunction_MouseClick)
	vm.AddFunction("keytap", customFunction_KeyTap)
	vm.AddFunction("keypress", customFunction_KeyPress)
	vm.AddFunction("keyrelease", customFunction_KeyRelease)
	vm.AddFunction("getcolor", customFunction_GetColor)
	vm.AddFunction("colormatch", customFunction_ColorMatch)

	vm.Run(string(source), args[1])
}
