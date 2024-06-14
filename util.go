package main

import (
	"strconv"
	"strings"
)

func colorsMatch(c1, c2 string, threshold int) bool {
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
