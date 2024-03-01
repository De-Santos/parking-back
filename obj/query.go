package obj

type ParkingSearchQuery struct {
	Limit      int    `form:"limit" validate:"required,gt=0"`
	Page       int    `form:"page" validate:"gte=0"`
	SearchText string `form:"search_text"`
}

func (psq *ParkingSearchQuery) GetLimit() int {
	return psq.Limit
}

func (psq *ParkingSearchQuery) GetPage() int {
	return psq.Page
}
