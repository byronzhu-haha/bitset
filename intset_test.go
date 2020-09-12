package bitset

import (
	"fmt"
	"testing"
)

func TestInt64Set_Add(t *testing.T) {
	set := NewBitSet()
	set.Add(1)
	if set.Len() != 1 {
		t.Error("add failed, len of set should be 1")
		return
	}
	set.Add(1)
	if set.Len() != 1 {
		t.Error("add failed, len of set should be 1")
		return
	}
	set.Add(100)
	if set.Len() != 2 {
		t.Error("add failed, len of set should be 2")
		return
	}
	set.Add(64*1024)
	if set.Len() != 3 {
		t.Error("add failed, len of set should be 3")
		return
	}
	set.Add(64*10240)
	if set.Len() != 4 {
		t.Error("add failed, len of set should be 4")
	}
}

func TestInt64Set_Has(t *testing.T) {
	set := NewBitSet()
	set.Add(100)
	set.Add(64*1024)
	set.Add(64*10240)
	if !set.Has(100) || !set.Has(64*1024) || !set.Has(64*10240) {
		t.Error("has is incorrect, they are in set")
		return
	}
	if set.Has(1000) || set.Has(64*1025) || set.Has(64*10250) {
		t.Error("has is incorrect, they are not int set")
	}
}

func TestInt64Set_Clear(t *testing.T) {
	set := NewBitSet()
	set.Add(100)
	set.Add(101)
	set.Add(102)
	set.Clear()
	if set.Len() != 0 || set.Has(100) || set.Has(101) || set.Has(102) {
		t.Error("clear() is incorrect, they are cleared")
	}
}

func TestInt64Set_Clone(t *testing.T) {
	set := NewBitSet()
	set.Add(100)
	set.Add(99)
	cp := set.Clone()
	if cp.Len() != set.Len() {
		t.Error("clone() is incorrect, cp.Len() should be equal set.Len()")
		return
	}
	if !cp.Has(100) || !cp.Has(99) {
		t.Error("clone() is incorrect, 100 and 99 should be in cp")
	}
}

func TestInt64Set_Remove(t *testing.T) {
	set := NewBitSet()
	set.Add(100)
	set.Add(99)
	set.Remove(100)
	if set.Len() != 1 {
		t.Error("remove() is incorrect, len of set should be 1")
		return
	}
	if set.Has(100) {
		t.Error("remove() is incorrect, 100 should be not in set")
	}
}

func TestInt64Set_Union(t *testing.T) {
	set := NewBitSet()
	set.Add(1)
	set.Add(2)
	set2 := NewBitSet()
	set2.Add(3)
	set2.Add(4)
	set.Union(set2)
	if set2.Len() != 4 {
		t.Error("union() is incorrect, len of set2 should be 4")
		return
	}
	if !set2.Has(1) || !set2.Has(2) {
		t.Error("union() is incorrect, they are in set")
	}
}

func TestInt64Set_String(t *testing.T) {
	set := NewBitSet()
	set.Add(1)
	set.Add(2)
	s := fmt.Sprintf("%+v", set)
	if s != "[1, 2]" {
		t.Error("string() is incorrect, s should be [1, 2]")
	}
}
