package obj

type Pageable interface {
	GetLimit() int
	GetPage() int
}
