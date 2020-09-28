package filetype

import (
	"github.com/h2non/filetype"
	"github.com/h2non/filetype/types"
)

type Type struct {
	types.Type
}

func (t *Type) ELF() bool {
	if t.Extension == "elf" {
		return true
	}
	return false
}

func Detect(buf []byte) (*Type, error) {
	t, err := filetype.Match(buf)
	if err != nil {
		return nil, err
	}
	return &Type{t}, nil
}
