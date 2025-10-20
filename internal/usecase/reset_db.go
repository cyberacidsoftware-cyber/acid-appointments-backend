package usecase

func ResetDb(state State) *UsecaseError {
	err := state.Queries().ResetDb(state.Context())
	if err != nil {
		return NewUnexpectedError(err)
	}
	return nil
}
