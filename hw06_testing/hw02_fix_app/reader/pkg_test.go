package reader

import (
	"testing"

	"github.com/Konstanta100/home_work_basic/hw06_testing/hw02_fix_app/types"
	"github.com/stretchr/testify/assert"
)

func TestReadJSON_MainTableTest(t *testing.T) {
	testCases := []struct {
		name      string
		path      string
		wantCount int
		wantData  []types.Employee
	}{
		{
			name:      "MainCase",
			path:      "data.json",
			wantCount: 2,
			wantData: []types.Employee{
				{UserID: 10, Age: 25, Name: "Rob", DepartmentID: 3},
				{UserID: 11, Age: 30, Name: "George", DepartmentID: 2},
			},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			staff, err := ReadJSON(tC.path)

			assert.Len(t, staff, tC.wantCount)
			assert.NoError(t, err)
			assert.Equal(t, tC.wantData, staff)
		})
	}
}

func TestReadJSON_FailedTableTest(t *testing.T) {
	testCases := []struct {
		name string
		path string
	}{
		{
			name: "EmptyPathCase",
			path: "",
		},
		{
			name: "EmptyFileCase",
			path: "empty_data.json",
		},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			staff, err := ReadJSON(tC.path)

			assert.Nil(t, staff)
			assert.Error(t, err)
		})
	}
}
