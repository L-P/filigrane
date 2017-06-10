package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	_ "image/png"
	"log"
	"os"
	"path/filepath"
)

func main() {
	l := log.New(os.Stderr, "", 0)
	mark, err := getMark()
	if err != nil {
		panic(err)
	}

	for _, path := range os.Args[1:] {
		name := filepath.Base(path)
		dir := filepath.Dir(path)

		fmt.Printf("Watermarking %s\n", name)
		err := watermark(
			path,
			fmt.Sprintf("%s/filigrane_%s", dir, name),
			mark,
		)

		if err != nil {
			l.Println(err)
		}
	}
}

func getMark() (image.Image, error) {
	dir := filepath.Dir(os.Args[0])
	file, err := os.Open(dir + "/watermark.png")
	if err != nil {
		return nil, fmt.Errorf("Unable to open watermark: %s", err)
	}

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("Unable to decode watermark: %s", err)
	}

	return img, err
}

func watermark(path, newPath string, mark image.Image) error {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("Unable to open %s: %s", path, err)
	}
	defer file.Close()

	raw, _, err := image.Decode(file)
	if err != nil {
		return fmt.Errorf("Unable to decode %s: %s", path, err)
	}

	dest, err := os.Create(newPath)
	if err != nil {
		return fmt.Errorf("Unable to create dest %s: %s", newPath, err)
	}

	img := toRGBA(raw.(*image.YCbCr))

	draw.Draw(img, img.Bounds(), mark, image.Point{
		-(img.Bounds().Max.X - mark.Bounds().Max.X),
		-(img.Bounds().Max.Y - mark.Bounds().Max.Y),
	}, draw.Over)

	opt := &jpeg.Options{Quality: 80}
	if err := jpeg.Encode(dest, img, opt); err != nil {
		return fmt.Errorf("Unable to encode %s: %s", newPath, err)
	}

	return nil
}

func toRGBA(src *image.YCbCr) *image.RGBA {
	b := src.Bounds()
	m := image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))

	draw.Draw(m, m.Bounds(), src, b.Min, draw.Src)

	return m
}
