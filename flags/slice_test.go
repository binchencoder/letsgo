package flags

import (
	"flag"
	"testing"
)

func TestStringSlice(t *testing.T) {
	ss := StringSlice{}

	fs := flag.NewFlagSet("test", flag.ContinueOnError)
	fs.Var(&ss, "item", "The item flag to put in slice")

	args := []string{
		"--item=apple",
		"--item=banana",
		"--item=hammer",
	}
	if err := fs.Parse(args); err != nil {
		t.Fatal(err)
	}

	if len(ss) != 3 {
		t.Fatalf("Expect %d items but got %d", 3, len(ss))
	}
	if ss[0] != "apple" {
		t.Errorf("Expect 1st item as apple but got %s", ss[0])
	}
	if ss[1] != "banana" {
		t.Errorf("Expect 2nd item as banana but got %s", ss[1])
	}
	if ss[2] != "hammer" {
		t.Errorf("Expect 3rd item as hammer but got %s", ss[2])
	}
}
