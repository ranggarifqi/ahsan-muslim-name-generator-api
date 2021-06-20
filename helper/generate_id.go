package helper

import (
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func GenerateID(length int) string {

	result := make([]byte, length)

	for i := range result {
		result[i] = charset[seededRand.Intn(len(charset))]
	}

	return string(result)
}
