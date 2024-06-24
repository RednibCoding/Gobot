package main

import (
	"fmt"
	"os"

	"github.com/RednibCoding/runevm"
)

func main() {
	args := []string{"gobot.exe", "test.tin"}
	// args := os.Args
	if len(args) < 2 {
		fmt.Println("Usage: gobot <script-file>")
		os.Exit(1)
	}
	source, err := os.ReadFile(args[1])
	if err != nil {
		fmt.Printf("ERROR: Can't find source file '%s'.\n", args[1])
		os.Exit(1)
	}

	filepath := args[1]

	vm := runevm.NewRuneVM()

	vm.SetFun("move", customFunction_Move)
	vm.SetFun("mouseclick", customFunction_MouseClick)
	vm.SetFun("keytap", customFunction_KeyTap)
	vm.SetFun("keypress", customFunction_KeyPress)
	vm.SetFun("keyrelease", customFunction_KeyRelease)
	vm.SetFun("getcolor", customFunction_GetColor)
	vm.SetFun("colormatch", customFunction_ColorMatch)

	vm.Run(string(source), filepath)

	// vm := tinvm.New()

	// vm.Run(string(source), args[1])
}
