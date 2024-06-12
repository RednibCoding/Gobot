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

func initializeKeyMap() map[string]string {
	keyMap := make(map[string]string)

	keyMap["lshift"] = "shift"
	keyMap["rshift"] = "shift"
	keyMap["lctrl"] = "ctrl"
	keyMap["rctrl"] = "ctrl"
	keyMap["lalt"] = "alt"
	keyMap["ralt"] = "alt"
	keyMap["space"] = "space"
	keyMap["enter"] = "enter"
	keyMap["backspace"] = "backspace"
	keyMap["tab"] = "tab"
	keyMap["esc"] = "esc"
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
	keyMap["f1"] = "f1"
	keyMap["f2"] = "f2"
	keyMap["f3"] = "f3"
	keyMap["f4"] = "f4"
	keyMap["f5"] = "f5"
	keyMap["f6"] = "f6"
	keyMap["f7"] = "f7"
	keyMap["f8"] = "f8"
	keyMap["f9"] = "f9"
	keyMap["f10"] = "f10"
	keyMap["f11"] = "f11"
	keyMap["f12"] = "f12"
	keyMap["numlock"] = "numlock"
	keyMap["capslock"] = "capslock"
	keyMap["scrolllock"] = "scrolllock"
	keyMap["pause"] = "pause"
	keyMap["printscreen"] = "printscreen"
	keyMap["windows"] = "win"
	keyMap["lmouse"] = "left"
	keyMap["rmouse"] = "right"

	for c := 'a'; c <= 'z'; c++ {
		keyMap[string(c)] = string(c)
	}
	for i := 0; i <= 9; i++ {
		keyMap[strconv.Itoa(i)] = strconv.Itoa(i)
	}

	// Numpad keys
	for i := 0; i <= 9; i++ {
		keyMap["numpad"+strconv.Itoa(i)] = "numpad" + strconv.Itoa(i)
	}
	keyMap["numpadadd"] = "numpadadd"
	keyMap["numpadsub"] = "numpadsub"
	keyMap["numpadmul"] = "numpadmul"
	keyMap["numpaddiv"] = "numpaddiv"
	keyMap["numpaddecimal"] = "numpaddecimal"
	keyMap["numpadenter"] = "numpadenter"

	// Symbols
	keyMap["semicolon"] = ";"
	keyMap["equals"] = "="
	keyMap["comma"] = ","
	keyMap["minus"] = "-"
	keyMap["period"] = "."
	keyMap["slash"] = "/"
	keyMap["backslash"] = "\\"
	keyMap["openbracket"] = "["
	keyMap["closebracket"] = "]"
	keyMap["quote"] = "'"

	return keyMap
}
