package stringstore

type StructStore struct {
	data string
}

func NewStructStore() *StructStore {
	return &StructStore{}
}

func (s *StructStore) Put(data string) {
	s.data = data
}

func (s *StructStore) Get() string {
	return s.data
}

