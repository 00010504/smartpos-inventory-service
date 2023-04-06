package helper

import (
	"math/rand"
	"time"
)

var letterRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
var numberRunes = []rune("0123456789")

func RandStringRunes(n int, format string) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	if format == "string" {
		for i := range b {
			b[i] = letterRunes[rand.Intn(len(letterRunes))]
		}
	} else if format == "number" {
		for i := range b {
			b[i] = numberRunes[rand.Intn(len(numberRunes))]
		}
	}
	return string(b)
}
