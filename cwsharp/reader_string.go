package cwsharp

type buffReader struct {
	src []rune
	off int
	Reader
}

func (r *buffReader) Init(src []rune) {
	r.src = src
	r.off = 0
}

func (r *buffReader) ReadRule() rune {
	if r.off == len(r.src) {
		return EOF
	}
	ch := r.src[r.off]
	r.off += 1
	return ch
}

func (r *buffReader) Peek() rune {
	if r.off == len(r.src) {
		return EOF
	}
	return r.src[r.off]
}

func (r *buffReader) Seek(offset int) {
	r.off = offset
}

func (r *buffReader) Pos() int {
	return r.off
}

func ReadString(src string) Reader {
	r := &buffReader{}
	r.Init([]rune(src))
	return r
}