package image

import (
	"github.com/google/go-containerregistry/pkg/name"
	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/remote"
)

func Get(refName string) (v1.Image, error) {
	ref, err := name.ParseReference(refName)
	if err != nil {
		panic(err)
	}

	return remote.Image(ref)
}
