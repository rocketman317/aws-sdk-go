package s3

import (
	"crypto/md5"
	"crypto/sha256"
)

type HasherFactory interface {
	AcquireMD5() Hasher
	AcquireSHA256() Hasher
}

type defaultHasherFactory struct {
	//
}

func (f *defaultHasherFactory) AcquireMD5() Hasher {
	return HashWrapper(md5.New())
}

func (f *defaultHasherFactory) AcquireSHA256() Hasher {
	return HashWrapper(sha256.New())
}
