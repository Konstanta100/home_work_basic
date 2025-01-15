package utilword

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountWords(t *testing.T) {
	testCases := []struct {
		name     string
		text     string
		wantLen  int
		wantData map[string]int
	}{
		{
			name:    "MainCase",
			text:    "Пришёл, увидел, победил! А если не пришёл, то проиграл ? Надо крест-накрест перекреститься",
			wantLen: 12,
			wantData: map[string]int{
				"а":              1,
				"пришёл":         2,
				"увидел":         1,
				"победил":        1,
				"если":           1,
				"не":             1,
				"то":             1,
				"проиграл":       1,
				"надо":           1,
				"крест":          1,
				"перекреститься": 1,
				"накрест":        1,
			},
		},
		{
			name:    "MainCase",
			text:    "Тестить-Тестить очень ////сложно-очень тестить    ....тестить!!!СлОжно очень",
			wantLen: 3,
			wantData: map[string]int{
				"тестить": 4,
				"очень":   3,
				"сложно":  2,
			},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			result := CountWords(tC.text)

			assert.Len(t, result, tC.wantLen)
			assert.Equal(t, tC.wantData, result)
		})
	}
}
