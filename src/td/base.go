package td

import "strings"

func Split(str, sep string) (result []string) {

	// 空字符串直接返回
	if str == "" {
		return []string{str}
	}
	// 空分隔符直接返回
	if sep == "" {
		return []string{str}
	}

	result = strings.Split(str, sep)
	return result
}
