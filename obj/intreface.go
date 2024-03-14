package obj

type Pageable interface {
	GetLimit() int
	GetPage() int
}

type ModelUpdated interface {
	GetUpdatedColumns() map[string]interface{}
}

type Pagination interface {
	Pageable
	SetLimit(limit int)
	SetPage(page int)
	SetTotalRows(rows int64)
	SetTotalPage(pages int)
}

type Search interface {
	GetSearchBy() string
	GetSearchText() string
	GetType() string
}
