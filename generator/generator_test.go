package generator

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerator_GenerateRandomIban(t *testing.T) {
	type fields struct {
		countryCode string
		bankCode    string
	}
	tests := []struct {
		name              string
		fields            fields
		wantCountryCode   string
		wantBankCode      string
		wantAccountLength int
	}{
		{
			name: "UA iban",
			fields: fields{
				countryCode: "UA",
				bankCode:    "123456",
			},
			wantCountryCode:   "UA",
			wantBankCode:      "123456",
			wantAccountLength: 14,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGenerator(tt.fields.countryCode, tt.fields.bankCode)
			got := g.GenerateRandomIban()
			assert.Equal(t, tt.wantCountryCode, got[:2])
			assert.Equal(t, tt.wantBankCode, got[4:10])
			assert.Equal(t, tt.wantAccountLength, len(got[15:]))

			fmt.Println(got[15:])
		})
	}
}
