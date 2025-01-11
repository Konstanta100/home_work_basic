package types

import (
	"testing"
)

func TestEmployeeString(t *testing.T) {
	employee := Employee{
		UserID:       1,
		Age:          18,
		Name:         "Kostya",
		DepartmentID: 456,
	}

	expected := "User ID: 1; Age: 18; Name: Kostya; Department ID: 456;"
	result := employee.String()

	if result != expected {
		t.Errorf("Ожидается %q, но получили %q", expected, result)
	}
}
