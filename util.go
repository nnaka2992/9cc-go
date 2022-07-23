package main

import (
	"strconv"
	"strings"
)

func StrToInt(s string) (v int, offset int) {
	offset = strings.IndexFunc(s, IsDigit)
	if offset == -1 {
		offset = len(s)
	}
	if offset == 0 {
		return
	} // Avoid Atoi on empty string
	v, _ = strconv.Atoi(s[:offset])
	return
}

func RuneToInt(rs []rune) (v int, offset int) {
	offset = 0
	for _, r := range rs {
		if IsDigit(r) {
			offset += 1
			continue
		}
		break
	}
	if offset == 0 {
		offset = len(rs)
	}
	v, _ = strconv.Atoi(string(rs[:offset]))
	return v, offset
}

func IsDigit(r rune) bool {
	return '0' <= r && r <= '9'
}
