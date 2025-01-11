package book

type Regime int

const (
	YEAR Regime = iota
	SIZE
	RATE
)

type Comparator struct {
	regime Regime
}

func NewComparator(regime Regime) *Comparator {
	return &Comparator{regime: regime}
}

func (b *Comparator) Compare(book1, book2 Book) bool {
	switch b.regime {
	case YEAR:
		return book1.Year() > book2.Year()
	case SIZE:
		return book1.Size() > book2.Size()
	case RATE:
		return book1.Rate() > book2.Rate()
	default:
		return false
	}
}
