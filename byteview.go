package zcache

type ByteView []byte

func (v ByteView) Len() int {
	return len(v)
}

func (v ByteView) ByteSlice() []byte {
	return cloneBytes(v)
}

func (v ByteView) String()  string {
	return string(v)
}

func cloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
