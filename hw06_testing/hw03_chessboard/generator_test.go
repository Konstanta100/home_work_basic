package chessboard

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateBoard_MainTableTest(t *testing.T) {
	testCases := []struct {
		name      string
		width     int
		height    int
		result    string
		exception string
	}{
		{
			name:      "MainCase",
			width:     5,
			height:    3,
			result:    " # # \n# # #\n # # \n",
			exception: "",
		},
		{
			name:      "SmallHeightCase",
			width:     8,
			height:    1,
			result:    " # # # #\n",
			exception: "",
		},
		{
			name:      "SmallWidthCase",
			width:     1,
			height:    8,
			result:    " \n#\n \n#\n \n#\n \n#\n",
			exception: "",
		},
		{
			name:      "EmptyWidthCase",
			width:     0,
			height:    1,
			result:    "",
			exception: "ширина и высота должны быть больше 0",
		},
		{
			name:      "EmptyHeightCase",
			width:     1,
			height:    0,
			result:    "",
			exception: "ширина и высота должны быть больше 0",
		},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			result, err := generateBoard(tC.width, tC.height)

			assert.Equal(t, tC.result, result)

			if err != nil {
				assert.Equal(t, tC.exception, err.Error())
			}
		})
	}
}
