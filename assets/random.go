package assets

import (
	"math/rand"
	"strings"
	"time"
)

// CreateRandomStr return random str with current lenght of symbols
func CreateRandomStr(length int) string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYabcdefghijklmnopqrstuvwxyz0123456789")
	var b strings.Builder

	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}

	str := b.String()
	return str
}
