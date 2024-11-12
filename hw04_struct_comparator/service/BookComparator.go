package service

import "github.com/Konstanta100/home_work_basic/hw04_struct_comparator/entity"

type Regime int

const (
	YEAR Regime = iota
	SIZE
	RATE
)

type BookComparator struct {
	regime Regime
}

func NewBookComparator(regime Regime) *BookComparator {
	return &BookComparator{regime: regime}
}

func (b *BookComparator) Compare(book1, book2 entity.Book) bool {
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
