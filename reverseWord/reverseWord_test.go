package reverseWord

import (
	"testing"
)

type testpair struct {
	values string
	result string
}




func TestReverseWord(t *testing.T) {
	testData := []testpair{
		{"This is an example!", "sihT si na !elpmaxe"},
		{"double  spaces", "elbuod  secaps"},
	}
	for _, pair := range testData {
		v := ReverseWord(pair.values)
		if v != pair.result {
			t.Error(
				"For", pair.values,
				"expected", pair.result,
				"got", v,
			)
		}
	}
}