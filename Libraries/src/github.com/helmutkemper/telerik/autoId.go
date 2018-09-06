package telerik

import (
	"math/rand"
	"time"
)

var autoIdBase []rune
var autoIdLength int

func GetAutoId() string {
	b := make([]rune, autoIdLength)
	for i := range b {
		b[i] = autoIdBase[rand.Intn(len(autoIdBase))]
	}
	return "autoId_" + string(b)
}

func init() {
	rand.Seed(time.Now().UnixNano())

	autoIdLength = 32
	autoIdBase = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ01234567890")
}
