package main

import (
	"fmt"

	"github.com/Konstanta100/home_work_basic/hw04_struct_comparator/entity"
	"github.com/Konstanta100/home_work_basic/hw04_struct_comparator/service"
)

func main() {
	book1 := entity.Book{}
	book1.SetID(1)
	book1.SetTitle("test1")
	book1.SetAuthor("Тестер1")
	book1.SetYear(2000)
	book1.SetSize(100)
	book1.SetRate(4.2)

	book2 := entity.Book{}
	book2.SetID(2)
	book2.SetTitle("test2")
	book2.SetAuthor("Тестер2")
	book2.SetYear(1998)
	book2.SetSize(200)
	book2.SetRate(4.2)

	comparator := service.NewBookComparator(service.YEAR)
	fmt.Println(comparator.Compare(book1, book2))
	comparator = service.NewBookComparator(service.SIZE)
	fmt.Println(comparator.Compare(book1, book2))
	comparator = service.NewBookComparator(service.RATE)
	fmt.Println(comparator.Compare(book1, book2))
}
