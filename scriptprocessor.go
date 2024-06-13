package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/go-vgo/robotgo"
)

type VarType int

const (
	Str VarType = iota
	Int
	Flt
)

func (vt VarType) String() string {
	switch vt {
	case Int:
		return "Integer"
	case Flt:
		return "Float"
	case Str:
		return "String"
	default:
		return "Unknown"
	}
}

type Variable struct {
	Type  VarType
	Value string
}

type ScriptProcessor struct {
	keyMap          map[string]string
	pressedKeys     map[string]bool
	executeNextLine bool
	variables       map[string]Variable
	labels          map[string]int
	anonymousLabels []int
	returnStack     []int
}

func NewScriptProcessor() *ScriptProcessor {
	sp := &ScriptProcessor{
		keyMap:          make(map[string]string),
		pressedKeys:     make(map[string]bool),
		variables:       make(map[string]Variable),
		labels:          make(map[string]int),
		executeNextLine: true,
		anonymousLabels: []int{},
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
			if line == "#@" {
				sp.anonymousLabels = append(sp.anonymousLabels, i)
			} else {
				sp.labels[line[1:]] = i
			}
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
			args = splitArgs(parts[1])
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

	case "print", "println":
		output := ""
		for _, arg := range args {
			arg = strings.TrimSpace(arg)
			if strings.HasPrefix(arg, "\"") && strings.HasSuffix(arg, "\"") {
				// It's a string literal, so strip the quotes
				output += arg[1 : len(arg)-1]
			} else if variable, exists := sp.variables[arg]; exists {
				// It's a variable, so use its value
				output += variable.Value
			} else {
				// It's a number literal value, so use it as is
				output += arg
			}
		}
		if command == "println" {
			fmt.Println(output)
		} else {
			fmt.Print(output)
		}
		return -1, nil

	case "move":
		if len(args) != 2 {
			return 0, fmt.Errorf("error on line %d: move command requires exactly 2 arguments", lineNumber)
		}

		x, errX := sp.resolveInt(args[0], lineNumber)
		if errX != nil {
			return 0, fmt.Errorf("error on line %d: invalid integer value for x-coordinate: %s", lineNumber, args[0])
		}

		y, errY := sp.resolveInt(args[1], lineNumber)
		if errY != nil {
			return 0, fmt.Errorf("error on line %d: invalid integer value for y-coordinate: %s", lineNumber, args[1])
		}

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
					robotgo.Toggle(key)
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
					robotgo.Toggle(key, "up")
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
					robotgo.Toggle(key)
				} else {
					robotgo.KeyToggle(key)
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
				if key == "left" || key == "right" || key == "center" {
					robotgo.Toggle(key, "up")
				} else {
					robotgo.KeyToggle(key, "up")
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

		if strings.HasPrefix(label, "@f") {
			steps := len(label) - 1 // Number of 'f' characters
			count := 0
			for i := 0; i < len(sp.anonymousLabels); i++ {
				if sp.anonymousLabels[i] > lineNumber {
					count++
					if count == steps {
						return sp.anonymousLabels[i], nil
					}
				}
			}
			err := fmt.Errorf("error on line %d: No sufficient forward anonymous labels found", lineNumber)
			return 0, err
		} else if strings.HasPrefix(label, "@b") {
			steps := len(label) - 1 // Number of 'b' characters
			count := 0
			for i := len(sp.anonymousLabels) - 1; i >= 0; i-- {
				if sp.anonymousLabels[i] < lineNumber {
					count++
					if count == steps {
						return sp.anonymousLabels[i], nil
					}
				}
			}
			err := fmt.Errorf("error on line %d: No sufficient backward anonymous labels found", lineNumber)
			return 0, err
		} else if line, ok := sp.labels[label]; ok {
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

	case "goclr":
		if len(args) != 0 {
			err := fmt.Errorf("error on line %d: goclr command requires no arguments", lineNumber)
			return 0, err
		}
		// clear the return stack
		sp.returnStack = sp.returnStack[:0]

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
			return 0, fmt.Errorf("error on line %d: set command requires exactly 2 arguments", lineNumber)
		}
		varName := strings.TrimSpace(args[0])
		varValue := strings.TrimSpace(args[1])

		// Check if varValue is another variable
		if refVar, exists := sp.variables[varValue]; exists {
			varValue = refVar.Value
			varType := refVar.Type
			// Check if variable exists and its type
			if variable, exists := sp.variables[varName]; exists {
				if variable.Type != varType {
					return 0, fmt.Errorf("error on line %d: cannot redefine variable %s with a different type", lineNumber, varName)
				}
			}
			sp.variables[varName] = Variable{Type: varType, Value: varValue}
			return -1, nil
		}

		// Determine the type of the variable
		varType := Str
		if strings.HasPrefix(varValue, "\"") && strings.HasSuffix(varValue, "\"") {
			// String type
			varValue = varValue[1 : len(varValue)-1]
		} else if strings.Contains(varValue, ".") {
			// Float type
			varType = Flt
			if _, err := strconv.ParseFloat(varValue, 64); err != nil {
				return 0, fmt.Errorf("error on line %d: invalid float value", lineNumber)
			}
		} else {
			// Integer type
			varType = Int
			if _, err := strconv.Atoi(varValue); err != nil {
				return 0, fmt.Errorf("error on line %d: invalid integer value", lineNumber)
			}
		}

		// Check if variable exists and its type
		if variable, exists := sp.variables[varName]; exists {
			if variable.Type != varType {
				return 0, fmt.Errorf("error on line %d: cannot redefine variable %s with a different type", lineNumber, varName)
			}
		}

		sp.variables[varName] = Variable{Type: varType, Value: varValue}
		return -1, nil

	case "add":
		if len(args) != 2 {
			err := fmt.Errorf("error on line %d: add command requires exactly 2 arguments", lineNumber)
			return 0, err
		}
		varName := strings.TrimSpace(args[0])
		varValue := strings.TrimSpace(args[1])

		variable, exists := sp.variables[varName]
		if !exists {
			err := fmt.Errorf("error on line %d: variable %s not defined", lineNumber, varName)
			return 0, err
		}

		otherVariable, exists := sp.variables[varValue]
		if !exists {
			otherVariable = Variable{Type: Int, Value: varValue}
			if strings.HasPrefix(varValue, "\"") && strings.HasSuffix(varValue, "\"") {
				otherVariable.Type = Str
				otherVariable.Value = varValue[1 : len(varValue)-1]
			} else if strings.Contains(varValue, ".") {
				otherVariable.Type = Flt
				if _, err := strconv.ParseFloat(varValue, 64); err != nil {
					return 0, fmt.Errorf("error on line %d: invalid float value", lineNumber)
				}
			} else {
				if _, err := strconv.Atoi(varValue); err != nil {
					return 0, fmt.Errorf("error on line %d: invalid integer value", lineNumber)
				}
			}
		}

		switch variable.Type {
		case Int:
			intValue, err := strconv.Atoi(variable.Value)
			if err != nil {
				return 0, fmt.Errorf("error on line %d: invalid integer value", lineNumber)
			}
			if otherVariable.Type == Int {
				argValue, err := strconv.Atoi(otherVariable.Value)
				if err != nil {
					return 0, fmt.Errorf("error on line %d: invalid integer value", lineNumber)
				}
				intValue += argValue
			} else if otherVariable.Type == Flt {
				argValue, err := strconv.ParseFloat(otherVariable.Value, 64)
				if err != nil {
					return 0, fmt.Errorf("error on line %d: invalid float value", lineNumber)
				}
				intValue += int(argValue)
			} else {
				return 0, fmt.Errorf("error on line %d: cannot add string to integer", lineNumber)
			}
			sp.variables[varName] = Variable{Type: Int, Value: strconv.Itoa(intValue)}

		case Flt:
			floatValue, err := strconv.ParseFloat(variable.Value, 64)
			if err != nil {
				return 0, fmt.Errorf("error on line %d: invalid float value", lineNumber)
			}
			if otherVariable.Type == Int {
				argValue, err := strconv.Atoi(otherVariable.Value)
				if err != nil {
					return 0, fmt.Errorf("error on line %d: invalid integer value", lineNumber)
				}
				floatValue += float64(argValue)
			} else if otherVariable.Type == Flt {
				argValue, err := strconv.ParseFloat(otherVariable.Value, 64)
				if err != nil {
					return 0, fmt.Errorf("error on line %d: invalid float value", lineNumber)
				}
				floatValue += argValue
			} else {
				return 0, fmt.Errorf("error on line %d: cannot add string to float", lineNumber)
			}
			sp.variables[varName] = Variable{Type: Flt, Value: fmt.Sprintf("%f", floatValue)}

		case Str:
			if otherVariable.Type == Str || otherVariable.Type == Int {
				varValue := variable.Value + otherVariable.Value
				sp.variables[varName] = Variable{Type: Str, Value: varValue}
			} else {
				return 0, fmt.Errorf("error on line %d: cannot add float to string", lineNumber)
			}

		default:
			return 0, fmt.Errorf("error on line %d: unsupported variable type for add command", lineNumber)
		}
		return -1, nil

	case "sub":
		if len(args) != 2 {
			err := fmt.Errorf("error on line %d: %s command requires exactly 2 arguments", lineNumber, command)
			return 0, err
		}
		varName := strings.TrimSpace(args[0])
		varValue := strings.TrimSpace(args[1])

		variable, exists := sp.variables[varName]
		if !exists {
			err := fmt.Errorf("error on line %d: variable %s not defined", lineNumber, varName)
			return 0, err
		}

		otherVariable, exists := sp.variables[varValue]
		if !exists {
			otherVariable = Variable{Type: Int, Value: varValue}
			if strings.Contains(varValue, ".") {
				otherVariable.Type = Flt
				if _, err := strconv.ParseFloat(varValue, 64); err != nil {
					return 0, fmt.Errorf("error on line %d: invalid float value", lineNumber)
				}
			} else {
				if _, err := strconv.Atoi(varValue); err != nil {
					return 0, fmt.Errorf("error on line %d: invalid integer value", lineNumber)
				}
			}
		}

		switch variable.Type {
		case Int:
			intValue, err := strconv.Atoi(variable.Value)
			if err != nil {
				return 0, fmt.Errorf("error on line %d: invalid integer value", lineNumber)
			}
			if otherVariable.Type == Int {
				argValue, err := strconv.Atoi(otherVariable.Value)
				if err != nil {
					return 0, fmt.Errorf("error on line %d: invalid integer value", lineNumber)
				}
				if command == "add" {
					intValue += argValue
				} else {
					intValue -= argValue
				}
			} else {
				argValue, err := strconv.ParseFloat(otherVariable.Value, 64)
				if err != nil {
					return 0, fmt.Errorf("error on line %d: invalid float value", lineNumber)
				}
				if command == "add" {
					intValue += int(argValue)
				} else {
					intValue -= int(argValue)
				}
			}
			sp.variables[varName] = Variable{Type: Int, Value: strconv.Itoa(intValue)}

		case Flt:
			floatValue, err := strconv.ParseFloat(variable.Value, 64)
			if err != nil {
				return 0, fmt.Errorf("error on line %d: invalid float value", lineNumber)
			}
			if otherVariable.Type == Int {
				argValue, err := strconv.Atoi(otherVariable.Value)
				if err != nil {
					return 0, fmt.Errorf("error on line %d: invalid integer value", lineNumber)
				}
				if command == "add" {
					floatValue += float64(argValue)
				} else {
					floatValue -= float64(argValue)
				}
			} else {
				argValue, err := strconv.ParseFloat(otherVariable.Value, 64)
				if err != nil {
					return 0, fmt.Errorf("error on line %d: invalid float value", lineNumber)
				}
				if command == "add" {
					floatValue += argValue
				} else {
					floatValue -= argValue
				}
			}
			sp.variables[varName] = Variable{Type: Flt, Value: fmt.Sprintf("%f", floatValue)}

		default:
			return 0, fmt.Errorf("error on line %d: unsupported variable type for %s command", lineNumber, command)
		}
		return -1, nil

	case "ifequal", "ifnotequal", "ifless", "ifgreater":
		if len(args) != 2 {
			return 0, fmt.Errorf("error on line %d: %s command requires exactly 2 arguments", lineNumber, command)
		}
		varName1 := strings.TrimSpace(args[0])
		varName2 := strings.TrimSpace(args[1])

		var1, exists1 := sp.variables[varName1]
		var var2 Variable

		// Check if the second argument is a variable or a constant
		if var2Temp, exists := sp.variables[varName2]; exists {
			var2 = var2Temp
		} else {
			// Determine the type of the second argument
			if strings.HasPrefix(varName2, "\"") && strings.HasSuffix(varName2, "\"") {
				var2 = Variable{Type: Str, Value: varName2[1 : len(varName2)-1]}
			} else if strings.Contains(varName2, ".") {
				if _, err := strconv.ParseFloat(varName2, 64); err == nil {
					var2 = Variable{Type: Flt, Value: varName2}
				} else {
					return 0, fmt.Errorf("error on line %d: invalid float value", lineNumber)
				}
			} else if _, err := strconv.Atoi(varName2); err == nil {
				var2 = Variable{Type: Int, Value: varName2}
			} else {
				return 0, fmt.Errorf("error on line %d: variable %s or constant %s not defined", lineNumber, varName1, varName2)
			}
		}

		if !exists1 {
			return 0, fmt.Errorf("error on line %d: variable %s not defined", lineNumber, varName1)
		}

		// Handle string comparisons for ifequal only
		if var1.Type == Str && var2.Type == Str {
			if command == "ifequal" {
				if var1.Value == var2.Value {
					sp.executeNextLine = true
				} else {
					sp.executeNextLine = false
				}
				return -1, nil
			} else if command == "ifnotequal" {
				if var1.Value != var2.Value {
					sp.executeNextLine = true
				} else {
					sp.executeNextLine = false
				}
				return -1, nil
			} else {
				return 0, fmt.Errorf("error on line %d: %s command not supported for string variables", lineNumber, command)
			}
		}

		// Handle integer and float comparisons
		var1Float, err1 := strconv.ParseFloat(var1.Value, 64)
		var2Float, err2 := strconv.ParseFloat(var2.Value, 64)
		if (var1.Type == Int || var1.Type == Flt) && (var2.Type == Int || var2.Type == Flt) {
			if err1 != nil || err2 != nil {
				return 0, fmt.Errorf("error on line %d: invalid numeric value for comparison", lineNumber)
			}
			switch command {
			case "ifequal":
				if var1Float == var2Float {
					sp.executeNextLine = true
				} else {
					sp.executeNextLine = false
				}
			case "ifnotequal":
				if var1Float != var2Float {
					sp.executeNextLine = true
				} else {
					sp.executeNextLine = false
				}
			case "ifless":
				if var1Float < var2Float {
					sp.executeNextLine = true
				} else {
					sp.executeNextLine = false
				}
			case "ifgreater":
				if var1Float > var2Float {
					sp.executeNextLine = true
				} else {
					sp.executeNextLine = false
				}
			}
			return -1, nil
		}

		return 0, fmt.Errorf("error on line %d: incompatible types for comparison: %s and %s", lineNumber, var1.Type, var2.Type)

	case "getcolor":
		if len(args) != 3 {
			return 0, fmt.Errorf("error on line %d: getcolor command requires exactly 3 arguments", lineNumber)
		}
		varName := strings.TrimSpace(args[0])

		x, errX := sp.resolveInt(args[1], lineNumber)
		if errX != nil {
			return 0, errX
		}

		y, errY := sp.resolveInt(args[2], lineNumber)
		if errY != nil {
			return 0, errY
		}

		colorHex := robotgo.GetPixelColor(x, y)

		variable, exists := sp.variables[varName]
		if exists && variable.Type != Str {
			return 0, fmt.Errorf("error on line %d: variable %s must be of type Str", lineNumber, varName)
		}

		sp.variables[varName] = Variable{Type: Str, Value: "#" + colorHex}
		return -1, nil

	case "colorsmatch":
		if len(args) != 3 {
			return 0, fmt.Errorf("error on line %d: colorsmatch command requires exactly 3 arguments", lineNumber)
		}

		color1 := strings.TrimSpace(args[0])
		color2 := strings.TrimSpace(args[1])
		thresholdHex := strings.TrimSpace(args[2])

		// Check if color1 and color2 are variables or string literals
		if var1, exists := sp.variables[color1]; exists && var1.Type == Str {
			color1 = var1.Value
		} else if strings.HasPrefix(color1, "\"") && strings.HasSuffix(color1, "\"") {
			color1 = color1[1 : len(color1)-1]
		} else {
			return 0, fmt.Errorf("error on line %d: invalid color1 value", lineNumber)
		}

		if var2, exists := sp.variables[color2]; exists && var2.Type == Str {
			color2 = var2.Value
		} else if strings.HasPrefix(color2, "\"") && strings.HasSuffix(color2, "\"") {
			color2 = color2[1 : len(color2)-1]
		} else {
			return 0, fmt.Errorf("error on line %d: invalid color2 value", lineNumber)
		}

		// Validate colors
		color1 = strings.TrimPrefix(color1, "#")
		color2 = strings.TrimPrefix(color2, "#")
		if len(color1) != 6 || len(color2) != 6 {
			return 0, fmt.Errorf("error on line %d: Color hex must be exactly 6 characters", lineNumber)
		}

		// Validate threshold
		thresholdHex = strings.TrimPrefix(thresholdHex, "#")
		if len(thresholdHex) != 2 {
			return 0, fmt.Errorf("error on line %d: Threshold hex must be exactly 2 characters", lineNumber)
		}

		threshold, err := strconv.ParseInt(thresholdHex, 16, 0)
		if err != nil {
			return 0, fmt.Errorf("error on line %d: Threshold is not a valid hexadecimal number: %s", lineNumber, args[2])
		}

		if !sp.colorsMatch(color1, color2, int(threshold)) {
			sp.executeNextLine = false
		}

		return -1, nil

	default:
		err := fmt.Errorf("error on line %d: Unknown command: %s", lineNumber, command)
		return 0, err

	}
	return -1, nil
}

// Function to resolve an argument to its integer value
func (sp *ScriptProcessor) resolveInt(arg string, lineNumber int) (int, error) {
	arg = strings.TrimSpace(arg)
	if variable, exists := sp.variables[arg]; exists {
		if variable.Type != Int {
			return 0, fmt.Errorf("error on line %d: variable %s is not an integer", lineNumber, arg)
		}
		return strconv.Atoi(variable.Value)
	}
	if value, err := strconv.Atoi(arg); err == nil {
		return value, nil
	}
	return 0, fmt.Errorf("error on line %d: variable %s not declared", lineNumber, arg)
}
