package rowsumoddnumbers

import (
	"testing"
)

type testpair struct {
	values int
	result int
}

func TestRowSumOddNumbers (t *testing.T) {
	testData := []testpair{
		{1, 1},
		{2, 8},
		{3, 27},
	}
	for _, pair := range testData {
		v := RowSumOddNumbers(pair.values)
		if v != pair.result {
			t.Error(
				"For", pair.values,
				"expected", pair.result,
				"got", v,
			)
		}
	}
}