package limit

import "io"

type Reader struct {
	n      int
	reader io.Reader
}

func (lr *Reader) Read(lp []byte) (int, error) {
	//p := make([]byte, lr.n)
	//m, err := lr.reader.Read(p)
	m, err := lr.reader.Read(lp)
	if m == lr.n {
		//copy(lp, p)
		return m, io.EOF
	}
	if m < lr.n && err == io.EOF {
		//copy(lp, p)
		return m, err
	}
	if m < lr.n && err != nil {
		//copy(lp, p)
		lr.n -= m
		return m, err
	}
	//p1 := p[:lr.n]
	//copy(lp, p1)
	lp = lp[:lr.n]
	return lr.n, io.EOF
}

func LimitReader(r io.Reader, n int) io.Reader {
	return &Reader{n, r}
}
