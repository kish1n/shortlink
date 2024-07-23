package data

type LinksQ interface {
	Get() (*CoupleLinks, error)
	Select() ([]CoupleLinks, error)
	Insert(value CoupleLinks) (*CoupleLinks, error)
	Update(value CoupleLinks) (*CoupleLinks, error)
	Delete() error
	FilterByShortened(addresses ...string) LinksQ
	FilterByOriginal(addresses ...string) LinksQ
}

type CoupleLinks struct {
	Shortened string `db:"shortened" structs:"shortened"`
	Original  string `db:"original" structs:"original"`
}
