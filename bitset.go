package bitset

type BitSet interface {
	Add(i interface{})
	Has(i interface{}) bool
	Len() int
	Clear()
	Remove(i interface{})
	Union(t BitSet)
	Clone() BitSet
	String() string
}
