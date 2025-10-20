package validate

var BLACKLIST_CNPJ []string = []string{
	"00000000000000",
	"11111111111111",
	"22222222222222",
	"33333333333333",
	"44444444444444",
	"55555555555555",
	"66666666666666",
	"77777777777777",
	"88888888888888",
	"99999999999999",
}

func IsCnpjValid(cnpj string) bool {
	if len(cnpj) != 14 {
		return false
	}
	for _, ch := range cnpj {
		if !('0' <= ch && ch <= '9') {
			return false
		}
	}
	for _, invalid := range BLACKLIST_CNPJ {
		if cnpj == invalid {
			return false
		}
	}
	a0 := int(cnpj[0] - '0')
	a1 := int(cnpj[1] - '0')
	a2 := int(cnpj[2] - '0')
	a3 := int(cnpj[3] - '0')
	a4 := int(cnpj[4] - '0')
	a5 := int(cnpj[5] - '0')
	a6 := int(cnpj[6] - '0')
	a7 := int(cnpj[7] - '0')
	a8 := int(cnpj[8] - '0')
	a9 := int(cnpj[9] - '0')
	a10 := int(cnpj[10] - '0')
	a11 := int(cnpj[11] - '0')
	a12 := int(cnpj[12] - '0')
	a13 := int(cnpj[13] - '0')
	sum12 := a0*5 + a1*4 + a2*3 + a3*2 + a4*9 + a5*8 + a6*7 + a7*6 + a8*5 + a9*4 + a10*3 + a11*2
	sum13 := a0*6 + a1*5 + a2*4 + a3*3 + a4*2 + a5*9 + a6*8 + a7*7 + a8*6 + a9*5 + a10*4 + a11*3 + a12*2
	return ((sum12*10)%11)%10 == a12 && ((sum13*10)%11)%10 == a13
}
