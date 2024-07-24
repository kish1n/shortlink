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

type CoupleLinksBuilder struct {
	coupleLinks CoupleLinks
}

func NewCoupleLinksBuilder() *CoupleLinksBuilder {
	return &CoupleLinksBuilder{
		coupleLinks: CoupleLinks{},
	}
}

func (b *CoupleLinksBuilder) SetShortened(shortened string) *CoupleLinksBuilder {
	b.coupleLinks.Shortened = shortened
	return b
}

func (b *CoupleLinksBuilder) SetOriginal(original string) *CoupleLinksBuilder {
	b.coupleLinks.Original = original
	return b
}

func (b *CoupleLinksBuilder) Build() CoupleLinks {
	return b.coupleLinks
}
