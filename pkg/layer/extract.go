package layer

import (
	"archive/tar"
	"bytes"
	"errors"
	v1 "github.com/google/go-containerregistry/pkg/v1"
	"io"
	"io/ioutil"
	"log"
	"path"
)

// ref to https://github.com/wagoodman/dive/blob/35db1ec914ddfb2e206e1e364bd8c1661d73a47d/dive/image/docker/image_archive.go
func ExtractSingleFile(img v1.Image, fpath string) (*bytes.Buffer, error) {
	manifest, err := img.Manifest()
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	for _, l := range manifest.Layers {
		lbd, err := img.LayerByDigest(l.Digest)
		if err != nil {
			panic(err)
		}

		rc, _ := lbd.Uncompressed()
		rbuf, err := ioutil.ReadAll(rc)
		if err != nil {
			return nil, err
		}

		tarReader := tar.NewReader(bytes.NewReader(rbuf))
		for {
			header, err := tarReader.Next()
			if err == io.EOF {
				break
			} else if err != nil {
				return nil, err
			}

			nm := path.Clean(header.Name)
			if nm == "." {
				continue
			}

			switch header.Typeflag {
			case tar.TypeDir, tar.TypeXHeader, tar.TypeXGlobalHeader:
			default:
				if "/"+nm == fpath {
					log.Println("detected file!")
					io.Copy(buf, tarReader)
					return buf, nil
				}
			}
		}
	}
	return nil, errors.New("not found single file")
}
