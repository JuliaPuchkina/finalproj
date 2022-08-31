package censor

import "strings"

// определение наличия "нехороших слов" в комментарии
// сорри за мат)))
func IsCommentBad(text string) bool {
	badWords := []string{"блядь", "пидор"}

	var isBad bool

	for i := 0; i < len(badWords); i++ {
		isBad = strings.ContainsAny(text, badWords[i])
	}

	return isBad
}
