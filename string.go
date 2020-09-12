package bitset

type stringSet struct {
	set *int64Set
}

func NewStringBitSet() BitSet {
	return &stringSet{set: &int64Set{}}
}

func (s *stringSet) Add(str interface{}) {

}

func (s *stringSet) Has(str interface{}) bool {

	return false
}

func (s *stringSet) Len() int {
	return s.set.Len()
}

func (s *stringSet) Clear() {
	s.set.Clear()
}

func (s *stringSet) Remove(str interface{}) {

}

func (s *stringSet) Union(t BitSet) {

}

func (s *stringSet) Clone() BitSet {
	return &stringSet{
		set: s.set.Clone().(*int64Set),
	}
}

func (s *stringSet) String() string {

	return ""
}