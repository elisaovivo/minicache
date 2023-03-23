package cache_

type ByteView struct {
	b []byte
}

func (v ByteView) Len() int {
	return len(v.b)
}

func cloneBytes(b []byte) []byte {
	tmp := make([]byte, len(b))
	copy(tmp, b)
	return tmp
}

func (v ByteView) String() string {
	return string(v.b)
}

func (v ByteView) ByteSlice() []byte {
	return cloneBytes(v.b)
}
