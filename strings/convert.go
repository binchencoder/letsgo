package strings

import (
	"encoding/csv"
	"log"
	"strings"
)

// CsvToSlice converts a comma separated string to a string slice.
// Return empty string if the string is a malformed CSV record.
// Leading/tailing spaces are trimmed for each field.
//
// Limitations:
// 1. The string must be single line.
// 2. If there is comma in field value, the field value must be double quoted.
// 3. "a" -> a,  \"a\" -> "a"
func CsvToSlice(in string) []string {
	in = strings.TrimSpace(in)

	r := csv.NewReader(strings.NewReader(in))

	// Set TrimLeadingSpace to true to avoid error: `bare " in non-quoted-field`.
	// This error happens when there are one or more space(s) after field
	// separator (comma), for example: `a, "b,c"`
	r.TrimLeadingSpace = true

	// Read a csv record.
	record, err := r.Read()
	if err != nil {
		// Panic? since most usages are during app start.
		log.Printf("letsgo/strings.CsvToSlice(%s) failed: %v", in, err)
		return []string{}
	}

	slice := []string{}
	for _, v := range record {
		v = strings.TrimSpace(v)
		if v != "" {
			slice = append(slice, v)
		}
	}

	return slice
}
