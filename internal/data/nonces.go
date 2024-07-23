package data

type LinksQ interface {
	Get() (*CoupleLinks, error)
	Select() ([]CoupleLinks, error)
	Insert(value CoupleLinks) (*CoupleLinks, error)
	Update(value CoupleLinks) (*CoupleLinks, error)
	Delete() error

	FilterByAddress(addresses ...string) LinksQ
	FilterExpired() LinksQ
}

type CoupleLinks struct {
	Shortened string `db:"shortened" structs:"shortened"`
	Original  string `db:"original" structs:"original"`
}
