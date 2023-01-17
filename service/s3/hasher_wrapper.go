package s3

import "hash"

type Wrapper struct {
	hash.Hash
}

// HashWrapper function transforms a hash.Hash into a Hasher interface
func HashWrapper(h hash.Hash) Hasher {
	return &Wrapper{h}
}

// Close function is a no-op for FallbackWrapper
func (w *Wrapper) Close() {
	// noop
}

// Sum function is a wrapper for hash.Hash::Sum
func (w *Wrapper) Sum(b []byte) []byte {
	return w.Hash.Sum(b)
}

// Write function is a wrapper for hash.Hash::Write
func (w *Wrapper) Write(p []byte) (n int, err error) {
	return w.Hash.Write(p)
}

func (w *Wrapper) Size() int {
	return w.Hash.Size()
}

func (w *Wrapper) BlockSize() int {
	return w.Hash.BlockSize()
}

func (w *Wrapper) Reset() {
	w.Hash.Reset()
}
