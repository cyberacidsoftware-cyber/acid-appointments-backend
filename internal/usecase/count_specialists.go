package usecase

type CountSpecialistArgs struct {
	Cpf   string
	Phone string
	Name  string
}

func (args *CountSpecialistArgs) Validate() *UsecaseError {
	return nil
}

func CountSpecialists(state State, args CountSpecialistArgs) (int64, *UsecaseError) {
	count, err := state.Queries().CountSpecialists(state.Context())
	if err == nil {
		return count, nil
	}
	return 0, NewUnexpectedError(err)
}
