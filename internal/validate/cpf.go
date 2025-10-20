package validate

// https://dicasdeprogramacao.com.br/algoritmo-para-validar-cpf/
var BLACKLIST_CPF []string = []string{
	"00000000000",
	"11111111111",
	"22222222222",
	"33333333333",
	"44444444444",
	"55555555555",
	"66666666666",
	"77777777777",
	"88888888888",
	"99999999999",
	"12345678909",
}

func IsCpfValid(cpf string) bool {
	if len(cpf) != 11 {
		return false
	}
	for _, ch := range cpf {
		if !('0' <= ch && ch <= '9') {
			return false
		}
	}
	for _, invalid := range BLACKLIST_CPF {
		if cpf == invalid {
			return false
		}
	}
	a0 := int(cpf[0] - '0')
	a1 := int(cpf[1] - '0')
	a2 := int(cpf[2] - '0')
	a3 := int(cpf[3] - '0')
	a4 := int(cpf[4] - '0')
	a5 := int(cpf[5] - '0')
	a6 := int(cpf[6] - '0')
	a7 := int(cpf[7] - '0')
	a8 := int(cpf[8] - '0')
	a9 := int(cpf[9] - '0')
	a10 := int(cpf[10] - '0')
	sum9 := a0*10 + a1*9 + a2*8 + a3*7 + a4*6 + a5*5 + a6*4 + a7*3 + a8*2
	sum10 := a0*11 + a1*10 + a2*9 + a3*8 + a4*7 + a5*6 + a6*5 + a7*4 + a8*3 + a9*2
	return ((sum9*10)%11)%10 == a9 && ((sum10*10)%11)%10 == a10
}
