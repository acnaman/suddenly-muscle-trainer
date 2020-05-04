package main

import (
	"archive/zip"
	"flag"
	"fmt"
	"io"
	"os"
	"path"
)

func main() {
	var (
		outfile string
		infiles []string
	)
	flag.StringVar(&outfile, "o", "out.zip", "name of zip file")
	flag.Parse()
	infiles = flag.Args()

	fmt.Println("infiles:", infiles)

	dest, err := os.Create(outfile)
	if err != nil {
		panic(err)
	}

	zipWriter := zip.NewWriter(dest)
	//defer zipWriter.Close()

	for _, s := range infiles {
		if err := addToZip(s, zipWriter); err != nil {
			panic(err)
		}
	}
	zipWriter.Close()
}

func addToZip(filename string, zipWriter *zip.Writer) error {
	src, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer src.Close()

	info, err := src.Stat()
	if err != nil {
		return err
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}
	fname := path.Base(filename)
	header.Name = fname
	//	header.Method = zip.Deflate

	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}
	_, err = io.Copy(writer, src)
	if err != nil {
		return err
	}

	return nil
}
