package obj

type Pageable interface {
	GetLimit() int
	GetPage() int
}

type ModelUpdated interface {
	GetUpdatedColumns() map[string]interface{}
}
