package complementaryDNA

import (
	"testing"
)

type testpair struct {
	values string
	result string
}


func TestComplementaryDNA(t *testing.T) {
	testData := []testpair{
		{"ATTGC", "TAACG"},
		{"GTAT", "CATA"},
	}
	for _, pair := range testData {
		v := ComplementaryDNA(pair.values)
		if v != pair.result {
			t.Error(
				"For", pair.values,
				"expected", pair.result,
				"got", v,
			)
		}
	}
}