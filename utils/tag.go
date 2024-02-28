package utils

import "strings"

func GetTagList(tag string) []string {
	return strings.Split(tag, ";")
}
