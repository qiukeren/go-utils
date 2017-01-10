package utils

import (
	"math/rand"
	"time"
)

const lower = "abcdefghijklmnopqrstuvwxyz"
const upper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const digit = "0123456789"
const punct = "~!@#$%^&*()_+-="

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func RandString(size int, include string) string {
	all := include

	lenAll := len(all)

	buf := make([]byte, size)
	for i := 0; i < size; i++ {
		b := all[rand.Intn(lenAll)]
		buf[i] = b
	}
	return string(buf)
}
