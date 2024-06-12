package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/go-vgo/robotgo"
)

type ScriptProcessor struct {
	keyMap          map[string]string
	pressedKeys     map[string]bool
	savedColor      string
	executeNextLine bool
	variables       map[string]int
	labels          map[string]int
	returnStack     []int
}

func NewScriptProcessor() *ScriptProcessor {
	sp := &ScriptProcessor{
		keyMap:          make(map[string]string),
		pressedKeys:     make(map[string]bool),
		variables:       make(map[string]int),
		labels:          make(map[string]int),
		executeNextLine: true,
		returnStack:     []int{},
	}
	sp.keyMap = initializeKeyMap()
	return sp
}

func (sp *ScriptProcessor) executeScript(lines []string) {
	// First pass: store labels and their line numbers
	for i, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "#") {
			sp.labels[line[1:]] = i
		}
	}

	// Second pass: execute commands
	for i := 0; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if line == "" || strings.HasPrefix(line, "#") || strings.HasPrefix(line, ";") {
			continue
		}

		if !sp.executeNextLine {
			sp.executeNextLine = true
			continue
		}

		parts := strings.SplitN(line, ":", 2)
		command := parts[0]
		args := []string{}
		if len(parts) > 1 {
			args = strings.Split(parts[1], ",")
		}

		lineNumber := i + 1
		newLine, err := sp.executeCommand(command, args, lineNumber)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(0)
		}
		if newLine >= 0 {
			i = newLine - 1
		}
	}
}

func (sp *ScriptProcessor) executeCommand(command string, args []string, lineNumber int) (int, error) {
	switch command {
	case "println":
		if len(args) < 1 {
			err := fmt.Errorf("error on line %d: println command requires at least 1 argument", lineNumber)
			return 0, err
		}
		output, err := sp.concatenateStringArguments(args, lineNumber)
		if err != nil {
			return 0, err
		}
		fmt.Println(output)

	case "print":
		if len(args) < 1 {
			err := fmt.Errorf("error on line %d: print command requires at least 1 argument", lineNumber)
			return 0, err
		}
		output, err := sp.concatenateStringArguments(args, lineNumber)
		if err != nil {
			return 0, err
		}
		fmt.Print(output)
	case "move":
		if len(args) != 2 {
			err := fmt.Errorf("error on line %d: move command requires exactly 2 arguments", lineNumber)
			return 0, err
		}
		x, _ := strconv.Atoi(strings.TrimSpace(args[0]))
		y, _ := strconv.Atoi(strings.TrimSpace(args[1]))
		robotgo.Move(x, y)
	case "autopress":
		if len(args) < 1 {
			err := fmt.Errorf("error on line %d: autopress command requires at least 1 argument", lineNumber)
			return 0, err
		}
		time.Sleep(80 * time.Millisecond)
		for _, arg := range args {
			arg = strings.TrimSpace(arg)
			key := sp.keyMap[arg]
			if key != "" {
				if key == "left" || key == "right" {
					robotgo.Click(key, false)
				} else {
					robotgo.KeyTap(key)
				}
				sp.pressedKeys[arg] = true
			} else {
				err := fmt.Errorf("error on line %d: Invalid key: %s", lineNumber, arg)
				return 0, err
			}
			time.Sleep(80 * time.Millisecond)
		}
		time.Sleep(80 * time.Millisecond)
		for _, arg := range args {
			arg = strings.TrimSpace(arg)
			key := sp.keyMap[arg]
			if key != "" {
				if key == "left" || key == "right" {
					robotgo.Click(key, true)
				} else {
					robotgo.KeyTap(key)
				}
				delete(sp.pressedKeys, arg)
			} else {
				err := fmt.Errorf("error on line %d: Invalid key: %s", lineNumber, arg)
				return 0, err
			}
			time.Sleep(80 * time.Millisecond)
		}
		time.Sleep(80 * time.Millisecond)
	case "press":
		if len(args) < 1 {
			err := fmt.Errorf("error on line %d: press command requires at least 1 argument", lineNumber)
			return 0, err
		}
		time.Sleep(40 * time.Millisecond)
		for _, arg := range args {
			arg = strings.TrimSpace(arg)
			key := sp.keyMap[arg]
			if key != "" {
				if key == "left" || key == "right" {
					robotgo.Click(key, false)
				} else {
					robotgo.KeyTap(key)
				}
				sp.pressedKeys[arg] = true
			} else {
				err := fmt.Errorf("error on line %d: Invalid key: %s", lineNumber, arg)
				return 0, err
			}
			time.Sleep(40 * time.Millisecond)
		}
	case "release":
		if len(args) < 1 {
			err := fmt.Errorf("error on line %d: release command requires at least 1 argument", lineNumber)
			return 0, err
		}
		time.Sleep(40 * time.Millisecond)
		for _, arg := range args {
			arg = strings.TrimSpace(arg)
			key := sp.keyMap[arg]
			if key != "" {
				if key == "left" || key == "right" {
					robotgo.Click(key, true)
				} else {
					robotgo.KeyTap(key)
				}
				delete(sp.pressedKeys, arg)
			} else {
				err := fmt.Errorf("error on line %d: Invalid key: %s", lineNumber, arg)
				return 0, err
			}
			time.Sleep(40 * time.Millisecond)
		}
	case "ifpressed":
		if len(args) != 1 {
			err := fmt.Errorf("error on line %d: ifpressed command requires exactly 1 argument", lineNumber)
			return 0, err
		}
		keystr := strings.TrimSpace(args[0])
		if !sp.pressedKeys[keystr] {
			sp.executeNextLine = false
		}
	case "ifnotpressed":
		if len(args) != 1 {
			err := fmt.Errorf("error on line %d: ifnotpressed command requires exactly 1 argument", lineNumber)
			return 0, err
		}
		keystr := strings.TrimSpace(args[0])
		if sp.pressedKeys[keystr] {
			sp.executeNextLine = false
		}
	case "wait":
		if len(args) != 1 {
			err := fmt.Errorf("error on line %d: wait command requires exactly 1 argument", lineNumber)
			return 0, err
		}
		duration, _ := strconv.Atoi(strings.TrimSpace(args[0]))
		time.Sleep(time.Duration(duration) * time.Millisecond)
	case "goto":
		if len(args) != 1 {
			err := fmt.Errorf("error on line %d: goto command requires exactly 1 argument", lineNumber)
			return 0, err
		}
		label := strings.TrimSpace(args[0])
		if line, ok := sp.labels[label]; ok {
			return line, nil
		} else {
			err := fmt.Errorf("error on line %d: Undefined label: %s", lineNumber, label)
			return 0, err
		}

	case "gosub":
		if len(args) != 1 {
			err := fmt.Errorf("error on line %d: gosub command requires exactly 1 argument", lineNumber)
			return 0, err
		}
		label := strings.TrimSpace(args[0])
		if line, ok := sp.labels[label]; ok {
			sp.returnStack = append(sp.returnStack, lineNumber) // Save the return address
			return line, nil                                    // Jump to the label
		} else {
			err := fmt.Errorf("error on line %d: Undefined label: %s", lineNumber, label)
			return 0, err
		}

	case "return":
		if len(args) != 0 {
			err := fmt.Errorf("error on line %d: return command requires no arguments", lineNumber)
			return 0, err
		}
		if len(sp.returnStack) == 0 {
			err := fmt.Errorf("error on line %d: return command called with an empty return stack", lineNumber)
			return 0, err
		}
		returnAddress := sp.returnStack[len(sp.returnStack)-1]
		sp.returnStack = sp.returnStack[:len(sp.returnStack)-1] // Pop the return address
		return returnAddress, nil
	case "set":
		if len(args) != 2 {
			err := fmt.Errorf("error on line %d: set command requires exactly 2 arguments", lineNumber)
			return 0, err
		}
		varName := strings.TrimSpace(args[0])
		value, _ := strconv.Atoi(strings.TrimSpace(args[1]))
		sp.variables[varName] = value
	case "add":
		if len(args) != 2 {
			err := fmt.Errorf("error on line %d: add command requires exactly 2 arguments", lineNumber)
			return 0, err
		}
		varName := strings.TrimSpace(args[0])
		if value, ok := sp.variables[varName]; ok {
			addValue, _ := strconv.Atoi(strings.TrimSpace(args[1]))
			sp.variables[varName] = value + addValue
		} else {
			err := fmt.Errorf("error on line %d: Variable not declared: %s", lineNumber, varName)
			return 0, err
		}
	case "sub":
		if len(args) != 2 {
			err := fmt.Errorf("error on line %d: sub command requires exactly 2 arguments", lineNumber)
			return 0, err
		}
		varName := strings.TrimSpace(args[0])
		if value, ok := sp.variables[varName]; ok {
			subValue, _ := strconv.Atoi(strings.TrimSpace(args[1]))
			sp.variables[varName] = value - subValue
		} else {
			err := fmt.Errorf("error on line %d: Variable not declared: %s", lineNumber, varName)
			return 0, err
		}
	case "ifequal":
		if len(args) != 2 {
			err := fmt.Errorf("error on line %d: ifequal command requires exactly 2 arguments", lineNumber)
			return 0, err
		}
		varName := strings.TrimSpace(args[0])
		if value, ok := sp.variables[varName]; ok {
			compareValue, _ := strconv.Atoi(strings.TrimSpace(args[1]))
			if value != compareValue {
				sp.executeNextLine = false
			}
		} else {
			err := fmt.Errorf("error on line %d: Variable not declared: %s", lineNumber, varName)
			return 0, err
		}
	case "ifgreater":
		if len(args) != 2 {
			err := fmt.Errorf("error on line %d: ifgreater command requires exactly 2 arguments", lineNumber)
			return 0, err
		}
		varName := strings.TrimSpace(args[0])
		if value, ok := sp.variables[varName]; ok {
			compareValue, _ := strconv.Atoi(strings.TrimSpace(args[1]))
			if value <= compareValue {
				sp.executeNextLine = false
			}
		} else {
			err := fmt.Errorf("error on line %d: Variable not declared: %s", lineNumber, varName)
			return 0, err
		}
	case "ifless":
		if len(args) != 2 {
			err := fmt.Errorf("error on line %d: ifless command requires exactly 2 arguments", lineNumber)
			return 0, err
		}
		varName := strings.TrimSpace(args[0])
		if value, ok := sp.variables[varName]; ok {
			compareValue, _ := strconv.Atoi(strings.TrimSpace(args[1]))
			if value >= compareValue {
				sp.executeNextLine = false
			}
		} else {
			err := fmt.Errorf("error on line %d: Variable not declared: %s", lineNumber, varName)
			return 0, err
		}
	case "savecolor":
		if len(args) != 2 {
			err := fmt.Errorf("error on line %d: savecolor command requires exactly 2 arguments", lineNumber)
			return 0, err
		}
		x, errX := strconv.Atoi(strings.TrimSpace(args[0]))
		y, errY := strconv.Atoi(strings.TrimSpace(args[1]))
		if errX != nil || errY != nil {
			err := fmt.Errorf("error on line %d: invalid arguments for savecolor command", lineNumber)
			return 0, err
		}
		sp.savedColor = robotgo.GetPixelColor(x, y)

	case "printcolorrgb":
		if len(args) != 0 {
			err := fmt.Errorf("error on line %d: printcolorrgb command requires no arguments", lineNumber)
			return 0, err
		}
		if sp.savedColor == "" {
			err := fmt.Errorf("error on line %d: No color saved, use savecolor command first", lineNumber)
			return 0, err
		}
		r, g, b := hexToRGB(sp.savedColor)
		fmt.Printf("RGB(%d, %d, %d)", r, g, b)

	case "printcolorhex":
		if len(args) != 0 {
			err := fmt.Errorf("error on line %d: printcolorhex command requires no arguments", lineNumber)
			return 0, err
		}
		if sp.savedColor == "" {
			err := fmt.Errorf("error on line %d: No color saved, use savecolor command first", lineNumber)
			return 0, err
		}
		fmt.Printf("#%s", sp.savedColor)

	case "ifcolor":
		if len(args) != 2 {
			err := fmt.Errorf("error on line %d: ifcolor command requires exactly 2 arguments", lineNumber)
			return 0, err
		}
		if sp.savedColor == "" {
			err := fmt.Errorf("error on line %d: No color saved, use savecolor command first", lineNumber)
			return 0, err
		}

		colorHex := strings.TrimSpace(args[0])
		colorHex = strings.TrimPrefix(colorHex, "#")
		if len(colorHex) != 6 {
			err := fmt.Errorf("error on line %d: Color hex must be exactly 6 characters, but got the number: %s", lineNumber, args[0])
			return 0, err
		}

		thresholdHex := strings.TrimSpace(args[1])
		thresholdHex = strings.TrimPrefix(thresholdHex, "#")
		if len(thresholdHex) != 2 {
			err := fmt.Errorf("error on line %d: Threshold hex must be exactly 2 characters, but got the number: %s", lineNumber, args[1])
			return 0, err
		}

		threshold, err := strconv.ParseInt(thresholdHex, 16, 0)
		if err != nil {
			err := fmt.Errorf("error on line %d: Threshold is not a valid hexadecimal number: %s", lineNumber, args[1])
			return 0, err
		}

		if !sp.colorsMatch(sp.savedColor, colorHex, int(threshold)) {
			sp.executeNextLine = false
		}

	default:
		err := fmt.Errorf("error on line %d: Unknown command: %s", lineNumber, command)
		return 0, err

	}
	return -1, nil
}

func (sp *ScriptProcessor) concatenateStringArguments(args []string, lineNumber int) (string, error) {
	var output strings.Builder
	for _, arg := range args {
		arg = strings.TrimSpace(arg)
		if strings.HasPrefix(arg, "\"") && strings.HasSuffix(arg, "\"") {
			// Argument is a string literal
			output.WriteString(strings.Trim(arg, "\""))
		} else {
			// Argument is a variable
			if value, ok := sp.variables[arg]; ok {
				output.WriteString(strconv.Itoa(value))
			} else {
				err := fmt.Errorf("error on line %d: Variable not declared: %s", lineNumber, arg)
				return "", err
			}
		}
	}
	return output.String(), nil
}
