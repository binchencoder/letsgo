package flags

import (
	"fmt"
)

// Type StringSlice implements flag.Value and allows defining a flag which can be
// provided in command line multiple times, each time a string can be specified.
// The strings will be collected in the slice.
//
// var myslice = flags.StringSlice{}
// flag.Var(&myslice, "item", "The item flag to put in slice")
// flag.Parse()
// fmt.Println(myslice)
//
// The above code snippet introduces an --item flag which can be used in command
// line:
//     myprogram --item=apple --item=rabbit --item=hammer
// The output would be: [apple rabbit hammer]
//
type StringSlice []string

// String returns the contents of the slice as a string.
func (s *StringSlice) String() string {
	return fmt.Sprintf("%s", *s)
}

// Set puts the input string into the slice.
func (s *StringSlice) Set(value string) error {
	*s = append(*s, value)
	return nil
}
