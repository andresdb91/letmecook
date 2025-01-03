package commons

type PagedList[T any] struct {
	// TotalItems  int
	// LastPage    int
	PageSize    int
	CurrentPage int
	HasNext     bool
	Items       []*T
	GetNextPage func() (*PagedList[T], error)
}

type QueryResult[T any] struct {
	Items       []*T
	Next        func() (*QueryResult[T], error)
	CurrentPage int
	PageSize    int
}

// Recibe un `next` y un parser; devuelve un `PagedList[T]` parseado para business
func ParseQueryPage[F any, T any](query func() (*QueryResult[F], error), parser func(dbPage *[]*F) ([]*T, error)) (*PagedList[T], error) {
	page, queryErr := query()
	if queryErr != nil {
		return nil, queryErr
	}
	items, parserErr := parser(&page.Items)
	if parserErr != nil {
		return nil, parserErr
	}
	pagedResult := PagedList[T]{
		Items:       items,
		PageSize:    page.PageSize,
		CurrentPage: page.CurrentPage,
		HasNext:     page.Next != nil,
		GetNextPage: func() (*PagedList[T], error) {
			return ParseQueryPage[F, T](page.Next, parser)
		},
	}
	return &pagedResult, nil
}
