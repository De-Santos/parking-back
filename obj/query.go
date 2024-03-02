package obj

type SearchQuery struct {
	Limit      int    `form:"limit" validate:"required,gt=0"`
	Page       int    `form:"page" validate:"gte=0"`
	SearchText string `form:"search_text"`
	Context    int    `form:"c"`
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
