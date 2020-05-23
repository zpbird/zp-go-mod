// Package zstrings import
package zstrings

import (
	"unicode"
)

// IncludeZhCn 判断字符串中是否包含汉字...
func IncludeZhCn(str string) bool {
	var count int
	for _, v := range str {
		if unicode.Is(unicode.Han, v) {
			count++
			break
		}
	}
	return count > 0
}
