package obj

type SearchQuery struct {
	Limit      int    `form:"limit" validate:"required,gt=0"`
	Page       int    `form:"page" validate:"gte=0"`
	SearchText string `form:"search_text"`
	Context    int    `form:"c"`
	SearchBy   string `form:"sb"`
	Type       string `form:"t"`
}

func (psq *SearchQuery) GetType() string {
	return psq.Type
}

func (psq *SearchQuery) GetSearchBy() string {
	return psq.SearchBy
}

func (psq *SearchQuery) GetSearchText() string {
	return psq.SearchText
}

func (psq *SearchQuery) GetLimit() int {
	return psq.Limit
}

func (psq *SearchQuery) GetPage() int {
	return psq.Page
}

type IdQuery struct {
	ID int `form:"id" validate:"required,gt=0"`
}

type StringQuery struct {
	String string `form:"s" validate:"required"`
}
