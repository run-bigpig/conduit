package utils

import "math/rand"

// RandomString 随机生成指定长度的包含数字小写字母的字符串
func RandomString(length int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyz0123456789")
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
