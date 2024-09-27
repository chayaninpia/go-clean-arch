package bytecompare

import (
	"bytes"
	"encoding/gob"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) CompareBy2Word(word1, word2 string) (int, int) {
	first := s.CompareByte("a")
	second := s.CompareByte("ก")
	return first, second
}

func (s *Service) CompareByte(word string) int {
	b := new(bytes.Buffer)
	gob.NewEncoder(b).Encode(word)
	return b.Len()
}
