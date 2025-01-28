package practice

import "strings"

func AddNewLineCharacter(s string) string {
	return strings.ReplaceAll(s, "\n", "\\n")
}
