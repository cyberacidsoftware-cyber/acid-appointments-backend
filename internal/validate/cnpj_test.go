package validate

// https://github.com/carvalhoviniciusluiz/cpf-cnpj-validator
import (
	"testing"
)

func TestIsCnpjValid(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		// invalid length
		{input: "", expected: false},
		{input: "0", expected: false},
		{input: "12", expected: false},
		{input: "123", expected: false},
		{input: "1234", expected: false},
		{input: "12345", expected: false},
		{input: "123456", expected: false},
		{input: "1234567", expected: false},
		{input: "12345678", expected: false},
		{input: "123456789", expected: false},
		{input: "1234567890", expected: false},
		{input: "12345678901", expected: false},
		{input: "123456789012", expected: false},
		{input: "1234567890123", expected: false},
		{input: "123456789012345", expected: false},
		// blacklist
		{input: "00000000000000", expected: false},
		{input: "11111111111111", expected: false},
		{input: "22222222222222", expected: false},
		{input: "33333333333333", expected: false},
		{input: "44444444444444", expected: false},
		{input: "55555555555555", expected: false},
		{input: "66666666666666", expected: false},
		{input: "77777777777777", expected: false},
		{input: "88888888888888", expected: false},
		{input: "99999999999999", expected: false},
		// format is not allowed
		{input: "54.550.752/0001-55", expected: false},
		// valid
		{input: "54550752000155", expected: true},
	}

	for _, test := range tests {
		result := IsCnpjValid(test.input)
		if result != test.expected {
			t.Errorf("IsCnpjValid(%v) = %v; want %v", test.input, result, test.expected)
		}
	}
}
