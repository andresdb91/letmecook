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
