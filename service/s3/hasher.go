package s3

import "hash"

type Hasher interface {
	hash.Hash
	Close()
}
