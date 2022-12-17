package utils

import (
	"strconv"
	"strings"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func MustInt(s string) int {
	i, err := strconv.Atoi(strings.TrimSpace(s))
	CheckError(err)
	return i
}

func StrSliceToInt(ss []string) []int {
	r := make([]int, len(ss))
	for idx, s := range ss {
		r[idx] = MustInt(s)
	}
	return r
}
