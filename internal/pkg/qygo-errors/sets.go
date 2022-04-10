package qygo_errors

type Empty struct{}

type String map[string]Empty

func NewString(items ...string) String {
	ss := String{}
	ss.Insert(items...)
	return ss
}

func (s String) Insert(items ...string) String {
	for _, item := range items {
		s[item] = Empty{}
	}
	return s
}
func (s String) Has(item string) bool {
	_, contained := s[item]
	return contained
}
