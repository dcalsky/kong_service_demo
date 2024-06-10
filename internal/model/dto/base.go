package dto

type EmptyRequest struct {
}

type EmptyResponse struct {
}

type PagingOption struct {
	PageSize int64
	PageNum  int64
}

type PagingResult struct {
	PageSize int64
	PageNum  int64
	Total    int64
}

type SortByOption struct {
	Items []SortByOptionItem
}

type SortByOptionItem struct {
	Field string
	Desc  bool
}

func (s PagingOption) Skip() int64 {
	if s.PageNum <= 1 {
		return 0
	}
	return (s.PageNum - 1) * s.PageSize
}
