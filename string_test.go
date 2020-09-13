package bitset

import "testing"

func TestBloomFilter_Add(t *testing.T) {
	f := NewBloomFilter()
	f.Add("add")
	if f.Len() != 1 {
		t.Error("add failed, len of f should be 1")
		return
	}
	f.Add("add")
	if f.Len() != 1 {
		t.Error("add failed, len of f should be 1")
		return
	}
	f.Add("add ")
	if f.Len() != 2 {
		t.Errorf("add failed, len of f should be 2, len: %+v", f.Len())
		return
	}
	f.Add("sub")
	if f.Len() != 3 {
		t.Error("add failed, len of f should be 3")
	}
}

func TestBloomFilter_Has(t *testing.T) {
	f := NewBloomFilter()
	f.Add("add")
	f.Add("add")
	f.Add("sub")
	f.Add("add ")
	if !f.Has("add") || !f.Has("sub") || !f.Has("add ") {
		t.Error("has failed, they should be in filter")
	}
}

func TestBloomFilter_Clear(t *testing.T) {
	f := NewBloomFilter()
	f.Add("333")
	f.Add("444")
	f.Add("555")
	f.Clear()
	if f.Len() != 0 {
		t.Error("clear failed, len of f should be 0")
		return
	}
	if f.Has("333") || f.Has("444") || f.Has("555") {
		t.Error("clear failed, they should be not in filter")
	}
}

func TestBloomFilter_Clone(t *testing.T) {
	f := NewBloomFilter()
	f.Add("333")
	f.Add("444")
	f.Add("555")
	nf := f.Clone()
	if nf.Len() != f.Len() {
		t.Error("clone failed, len of nf should be equal to len of f")
		return
	}
	if !nf.Has("333") || !nf.Has("444") || !nf.Has("555") {
		t.Error("clone failed, they should be in nf")
	}
}
