
go run ./cmd/cli/main.go create-admin --name Admin --email admin@staff.com --password 123mudar
go run ./cmd/cli/main.go create-secretary --name "Secretary" --email secretary@staff.com --password 123mudar --phone 55123456789 --birthdate "2010-10-10" --cpf "45678912398"
go run ./cmd/cli/main.go create-customer -name "Customer" -email customer@people.com -phone 55789456321 -birthdate "2000-01-01" -cpf "12345678954"
go run ./cmd/cli/main.go create-specialist -name "Specialist" -email specialist@people.com -phone 55456789123 -birthdate "2001-02-02" -cpf "78945612354" -cnpj "98765432198722"
go run ./cmd/cli/main.go create-specialization -name "Specialization" 
go run ./cmd/cli/main.go create-service -name "Service" -specialization "Specialization"
