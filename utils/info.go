package utils

import "strings"

func GetDetailInfo(txt string) string {
	result := strings.Split(txt, ": ")

	return result[1]
}
