package generator

import (
	"fmt"
	"math/rand"
)

func NewGenerator(countryCode string, bankCode string) *Generator {
	return &Generator{countryCode: countryCode, bankCode: bankCode}
}

type Generator struct {
	countryCode string
	bankCode    string
}

func (g Generator) GenerateRandomIban() string {
	return fmt.Sprintf("%s%s%s00000%s", g.countryCode, randStringRunes(2), g.bankCode, randStringRunes(14))
}

var numberRunes = []rune("1234567890")

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = numberRunes[rand.Intn(len(numberRunes))]
	}
	return string(b)
}
