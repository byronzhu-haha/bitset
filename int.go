package bitset

import (
	"bytes"
	"fmt"
)

type int64Set struct {
	len   int
	items []uint64
}

func NewBitSet() BitSet {
	return &int64Set{}
}

func (s *int64Set) unpack(i int) (idx, bit int) {
	return i >> 6, i & 63
}

func (s *int64Set) assert(intN interface{}) int {
	i, ok := intN.(int)
	if !ok {
		panic("the type of intN must be int")
	}
	return i
}

// Add 加入一个int型的元素
func (s *int64Set) Add(intN interface{}) {
	i := s.assert(intN)
	if i < 0 || s.Has(i) {
		return
	}
	var (
		idx, bit = s.unpack(i)
	)
	// 扩容
	s.grow(idx)
	s.len++
	s.items[idx] |= 1 << bit
}

func (s *int64Set) grow(idx int) {
	var (
		iLen = len(s.items)
	)
	if idx < iLen {
		return
	}
	// 如果idx比当前数组长度大，则触发扩容操作
	// 1、idx是iLen的两倍及以上，且iLen大于等于1024，则重新分配，然后进行拷贝
	// 2、两倍之内，直接append
	if (idx >> 1) >= iLen && iLen >= 1024 {
		alloc := make([]uint64, idx+1)
		copy(alloc, s.items)
		s.items = alloc
		return
	}

	for i := idx - iLen; i >= 0; i-- {
		s.items = append(s.items, 0)
	}
}

func (s *int64Set) Has(intN interface{}) bool {
	i := s.assert(intN)
	if i < 0 {
		return false
	}
	idx, bit := s.unpack(i)
	return idx < len(s.items) && s.items[idx]&(1<<bit) != 0
}

func (s *int64Set) Len() int {
	return s.len
}

func (s *int64Set) Clear() {
	s.len = 0
	s.items = nil
}

func (s *int64Set) Remove(intN interface{}) {
	if !s.Has(intN) {
		return
	}
	i := s.assert(intN)
	idx, bit := s.unpack(i)
	if idx >= len(s.items) {
		return
	}
	s.len--
	s.items[idx] &= ^(1 << bit)
}

func (s *int64Set) Clone() BitSet {
	items := make([]uint64, len(s.items))
	copy(items, s.items)
	return &int64Set{
		len:   s.len,
		items: items,
	}
}

// Union 取并集，结果保存在inset中
func (s *int64Set) Union(inset BitSet) {
	iset, ok := inset.(*int64Set)
	if !ok {
		return
	}
	for i, item := range s.items {
		if i < len(iset.items) {
			iset.items[i] |= item
			iset.len += s.countDiffBit(iset.items[i], item)
		} else {
			iset.items = append(iset.items, item)
			iset.len += s.countDiffBit(item, 0)
		}
	}
}

func (s *int64Set) countDiffBit(a, b uint64) int {
	var (
		//异或
		ans = a ^ b
		count int
	)
	for ans != 0 {
		ans &= ans -1
		count++
	}
	return count
}

// String 位图中元素不易理解，利用String()可让有关位图的打印输出更直白
func (s *int64Set) String() string {
	var buf bytes.Buffer
	n, _ := buf.WriteString("[")
	for i, item := range s.items {
		if item == 0 {
			continue
		}
		for j := uint(0); j < 64; j++ {
			if item&(1<<j) != 0 {
				if buf.Len() > n {
					buf.WriteString(", ")
				}
				_, _ = fmt.Fprintf(&buf, "%d", 64*i+int(j))
			}
		}
	}
	buf.WriteString("]")
	return buf.String()
}
