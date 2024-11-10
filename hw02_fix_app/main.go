package main

import (
	"fmt"

	"github.com/Konstanta100/home_work_basic/hw02_fix_app/printer"
	"github.com/Konstanta100/home_work_basic/hw02_fix_app/reader"
	"github.com/Konstanta100/home_work_basic/hw02_fix_app/types"
)

func main() {
	var path string

	fmt.Printf("Enter data file path: ")
	fmt.Scanln(&path)

	if len(path) == 0 {
		path = "data.json"
	}

	var staff []types.Employee
	var err error

	staff, err = reader.ReadJSON(path)

	if err != nil {
		fmt.Printf("Error: %v", err)
	} else {
		printer.PrintStaff(staff)
	}
}
