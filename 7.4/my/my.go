package my

import "io"

type Reader struct {
	s string
}

func (r *Reader) Read(p []byte) (int, error) {
	copy(p, []byte(r.s))
	return len(r.s), io.EOF
}

func NewReader(s string) *Reader {
	return &Reader{s}
}
