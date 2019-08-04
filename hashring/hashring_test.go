package hashring

import (
	"bufio"
	"os"
	"runtime"
	"sort"
	"strconv"
	"testing"
	"testing/quick"
)

func checkNum(num, expected int, t *testing.T) {
	if num != expected {
		t.Errorf("got %d, expected %d", num, expected)
	}
}

func TestNew(t *testing.T) {
	x := New()
	if x == nil {
		t.Errorf("expected obj")
	}
	checkNum(x.NumberOfReplicas, 20, t)
}

func TestAdd(t *testing.T) {
	x := New()
	x.Add(&Member{Key: "123456"})
	checkNum(len(x.circle), 20, t)
	checkNum(len(x.sortedHashes), 20, t)
	if sort.IsSorted(x.sortedHashes) == false {
		t.Errorf("expected sorted hashes to be sorted")
	}
	x.Add(&Member{Key: "qwer"})
	checkNum(len(x.circle), 40, t)
	checkNum(len(x.sortedHashes), 40, t)
	if !sort.IsSorted(x.sortedHashes) {
		t.Errorf("expected sorted hashes to be sorted")
	}
}

func TestRemove(t *testing.T) {
	x := New()
	x.Add(&Member{Key: "abcdefg"})
	x.Remove("abcdefg")
	checkNum(len(x.circle), 0, t)
	checkNum(len(x.sortedHashes), 0, t)
}

func TestRemoveNonExisting(t *testing.T) {
	x := New()
	x.Add(&Member{Key: "abcdefg"})
	x.Remove("abcdefghijk")
	checkNum(len(x.circle), 20, t)
}

func TestGetEmpty(t *testing.T) {
	x := New()
	_, err := x.Get("asdfsadfsadf")
	if err == nil {
		t.Errorf("expected error")
	}
	if err != ErrEmptyCircle {
		t.Errorf("expected empty circle error")
	}
}

func TestGetSingle(t *testing.T) {
	x := New()
	x.Add(&Member{Key: "abcdefg"})
	f := func(s string) bool {
		y, err := x.Get(s)
		if err != nil {
			t.Logf("error: %q", err)
			return false
		}
		t.Logf("s = %q, y = %q", s, y)
		return y == "abcdefg"
	}
	if err := quick.Check(f, nil); err != nil {
		t.Fatal(err)
	}
}

type gtest struct {
	in  string
	out string
}

var gmtests = []gtest{
	{"ggg", "abcdefg"},
	{"hhh", "opqrstu"},
	{"iiiii", "hijklmn"},
}

func TestGetMultiple(t *testing.T) {
	x := New()
	x.Add(&Member{Key: "abcdefg"})
	x.Add(&Member{Key: "hijklmn"})
	x.Add(&Member{Key: "opqrstu"})
	for i, v := range gmtests {
		result, err := x.Get(v.in)
		if err != nil {
			t.Fatal(err)
		}
		if result != v.out {
			t.Errorf("%d. got %q, expected %q", i, result, v.out)
		}
	}
}

func TestGetMultipleQuick(t *testing.T) {
	x := New()
	x.Add(&Member{Key: "abcdefg"})
	x.Add(&Member{Key: "hijklmn"})
	x.Add(&Member{Key: "opqrstu"})
	f := func(s string) bool {
		y, err := x.Get(s)
		if err != nil {
			t.Logf("error: %q", err)
			return false
		}
		t.Logf("s = %q, y = %q", s, y)
		return y == "abcdefg" || y == "hijklmn" || y == "opqrstu"
	}
	if err := quick.Check(f, nil); err != nil {
		t.Fatal(err)
	}
}

var rtestsBefore = []gtest{
	{"ggg", "abcdefg"},
	{"hhh", "opqrstu"},
	{"iiiii", "hijklmn"},
}

var rtestsAfter = []gtest{
	{"ggg", "abcdefg"},
	{"hhh", "opqrstu"},
	{"iiiii", "opqrstu"},
}

func TestGetMultipleRemove(t *testing.T) {
	x := New()
	x.Add(&Member{Key: "abcdefg"})
	x.Add(&Member{Key: "hijklmn"})
	x.Add(&Member{Key: "opqrstu"})
	for i, v := range rtestsBefore {
		result, err := x.Get(v.in)
		if err != nil {
			t.Fatal(err)
		}
		if result != v.out {
			t.Errorf("%d. got %q, expected %q before rm", i, result, v.out)
		}
	}
	x.Remove("hijklmn")
	for i, v := range rtestsAfter {
		result, err := x.Get(v.in)
		if err != nil {
			t.Fatal(err)
		}
		if result != v.out {
			t.Errorf("%d. got %q, expected %q after rm", i, result, v.out)
		}
	}
}

func TestGetMultipleRemoveQuick(t *testing.T) {
	x := New()
	x.Add(&Member{Key: "abcdefg"})
	x.Add(&Member{Key: "hijklmn"})
	x.Add(&Member{Key: "opqrstu"})
	x.Remove("opqrstu")
	f := func(s string) bool {
		y, err := x.Get(s)
		if err != nil {
			t.Logf("error: %q", err)
			return false
		}
		t.Logf("s = %q, y = %q", s, y)
		return y == "abcdefg" || y == "hijklmn"
	}
	if err := quick.Check(f, nil); err != nil {
		t.Fatal(err)
	}
}

func TestGetTwo(t *testing.T) {
	x := New()
	x.Add(&Member{Key: "abcdefg"})
	x.Add(&Member{Key: "hijklmn"})
	x.Add(&Member{Key: "opqrstu"})
	a, b, err := x.GetTwo("99999999")
	if err != nil {
		t.Fatal(err)
	}
	if a == b {
		t.Errorf("a shouldn't equal b")
	}
	if a != "abcdefg" {
		t.Errorf("wrong a: %q", a)
	}
	if b != "hijklmn" {
		t.Errorf("wrong b: %q", b)
	}
}

func TestGetTwoQuick(t *testing.T) {
	x := New()
	x.Add(&Member{Key: "abcdefg"})
	x.Add(&Member{Key: "hijklmn"})
	x.Add(&Member{Key: "opqrstu"})
	f := func(s string) bool {
		a, b, err := x.GetTwo(s)
		if err != nil {
			t.Logf("error: %q", err)
			return false
		}
		if a == b {
			t.Logf("a == b")
			return false
		}
		if a != "abcdefg" && a != "hijklmn" && a != "opqrstu" {
			t.Logf("invalid a: %q", a)
			return false
		}

		if b != "abcdefg" && b != "hijklmn" && b != "opqrstu" {
			t.Logf("invalid b: %q", b)
			return false
		}
		return true
	}
	if err := quick.Check(f, nil); err != nil {
		t.Fatal(err)
	}
}

func TestGetTwoOnlyTwoQuick(t *testing.T) {
	x := New()
	x.Add(&Member{Key: "abcdefg"})
	x.Add(&Member{Key: "hijklmn"})
	f := func(s string) bool {
		a, b, err := x.GetTwo(s)
		if err != nil {
			t.Logf("error: %q", err)
			return false
		}
		if a == b {
			t.Logf("a == b")
			return false
		}
		if a != "abcdefg" && a != "hijklmn" {
			t.Logf("invalid a: %q", a)
			return false
		}

		if b != "abcdefg" && b != "hijklmn" {
			t.Logf("invalid b: %q", b)
			return false
		}
		return true
	}
	if err := quick.Check(f, nil); err != nil {
		t.Fatal(err)
	}
}

func TestGetTwoOnlyOneInCircle(t *testing.T) {
	x := New()
	x.Add(&Member{Key: "abcdefg"})
	a, b, err := x.GetTwo("99999999")
	if err != nil {
		t.Fatal(err)
	}
	if a == b {
		t.Errorf("a shouldn't equal b")
	}
	if a != "abcdefg" {
		t.Errorf("wrong a: %q", a)
	}
	if b != "" {
		t.Errorf("wrong b: %q", b)
	}
}

func TestGetN(t *testing.T) {
	x := New()
	x.Add(&Member{Key: "abcdefg"})
	x.Add(&Member{Key: "hijklmn"})
	x.Add(&Member{Key: "opqrstu"})
	members, err := x.GetN("9999999", 3)
	if err != nil {
		t.Fatal(err)
	}
	if len(members) != 3 {
		t.Errorf("expected 3 members instead of %d", len(members))
	}
	if members[0] != "opqrstu" {
		t.Errorf("wrong members[0]: %q", members[0])
	}
	if members[1] != "abcdefg" {
		t.Errorf("wrong members[1]: %q", members[1])
	}
	if members[2] != "hijklmn" {
		t.Errorf("wrong members[2]: %q", members[2])
	}
}

func TestGetNLess(t *testing.T) {
	x := New()
	x.Add(&Member{Key: "abcdefg"})
	x.Add(&Member{Key: "hijklmn"})
	x.Add(&Member{Key: "opqrstu"})
	members, err := x.GetN("99999999", 2)
	if err != nil {
		t.Fatal(err)
	}
	if len(members) != 2 {
		t.Errorf("expected 2 members instead of %d", len(members))
	}
	if members[0] != "abcdefg" {
		t.Errorf("wrong members[0]: %q", members[0])
	}
	if members[1] != "hijklmn" {
		t.Errorf("wrong members[1]: %q", members[1])
	}
}

func TestGetNMore(t *testing.T) {
	x := New()
	x.Add(&Member{Key: "abcdefg"})
	x.Add(&Member{Key: "hijklmn"})
	x.Add(&Member{Key: "opqrstu"})
	members, err := x.GetN("9999999", 5)
	if err != nil {
		t.Fatal(err)
	}
	if len(members) != 3 {
		t.Errorf("expected 3 members instead of %d", len(members))
	}
	if members[0] != "opqrstu" {
		t.Errorf("wrong members[0]: %q", members[0])
	}
	if members[1] != "abcdefg" {
		t.Errorf("wrong members[1]: %q", members[1])
	}
	if members[2] != "hijklmn" {
		t.Errorf("wrong members[2]: %q", members[2])
	}
}

func TestGetNQuick(t *testing.T) {
	x := New()
	x.Add(&Member{Key: "abcdefg"})
	x.Add(&Member{Key: "hijklmn"})
	x.Add(&Member{Key: "opqrstu"})
	f := func(s string) bool {
		members, err := x.GetN(s, 3)
		if err != nil {
			t.Logf("error: %q", err)
			return false
		}
		if len(members) != 3 {
			t.Logf("expected 3 members instead of %d", len(members))
			return false
		}
		set := make(map[string]bool, 4)
		for _, member := range members {
			if set[member] {
				t.Logf("duplicate error")
				return false
			}
			set[member] = true
			if member != "abcdefg" && member != "hijklmn" && member != "opqrstu" {
				t.Logf("invalid member: %q", member)
				return false
			}
		}
		return true
	}
	if err := quick.Check(f, nil); err != nil {
		t.Fatal(err)
	}
}

func TestGetNLessQuick(t *testing.T) {
	x := New()
	x.Add(&Member{Key: "abcdefg"})
	x.Add(&Member{Key: "hijklmn"})
	x.Add(&Member{Key: "opqrstu"})
	f := func(s string) bool {
		members, err := x.GetN(s, 2)
		if err != nil {
			t.Logf("error: %q", err)
			return false
		}
		if len(members) != 2 {
			t.Logf("expected 2 members instead of %d", len(members))
			return false
		}
		set := make(map[string]bool, 4)
		for _, member := range members {
			if set[member] {
				t.Logf("duplicate error")
				return false
			}
			set[member] = true
			if member != "abcdefg" && member != "hijklmn" && member != "opqrstu" {
				t.Logf("invalid member: %q", member)
				return false
			}
		}
		return true
	}
	if err := quick.Check(f, nil); err != nil {
		t.Fatal(err)
	}
}

func TestGetNMoreQuick(t *testing.T) {
	x := New()
	x.Add(&Member{Key: "abcdefg"})
	x.Add(&Member{Key: "hijklmn"})
	x.Add(&Member{Key: "opqrstu"})
	f := func(s string) bool {
		members, err := x.GetN(s, 5)
		if err != nil {
			t.Logf("error: %q", err)
			return false
		}
		if len(members) != 3 {
			t.Logf("expected 3 members instead of %d", len(members))
			return false
		}
		set := make(map[string]bool, 4)
		for _, member := range members {
			if set[member] {
				t.Logf("duplicate error")
				return false
			}
			set[member] = true
			if member != "abcdefg" && member != "hijklmn" && member != "opqrstu" {
				t.Logf("invalid member: %q", member)
				return false
			}
		}
		return true
	}
	if err := quick.Check(f, nil); err != nil {
		t.Fatal(err)
	}
}

// allocBytes returns the number of bytes allocated by invoking f.
func allocBytes(f func()) uint64 {
	var stats runtime.MemStats
	runtime.ReadMemStats(&stats)
	t := stats.TotalAlloc
	f()
	runtime.ReadMemStats(&stats)
	return stats.TotalAlloc - t
}

func mallocNum(f func()) uint64 {
	var stats runtime.MemStats
	runtime.ReadMemStats(&stats)
	t := stats.Mallocs
	f()
	runtime.ReadMemStats(&stats)
	return stats.Mallocs - t
}

func BenchmarkAllocations(b *testing.B) {
	x := New()
	x.Add(&Member{Key: "stays"})
	b.ResetTimer()
	allocSize := allocBytes(func() {
		for i := 0; i < b.N; i++ {
			x.Add(&Member{Key: "Foo"})
			x.Remove("Foo")
		}
	})
	b.Logf("%d: Allocated %d bytes (%.2fx)", b.N, allocSize, float64(allocSize)/float64(b.N))
}

func BenchmarkMalloc(b *testing.B) {
	x := New()
	x.Add(&Member{Key: "stays"})
	b.ResetTimer()
	mallocs := mallocNum(func() {
		for i := 0; i < b.N; i++ {
			x.Add(&Member{Key: "Foo"})
			x.Remove("Foo")
		}
	})
	b.Logf("%d: Mallocd %d times (%.2fx)", b.N, mallocs, float64(mallocs)/float64(b.N))
}

func BenchmarkCycle(b *testing.B) {
	x := New()
	x.Add(&Member{Key: "nothing"})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x.Add(&Member{Key: "foo" + strconv.Itoa(i)})
		x.Remove("foo" + strconv.Itoa(i))
	}
}

func BenchmarkCycleLarge(b *testing.B) {
	x := New()
	for i := 0; i < 10; i++ {
		x.Add(&Member{Key: "start" + strconv.Itoa(i)})
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x.Add(&Member{Key: "foo" + strconv.Itoa(i)})
		x.Remove("foo" + strconv.Itoa(i))
	}
}

func BenchmarkGet(b *testing.B) {
	x := New()
	x.Add(&Member{Key: "nothing"})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x.Get("nothing")
	}
}

func BenchmarkGetLarge(b *testing.B) {
	x := New()
	for i := 0; i < 10; i++ {
		x.Add(&Member{Key: "start" + strconv.Itoa(i)})
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x.Get("nothing")
	}
}

func BenchmarkGetN(b *testing.B) {
	x := New()
	x.Add(&Member{Key: "nothing"})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x.GetN("nothing", 3)
	}
}

func BenchmarkGetNLarge(b *testing.B) {
	x := New()
	for i := 0; i < 10; i++ {
		x.Add(&Member{Key: "start" + strconv.Itoa(i)})
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x.GetN("nothing", 3)
	}
}

func BenchmarkGetTwo(b *testing.B) {
	x := New()
	x.Add(&Member{Key: "nothing"})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x.GetTwo("nothing")
	}
}

func BenchmarkGetTwoLarge(b *testing.B) {
	x := New()
	for i := 0; i < 10; i++ {
		x.Add(&Member{Key: "start" + strconv.Itoa(i)})
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x.GetTwo("nothing")
	}
}

// from @edsrzf on github:
func TestAddCollision(t *testing.T) {
	// These two strings produce several crc32 collisions after "|i" is
	// appended added by Consistent.eltKey.
	const s1 = "abear"
	const s2 = "solidiform"
	x := New()
	x.Add(&Member{Key: s1})
	x.Add(&Member{Key: s2})
	elt1, err := x.Get("abear")
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	y := New()
	// add elements in opposite order
	y.Add(&Member{Key: s2})
	y.Add(&Member{Key: s1})
	elt2, err := y.Get(s1)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	if elt1 != elt2 {
		t.Error(elt1, "and", elt2, "should be equal")
	}
}

// inspired by @or-else on github
func TestCollisionsCRC(t *testing.T) {
	t.SkipNow()
	c := New()
	f, err := os.Open("/usr/share/dict/words")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	found := make(map[uint32]string)
	scanner := bufio.NewScanner(f)
	count := 0
	for scanner.Scan() {
		word := scanner.Text()
		for i := 0; i < c.NumberOfReplicas; i++ {
			ekey := c.eltKey(word, i)
			// ekey := word + "|" + strconv.Itoa(i)
			k := c.hashKey(ekey)
			exist, ok := found[k]
			if ok {
				t.Logf("found collision: %s, %s", ekey, exist)
				count++
			} else {
				found[k] = ekey
			}
		}
	}
	t.Logf("number of collisions: %d", count)
}
