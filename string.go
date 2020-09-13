package bitset

import (
	"math/rand"
	"time"
)

const stub = 1 << 32

type simpleHash struct {
	seed int
}

func (h *simpleHash) hash(value string) int {
	var result int
	for i := 0; i < len(value); i++ {
		result = result*h.seed + int(value[i])
	}
	return stub & result
}

type BloomFilter struct {
	len     int
	set     BitSet
	hashFns [7]simpleHash
}

func NewBloomFilter() *BloomFilter {
	f := &BloomFilter{set: NewBitSet()}
	f.genHashFns()
	return f
}

func (s *BloomFilter) genHashFns() {
	rand.Seed(time.Now().UnixNano())
	var itr int
	for i := 0; i < len(s.hashFns); i++ {
		seed := rand.Intn(10+(1<<i)) + itr
		itr = seed
		s.hashFns[i] = simpleHash{
			seed: seed,
		}
	}
}

func (s *BloomFilter) Add(str string) {
	if str == "" {
		return
	}
	if s.Has(str) {
		return
	}
	s.len++
	for _, fn := range s.hashFns {
		s.set.Add(fn.hash(str))
	}
}

func (s *BloomFilter) Has(str string) (res bool) {
	if str == "" {
		return false
	}
	res = true
	for _, f := range s.hashFns {
		res = res && s.set.Has(f.hash(str))
	}
	return res
}

func (s *BloomFilter) Len() int {
	return s.len
}

func (s *BloomFilter) Clear() {
	s.len = 0
	s.set.Clear()
	s.hashFns = [7]simpleHash{}
}

func (s *BloomFilter) Clone() *BloomFilter {
	return &BloomFilter{
		len:     s.len,
		set:     s.set.Clone().(*int64Set),
		hashFns: s.hashFns,
	}
}
