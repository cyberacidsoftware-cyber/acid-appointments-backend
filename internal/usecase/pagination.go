package usecase

type PaginationArgs struct {
	PageSize int32
	Page     int32
}

func (args *PaginationArgs) Validate() *UsecaseError {
	if args.PageSize == 0 {
		args.PageSize = 10
	}

	if args.PageSize < 0 {
		return NewInvalidArgumentError(ErrExpectPositiveValue).InField("pageSize")
	}
	if args.Page < 0 {
		return NewInvalidArgumentError(ErrExpectPositiveValue).InField("page")
	}
	return nil
}
