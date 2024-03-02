package obj

type PageableDtoWrapper struct {
	Limit int           `json:"limit"`
	Page  int           `json:"page"`
	Body  []interface{} `json:"body"`
}

func (pdw PageableDtoWrapper) New(pageable Pageable, body []any) PageableDtoWrapper {
	if body == nil {
		body = []interface{}{}
	}
	return PageableDtoWrapper{Limit: pageable.GetLimit(), Page: pageable.GetPage(), Body: body}
}
