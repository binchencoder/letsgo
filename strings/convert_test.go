package strings

import (
	"testing"
)

func TestCsvToSlice(t *testing.T) {
	csv := "a, ,b, c"
	slice := CsvToSlice(csv)
	if len(slice) != 3 || slice[0] != "a" || slice[1] != "b" || slice[2] != "c" {
		t.Errorf("expect result [a b c] but got %v", slice)
	}

	csv = "http://192.168.0.1:2379, ,http://192.168.0.2:2379, http://192.168.0.3:2379"
	slice = CsvToSlice(csv)
	if len(slice) != 3 || slice[0] != "http://192.168.0.1:2379" || slice[1] != "http://192.168.0.2:2379" || slice[2] != "http://192.168.0.3:2379" {
		t.Errorf("expect result [http://192.168.0.1:2379 http://192.168.0.2:2379 http://192.168.0.3:2379] but got %v", slice)
	}

	// Space, double quotes, escape.
	csvs := map[string][]string{
		`a,"b,c"`:      {"a", "b,c"},
		` a, " b,c " `: {"a", "b,c"},
		`a,\"b\"`:      {"a", `"b"`},
	}

	for s, expected := range csvs {
		slice = CsvToSlice(s)
		if !strSliceEquals(slice, expected) {
			t.Errorf("%s: expect result %v but got %v", s, expected, slice)
		}
	}
}

func strSliceEquals(slice1, slice2 []string) bool {
	if len(slice1) != len(slice1) {
		return false
	}
	for i := 0; i < len(slice1); i++ {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}
