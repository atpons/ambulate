package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/atpons/ambulate/pkg/filetype"
	"github.com/atpons/ambulate/pkg/image"
	"github.com/atpons/ambulate/pkg/layer"
	"log"
	"os"
)

var (
	refName string
)

func main() {
	if err := command(); err != nil {
		fmt.Fprintf(os.Stderr, "%+v", err)
		os.Exit(1)
	}
}

func command() error {
	flag.StringVar(&refName, "name", "", "")
	flag.Parse()

	img, err := image.Get(refName)
	if err != nil {
		return err
	}

	cfg, err := img.ConfigFile()
	if err != nil {
		return err
	}

	entrypoints := cfg.Config.Entrypoint
	if len(entrypoints) < 1 {
		return errors.New("not found entrypoints, cannot to analyze")
	}

	ext, err := layer.ExtractSingleFile(img, entrypoints[0])
	if err != nil {
		return err
	}

	t, err := filetype.Detect(ext.Bytes())
	if err != nil {
		return err
	}

	if t.ELF() {
		log.Println("elf binary detected")
		return nil
	}

	log.Printf("elf binary is not found mime=%s extension=%s", t.MIME, t.Extension)
	return nil
}
