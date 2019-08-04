// Adapted from stathat.com/c/consistent which uses the MIT license.
// Please see the LICENSE file in the same folder.

// Package hashring implements a consistent hashing algorithm for gRPC
// load balancing.
package hashring

import (
	"errors"
	"hash/crc32"
	"sort"
	"strconv"
	"sync"
)

type uints []uint32

// Len returns the length of the uints array.
func (x uints) Len() int { return len(x) }

// Less returns true if element i is less than element j.
func (x uints) Less(i, j int) bool { return x[i] < x[j] }

// Swap exchanges elements i and j.
func (x uints) Swap(i, j int) { x[i], x[j] = x[j], x[i] }

// ErrEmptyCircle is the error returned when trying to get an element
// when nothing has been added to hash.
var ErrEmptyCircle = errors.New("empty circle")

// Member represents the key and value of the member in ring.
type Member struct {
	Key string
	Val interface{}
}

// HashRing holds the information about the members of the consistent
// hash circle.
type HashRing struct {
	circle           map[uint32]string
	members          map[string]Member
	sortedHashes     uints
	NumberOfReplicas int
	count            int64
	scratch          [64]byte
	sync.RWMutex
}

// New creates a new HashRing object with a default setting of 20 replicas
// for each entry.
//
// To change the number of replicas, set NumberOfReplicas before adding entries.
func New() *HashRing {
	hr := new(HashRing)
	hr.NumberOfReplicas = 20
	hr.circle = make(map[uint32]string)
	hr.members = make(map[string]Member)
	return hr
}

// eltKey generates a string key for an element with an index.
func (hr *HashRing) eltKey(elt string, idx int) string {
	// return elt + "|" + strconv.Itoa(idx)
	return strconv.Itoa(idx) + elt
}

// Add inserts a string element in the consistent hash.
func (hr *HashRing) Add(member *Member) {
	hr.Lock()
	defer hr.Unlock()
	hr.add(member)
}

func (hr *HashRing) add(member *Member) {
	key := member.Key
	for i := 0; i < hr.NumberOfReplicas; i++ {
		hr.circle[hr.hashKey(hr.eltKey(key, i))] = key
	}
	hr.members[key] = *member
	hr.updateSortedHashes()
	hr.count++
}

// Remove removes an element from the hash.
func (hr *HashRing) Remove(key string) {
	hr.Lock()
	defer hr.Unlock()
	hr.remove(key)
}

func (hr *HashRing) remove(key string) {
	for i := 0; i < hr.NumberOfReplicas; i++ {
		delete(hr.circle, hr.hashKey(hr.eltKey(key, i)))
	}
	delete(hr.members, key)
	hr.updateSortedHashes()
	hr.count--
}

// Members returns all members in a slice.
func (hr *HashRing) Members() []Member {
	hr.RLock()
	defer hr.RUnlock()
	var m []Member
	for _, v := range hr.members {
		m = append(m, v)
	}
	return m
}

// GetMember returns a member by key.
func (hr *HashRing) GetMember(key string) *Member {
	hr.RLock()
	defer hr.RUnlock()

	if m, ok := hr.members[key]; ok {
		return &m
	}
	return nil
}

// Get returns an element close to where name hashes to in the circle.
func (hr *HashRing) Get(name string) (string, error) {
	hr.RLock()
	defer hr.RUnlock()
	if len(hr.circle) == 0 {
		return "", ErrEmptyCircle
	}
	key := hr.hashKey(name)
	i := hr.search(key)
	return hr.circle[hr.sortedHashes[i]], nil
}

func (hr *HashRing) search(key uint32) (i int) {
	f := func(x int) bool {
		return hr.sortedHashes[x] > key
	}
	i = sort.Search(len(hr.sortedHashes), f)
	if i >= len(hr.sortedHashes) {
		i = 0
	}
	return
}

// GetTwo returns the two closest distinct elements to the name input in the circle.
func (hr *HashRing) GetTwo(name string) (string, string, error) {
	hr.RLock()
	defer hr.RUnlock()
	if len(hr.circle) == 0 {
		return "", "", ErrEmptyCircle
	}
	key := hr.hashKey(name)
	i := hr.search(key)
	a := hr.circle[hr.sortedHashes[i]]

	if hr.count == 1 {
		return a, "", nil
	}

	start := i
	var b string
	for i = start + 1; i != start; i++ {
		if i >= len(hr.sortedHashes) {
			i = 0
		}
		b = hr.circle[hr.sortedHashes[i]]
		if b != a {
			break
		}
	}
	return a, b, nil
}

// GetN returns the N closest distinct elements to the name input in the circle.
func (hr *HashRing) GetN(name string, n int) ([]string, error) {
	hr.RLock()
	defer hr.RUnlock()

	if len(hr.circle) == 0 {
		return nil, ErrEmptyCircle
	}

	if hr.count < int64(n) {
		n = int(hr.count)
	}

	var (
		key   = hr.hashKey(name)
		i     = hr.search(key)
		start = i
		res   = make([]string, 0, n)
		elem  = hr.circle[hr.sortedHashes[i]]
	)

	res = append(res, elem)

	if len(res) == n {
		return res, nil
	}

	for i = start + 1; i != start; i++ {
		if i >= len(hr.sortedHashes) {
			i = 0
		}
		elem = hr.circle[hr.sortedHashes[i]]
		if !sliceContainsMember(res, elem) {
			res = append(res, elem)
		}
		if len(res) == n {
			break
		}
	}

	return res, nil
}

func (hr *HashRing) hashKey(key string) uint32 {
	if len(key) < 64 {
		var scratch [64]byte
		copy(scratch[:], key)
		return crc32.ChecksumIEEE(scratch[:len(key)])
	}
	return crc32.ChecksumIEEE([]byte(key))
}

func (hr *HashRing) updateSortedHashes() {
	hashes := hr.sortedHashes[:0]
	//reallocate if we're holding on to too much (1/4th)
	if cap(hr.sortedHashes)/(hr.NumberOfReplicas*4) > len(hr.circle) {
		hashes = nil
	}
	for k := range hr.circle {
		hashes = append(hashes, k)
	}
	sort.Sort(hashes)
	hr.sortedHashes = hashes
}

func sliceContainsMember(set []string, member string) bool {
	for _, m := range set {
		if m == member {
			return true
		}
	}
	return false
}
