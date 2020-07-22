package assets

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// GenerateBundleName return random name for new bundle file
func GenerateBundleName(length int) string {
	return fmt.Sprintf("%s.bundle.js", сreateRandomStr(length))
}

// GenerateSelectorName return random name for selector for main tag
func GenerateSelectorName(length int) string {
	return fmt.Sprintf("%s", сreateRandomStr(length))
}

func сreateRandomStr(length int) string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYabcdefghijklmnopqrstuvwxyz0123456789")
	var b strings.Builder

	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}

	str := b.String()
	return str
}
