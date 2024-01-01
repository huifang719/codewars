package deadfishswim

import (
	"reflect"
	"testing"
)

type testpair struct {
	values string
	result []int
}
func TestDeadfishSwim(t *testing.T) {
	testData := []testpair{
		{"iiisdoso", []int{8, 64}},
		{"iiisdosodddddiso", []int{8, 64, 3600}},
	}
	for _, pair := range testData {
		v := DeadfishSwim(pair.values)
		if !reflect.DeepEqual(v, pair.result) {
			t.Error(
				"For", pair.values,
				"expected", pair.result,
				"got", v,
			)
		}
	}
}