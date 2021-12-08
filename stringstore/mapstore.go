package stringstore

type MapStore struct {
	data map[string]bool
}

func NewMapStore() *MapStore {
	return &MapStore{
		data: map[string]bool{},
	}
}

func (s *MapStore) Put(data string) {
	s.data[data] = true
}

func (s *MapStore) Get() string {
	for k := range s.data {
		return k
	}
	return ""
}
