// Package translitger provides a function to transliterate `string`s with German
// umlauts to only use ASCII characters.
package translitger

import "strings"

var mapping = map[string]string{
	"ä": "ae",
	"Ä": "Ae",
	"ü": "ue",
	"Ü": "Ue",
	"ö": "oe",
	"Ö": "Oe",
	"ß": "ss",
}

// Translit transforms all German umlauts into ASCII charater sequences
func Translit(input string) string {
	for k, v := range mapping {
		input = strings.ReplaceAll(input, k, v)
	}
	return input
}
