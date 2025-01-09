package book

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompare_MainTableTest(t *testing.T) {
	book1 := Book{}
	book1.SetID(1)
	book1.SetTitle("test1")
	book1.SetAuthor("Тестер1")
	book1.SetYear(2000)
	book1.SetSize(100)
	book1.SetRate(4.2)

	book2 := Book{}
	book2.SetID(2)
	book2.SetTitle("test2")
	book2.SetAuthor("Тестер2")
	book2.SetYear(1998)
	book2.SetSize(200)
	book2.SetRate(4.2)

	testCases := []struct {
		name         string
		comparator   *Comparator
		result       bool
		resultRevert bool
	}{
		{
			name:         "YearCase",
			comparator:   NewComparator(YEAR),
			result:       true,
			resultRevert: false,
		},
		{
			name:         "SizeCase",
			comparator:   NewComparator(SIZE),
			result:       false,
			resultRevert: true,
		},
		{
			name:         "RateCompareCase",
			comparator:   NewComparator(RATE),
			result:       false,
			resultRevert: false,
		},
		{
			name:         "EmptyRegimeCase",
			comparator:   &Comparator{},
			result:       true,
			resultRevert: false,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			result := tC.comparator.Compare(book1, book2)
			resultRevert := tC.comparator.Compare(book2, book1)

			assert.Equal(t, tC.result, result)
			assert.Equal(t, tC.resultRevert, resultRevert)
		})
	}
}
