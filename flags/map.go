package flags

import (
	"fmt"
	"strings"
)

// Type StringMap implements flag.Value and allows defining a flag which can be
// provided in command line multiple times, each time a "key=value" string can
// be specified. The key/value pairs will be collected in the map. Example:
//
// var mymap = flags.StringMap{}
// flag.Var(&mymap, "item", "The item flag to put in map")
// flag.Parse()
// fmt.Println(mymap)
//
// The above code snippet introduces an --item flag which can be used in command
// line:
//     myprogram --item=apple=fruit --item=rabbit=animal --item=hammer=tool
// The output would be: map[apple:fruit rabbit:animal hammer:tool]
//
type StringMap map[string]string

// String returns the contents of the map as a string.
func (s *StringMap) String() string {
	return fmt.Sprintf("%q", *s)
}

// Set parses the input value (in format key=value) and put the key and value
// into the map.
func (s *StringMap) Set(value string) error {
	parts := strings.SplitN(value, "=", 2)
	if len(parts) != 2 {
		return fmt.Errorf("expect key=value")
	}

	(*s)[parts[0]] = parts[1]
	return nil
}
