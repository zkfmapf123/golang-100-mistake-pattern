package main

import "strings"

func _25_trim() {
	s := "oxo123123oxo"

	strings.TrimLeft(s, "ox")   // 123123oxo -> left에 o, x 모두 제거
	strings.TrimRight(s, "ox")  // oxo123123 -> right에 o, x 모두 제거
	strings.TrimSuffix(s, "ox") // oxo123123 -> suffix에 ox 제거
	strings.Trim(s, "ox")       // 123123 -> left, right에 o, x 모두 제거 (TrimLeft + TrimRight)
}
