package bitstream

import (
	"github.com/gravestench/bitstream/pkg"
)

// these exports are here to prevent you from needing to import from pkg

type (
	Bits   = pkg.Bits
	Reader = pkg.Reader
	Writer = pkg.Writer
)

func NewReader(args ...interface{}) *Reader {
	return pkg.NewReader(args)
}

func ReaderFromBytes(data ...byte) *Reader {
	return pkg.NewReader(data)
}

func BitsFromByte(bb byte) Bits {
	return pkg.BitsFromByte(bb)
}
