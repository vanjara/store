package stringstore

type Store interface {
	Put(string)
	Get() string
}
