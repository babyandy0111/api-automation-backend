package helper

import (
	"strconv"
	"strings"
)

func Int64Unique(intSlice []int64) []int64 {
	keys := make(map[int64]bool)
	var list []int64
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func Int64Contains(intSlice []int64, matchInt int64) bool {
	for _, each := range intSlice {
		if each == matchInt {
			return true
		}
	}
	return false
}

func StringContains(strSlice []string, matchStr string) bool {
	for _, each := range strSlice {
		if each == matchStr {
			return true
		}
	}
	return false
}

func Int64Diff(intSlice []int64, matchIntSlice ...[]int64) (data []int64) {
	if len(matchIntSlice) == 0 {
		return intSlice
	}

	i := 0
loop:
	for {
		if i == len(intSlice) {
			break
		}
		v := intSlice[i]
		for _, arr := range matchIntSlice {
			for _, val := range arr {
				if v == val {
					i++
					continue loop
				}
			}
		}
		data = append(data, v)
		i++
	}
	return data
}

// 將 int array join 轉換為 string
func IntJoin(intSlice []int, separator string) string {
	data := make([]string, len(intSlice))

	for i, val := range intSlice {
		data[i] = strconv.Itoa(val)
	}

	return strings.Join(data, separator)
}

func StringUnique(strSlice []string) []string {
	keys := make(map[string]bool)
	var list []string
	for _, entry := range strSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func ReverseSliceString(ss []string) []string {
	n := len(ss)
	runes := make([]string, n)
	for _, rune := range ss {
		n--
		runes[n] = rune
	}

	return runes
}
