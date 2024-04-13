package hash

type HashService interface {
	Hash(string) (string, error)
	Compare(data string, encrypted string) error
}
