// Package util contains utility code for demosntration of go-libjpeg.
package util

import (
	"bufio"
	"image"
	"image/png"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

// SubsampledImages contains example images which have subsampling colors.
var SubsampledImages = []string{
	"../testdata/images/checkerboard_444.jpg",
	"../testdata/images/checkerboard_440.jpg",
	"../testdata/images/checkerboard_422.jpg",
	"../testdata/images/checkerboard_420.jpg",
}

// GetOutFilePath returns the path of specified out file.
func GetOutFilePath(name string) string {
    return filepath.Join("../testdata/out", name)
}

// OpenFile opens specified example file
func OpenFile(name string) (io io.Reader) {
	io, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	return
}

// ReadFile reads and returns data bytes of specified example file.
func ReadFile(name string) (data []byte) {
	data, err := ioutil.ReadFile(name)
	if err != nil {
		panic(err)
	}
	return
}

// CreateFile opens specified example file
func CreateFile(name string) (f *os.File) {
	f, err := os.Create(GetOutFilePath(name))
	if err != nil {
		panic(err)
	}
	return
}

// WritePNG encodes and writes image into PNG file.
func WritePNG(img image.Image, name string) {
	f, err := os.Create(GetOutFilePath(name))
	if err != nil {
		panic(err)
	}
	b := bufio.NewWriter(f)
	defer func() {
		b.Flush()
		f.Close()
	}()

	if err := png.Encode(b, img); err != nil {
		panic(err)
	}
	return
}
