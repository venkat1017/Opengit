package objects

import (
	"crypto/sha1"
	"fmt"
)

type Blob struct {
	Data []byte
}

func NewBlob(data []byte) *Blob {
	return &Blob{Data: data}
}

func (b *Blob) Serialize() []byte {
	header := fmt.Sprintf("blob %d\x00", len(b.Data))
	return append([]byte(header), b.Data...)
}

func (b *Blob) Hash() string {
	data := b.Serialize()
	hash := sha1.Sum(data)
	return fmt.Sprintf("%x", hash)
}
