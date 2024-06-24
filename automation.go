package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/go-vgo/robotgo"
)

func customFunction_Move(args ...interface{}) interface{} {
	if len(args) != 2 {
		return fmt.Errorf("move requires exactly 2 arguments")
	}

	// Using type assertions to check if x and y are of type int
	x, ok1 := args[0].(int)
	y, ok2 := args[1].(int)

	if !ok1 || !ok2 {
		return fmt.Errorf("both arguments must be of type int, got %T and %T", x, y)
	}

	robotgo.Move(x, y)
	return nil
}

func customFunction_MouseClick(args ...interface{}) interface{} {
	if len(args) != 1 {
		return fmt.Errorf("mouseclick requires exactly 1 argument")
	}

	// Using type assertions to check if variables are of correct type
	button, ok := args[0].(string)

	if !ok {
		return fmt.Errorf("argument must be of type string, got: %T", button)
	}

	time.Sleep(time.Duration(80) * time.Millisecond)
	robotgo.Toggle(button, "down")
	time.Sleep(time.Duration(100) * time.Millisecond)
	robotgo.Toggle(button, "up")
	time.Sleep(time.Duration(80) * time.Millisecond)
	return nil
}

func customFunction_KeyTap(args ...interface{}) interface{} {
	if len(args) < 1 {
		return fmt.Errorf("keytap requires at least 1 argument")
	}

	var keys []string

	for _, arg := range args {
		switch key := arg.(type) {
		case string:
			keys = append(keys, key)
		default:
			return fmt.Errorf("unsupported argument type: %T", arg)
		}
	}

	if len(keys) == 0 {
		return fmt.Errorf("no valid keys provided")
	}

	if len(keys) == 1 {
		robotgo.KeyTap(keys[0])
	} else if len(keys) > 1 {
		robotgo.KeyTap(keys[0], keys[1:])
	}

	return nil
}

func customFunction_KeyPress(args ...interface{}) interface{} {
	if len(args) < 1 {
		return fmt.Errorf("keypress requires at least 1 argument")
	}

	var keys []string

	for _, arg := range args {
		switch key := arg.(type) {
		case string:
			keys = append(keys, key)
		default:
			return fmt.Errorf("unsupported argument type: %T", arg)
		}
	}

	if len(keys) == 0 {
		return fmt.Errorf("no valid keys provided")
	}

	if len(keys) == 1 {
		robotgo.KeyToggle(keys[0])
	} else if len(keys) > 1 {
		robotgo.KeyToggle(keys[0], keys[1:])
	}

	return nil
}

func customFunction_KeyRelease(args ...interface{}) interface{} {
	if len(args) < 1 {
		return fmt.Errorf("keyrelease requires at least 1 argument")
	}

	var keys []string

	for _, arg := range args {
		switch key := arg.(type) {
		case string:
			keys = append(keys, key)
		default:
			return fmt.Errorf("unsupported argument type: %T", arg)
		}
	}

	if len(keys) == 0 {
		return fmt.Errorf("no valid keys provided")
	}

	if len(keys) == 1 {
		robotgo.KeyToggle(keys[0])
	} else if len(keys) > 1 {
		robotgo.KeyToggle(keys[0], keys[1:])
	}

	return nil
}

func customFunction_GetColor(args ...interface{}) interface{} {
	if len(args) != 2 {
		return fmt.Errorf("getcolor requires exactly 3 argument")
	}

	// Using type assertions to check if variables are of correct type
	x, ok1 := args[1].(int)
	y, ok2 := args[2].(int)

	if !ok1 || !ok2 {
		return fmt.Errorf("argument for x and y must be of type int, got %T and %T", x, y)
	}

	color := robotgo.GetPixelColor(x, y)

	return color
}

func customFunction_ColorMatch(args ...interface{}) interface{} {
	if len(args) != 4 {
		return fmt.Errorf("getcolor requires exactly 4 argument")
	}

	// Using type assertions to check if variables are of correct type
	color1, ok1 := args[1].(string)
	color2, ok2 := args[2].(string)
	threshold, ok3 := args[3].(string)

	if !ok1 || !ok2 {
		return fmt.Errorf("argument for color 1 and color 2 must be of type string, got %T and %T", color1, color2)
	}

	if !ok3 {
		return fmt.Errorf("argument for threshold must be of type string, got %T", threshold)
	}

	color1 = strings.TrimSpace(color1)
	color2 = strings.TrimSpace(color2)
	threshold = strings.TrimSpace(threshold)

	// Validate colors
	color1 = strings.TrimPrefix(color1, "#")
	color2 = strings.TrimPrefix(color2, "#")
	if len(color1) != 6 || len(color2) != 6 {
		return fmt.Errorf("colors hex must be exactly 6 characters")
	}

	// Validate threshold
	threshold = strings.TrimPrefix(threshold, "#")
	if len(threshold) != 2 {
		return fmt.Errorf("threshold hex must be exactly 2 characters")
	}

	thresholdInt, err := strconv.ParseInt(threshold, 16, 0)
	if err != nil {
		return fmt.Errorf("threshold is not a valid hexadecimal number: %s", threshold)
	}

	returnResult := false

	if colorsMatch(color1, color2, int(thresholdInt)) {
		returnResult = true
	}

	return returnResult
}
