package flags

import (
	"flag"
	"testing"
)

func TestStringMap(t *testing.T) {
	m := StringMap{}

	fs := flag.NewFlagSet("test", flag.ContinueOnError)
	fs.Var(&m, "item", "The item flag to put in map")

	args := []string{
		"--item=apple=fruit",
		"--item=rabbit=animal",
		"--item=hammer=tool",
	}
	if err := fs.Parse(args); err != nil {
		t.Fatal(err)
	}

	if len(m) != 3 {
		t.Fatalf("Expect %d items but got %d", 3, len(m))
	}
	if v, ok := m["apple"]; !ok || v != "fruit" {
		t.Fatalf("Expect value fruit for key apple but got %s", v)
	}
	if v, ok := m["rabbit"]; !ok || v != "animal" {
		t.Fatalf("Expect value animal for key rabbit but got %s", v)
	}
	if v, ok := m["hammer"]; !ok || v != "tool" {
		t.Fatalf("Expect value fruit for key hammer but got %s", v)
	}
}
