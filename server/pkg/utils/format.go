package utils

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func FormatText(text string) string {
	caser := cases.Title(language.English)

	firstChar := text[:1]
	restChars := text[1:]

	return caser.String(firstChar) + restChars
}