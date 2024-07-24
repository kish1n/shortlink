package data

type LinksQ interface {
	FilterByOriginal(original string) (CoupleLinks, error)
	FilterByShortened(shortened string) (CoupleLinks, error)
	Insert(value CoupleLinks) (*CoupleLinks, error)
}

type CoupleLinks struct {
	Shortened string `db:"shortened" structs:"shortened"`
	Original  string `db:"original" structs:"original"`
}

//TODO replace builder with a struct
