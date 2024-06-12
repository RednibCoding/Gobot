package main

import (
	"strconv"
	"strings"
)

func (sp *ScriptProcessor) colorsMatch(c1, c2 string, threshold int) bool {
	r1, g1, b1 := hexToRGB(c1)
	r2, g2, b2 := hexToRGB(c2)
	rDiff := abs(r1 - r2)
	gDiff := abs(g1 - g2)
	bDiff := abs(b1 - b2)

	return rDiff <= threshold && gDiff <= threshold && bDiff <= threshold
}

func hexToRGB(hexColor string) (int, int, int) {
	hexColor = strings.TrimPrefix(hexColor, "#")
	r, _ := strconv.ParseInt(hexColor[0:2], 16, 0)
	g, _ := strconv.ParseInt(hexColor[2:4], 16, 0)
	b, _ := strconv.ParseInt(hexColor[4:6], 16, 0)
	return int(r), int(g), int(b)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func splitArgs(argsStr string) []string {
	args := []string{}
	current := strings.Builder{}
	inQuotes := false

	for _, r := range argsStr {
		switch {
		case r == '"' && (len(current.String()) == 0 || current.String()[len(current.String())-1] != '\\'):
			inQuotes = !inQuotes
			current.WriteRune(r)
		case r == ',' && !inQuotes:
			args = append(args, strings.TrimSpace(current.String()))
			current.Reset()
		default:
			current.WriteRune(r)
		}
	}
	args = append(args, strings.TrimSpace(current.String()))
	return args
}

func initializeKeyMap() map[string]string {
	keyMap := make(map[string]string)

	keyMap["lshift"] = "lshift"
	keyMap["rshift"] = "rshift"
	keyMap["lctrl"] = "lctrl"
	keyMap["rctrl"] = "rctrl"
	keyMap["lalt"] = "lalt"
	keyMap["ralt"] = "ralt"
	keyMap["space"] = "space"
	keyMap["enter"] = "enter"
	keyMap["backspace"] = "backspace"
	keyMap["tab"] = "tab"
	keyMap["esc"] = "esc"
	keyMap["escape"] = "esc"
	keyMap["delete"] = "delete"
	keyMap["insert"] = "insert"
	keyMap["home"] = "home"
	keyMap["end"] = "end"
	keyMap["pageup"] = "pageup"
	keyMap["pagedown"] = "pagedown"
	keyMap["up"] = "up"
	keyMap["down"] = "down"
	keyMap["left"] = "left"
	keyMap["right"] = "right"

	// Function keys
	for i := 1; i <= 24; i++ {
		keyMap["f"+strconv.Itoa(i)] = "f" + strconv.Itoa(i)
	}

	keyMap["cmd"] = "cmd"
	keyMap["lcmd"] = "lcmd"
	keyMap["rcmd"] = "rcmd"
	keyMap["alt"] = "alt"
	keyMap["ctrl"] = "ctrl"
	keyMap["control"] = "control"
	keyMap["shift"] = "shift"
	keyMap["capslock"] = "capslock"
	keyMap["print"] = "print"
	keyMap["printscreen"] = "printscreen"
	keyMap["menu"] = "menu"
	keyMap["audio_mute"] = "audio_mute"
	keyMap["audio_vol_down"] = "audio_vol_down"
	keyMap["audio_vol_up"] = "audio_vol_up"
	keyMap["audio_play"] = "audio_play"
	keyMap["audio_stop"] = "audio_stop"
	keyMap["audio_pause"] = "audio_pause"
	keyMap["audio_prev"] = "audio_prev"
	keyMap["audio_next"] = "audio_next"
	keyMap["audio_rewind"] = "audio_rewind"
	keyMap["audio_forward"] = "audio_forward"
	keyMap["audio_repeat"] = "audio_repeat"
	keyMap["audio_random"] = "audio_random"

	// Number keys
	for i := 0; i <= 9; i++ {
		keyMap[strconv.Itoa(i)] = strconv.Itoa(i)
	}

	// Numpad keys
	keyMap["num0"] = "num0"
	keyMap["num1"] = "num1"
	keyMap["num2"] = "num2"
	keyMap["num3"] = "num3"
	keyMap["num4"] = "num4"
	keyMap["num5"] = "num5"
	keyMap["num6"] = "num6"
	keyMap["num7"] = "num7"
	keyMap["num8"] = "num8"
	keyMap["num9"] = "num9"
	keyMap["num_lock"] = "num_lock"
	keyMap["num."] = "num."
	keyMap["num+"] = "num+"
	keyMap["num-"] = "num-"
	keyMap["num*"] = "num*"
	keyMap["num/"] = "num/"
	keyMap["num_clear"] = "num_clear"
	keyMap["num_enter"] = "num_enter"
	keyMap["num_equal"] = "num_equal"

	// Monitor and keyboard light control keys
	keyMap["lights_mon_up"] = "lights_mon_up"
	keyMap["lights_mon_down"] = "lights_mon_down"
	keyMap["lights_kbd_toggle"] = "lights_kbd_toggle"
	keyMap["lights_kbd_up"] = "lights_kbd_up"
	keyMap["lights_kbd_down"] = "lights_kbd_down"

	// Alphabet keys
	for c := 'a'; c <= 'z'; c++ {
		keyMap[string(c)] = string(c)
	}
	for c := 'A'; c <= 'Z'; c++ {
		keyMap[string(c)] = string(c)
	}

	// Symbols
	keyMap[";"] = "semicolon"
	keyMap["="] = "equals"
	keyMap[","] = "comma"
	keyMap["-"] = "minus"
	keyMap["."] = "period"
	keyMap["/"] = "slash"
	keyMap["\\"] = "backslash"
	keyMap["["] = "openbracket"
	keyMap["]"] = "closebracket"
	keyMap["'"] = "quote"

	return keyMap
}
