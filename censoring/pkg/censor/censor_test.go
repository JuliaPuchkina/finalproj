package censor

import "testing"

func TestIsCommentBad(t *testing.T) {
	badComment := "ты блядь"

	result := IsCommentBad(badComment)

	if result == false {
		t.Fatal("тест не пройден")
	}

}
