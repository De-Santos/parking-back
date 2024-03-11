package obj

type PageableWrapper struct {
	Limit      int           `json:"limit"`
	Page       int           `json:"page"`
	TotalRows  int64         `json:"total_rows"`
	TotalPages int           `json:"total_pages"`
	Body       []interface{} `json:"body"`
}

func (w *PageableWrapper) GetLimit() int {
	return w.Limit
}

func (w *PageableWrapper) GetPage() int {
	return w.Page
}

func (w *PageableWrapper) SetLimit(limit int) {
	w.Limit = limit
}

func (w *PageableWrapper) SetPage(page int) {
	w.Page = page
}

func (w *PageableWrapper) SetTotalRows(rows int64) {
	w.TotalRows = rows
}

func (w *PageableWrapper) SetTotalPage(pages int) {
	w.TotalPages = pages
}

func (w *PageableWrapper) SetBody(body []any) {
	w.Body = body
}

func (w *PageableWrapper) OffMigrate(pageable Pageable) {
	w.Limit = pageable.GetLimit()
	w.Page = pageable.GetPage()
}
