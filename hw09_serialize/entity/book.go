package entity

import "fmt"

type Book struct {
	ID     int     `json:"id" xml:"id,attr" yaml:"id" bson:"id"`
	Title  string  `json:"title" xml:"title" yaml:"title" bson:"title"`
	Author string  `json:"author" xml:"author" yaml:"author" bson:"author"`
	Year   int     `json:"year,omitempty" xml:"year,omitempty" yaml:"year,omitempty" bson:"year,omitempty"`
	Size   int     `json:"size" xml:"size" yaml:"size" bson:"size"`
	Rate   float64 `json:"rate" xml:"rate" yaml:"rate" bson:"rate"`
	Sample []byte  `json:"sample" xml:"sample" yaml:"sample" bson:"sample"`
}

func (book Book) String() string {
	return fmt.Sprintf("ID: %d; Title: %s; Author: %s; Year: %d; Size: %d, Rate: %f, Sample: %s ",
		book.ID, book.Title, book.Author, book.Year, book.Size, book.Rate, book.Sample,
	)
}
