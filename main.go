package main

import "fmt"

type ImageCompressionStrategy interface {
	Compress(imageData []byte) []byte
}
type JpegCompressionStrategy struct{}

func (j *JpegCompressionStrategy) Compress(imageData []byte) []byte {
	fmt.Println("Сжатие изображения в формате JPEG")
	return []byte("JPEG compressed data")
}
type PngCompressionStrategy struct{}

func (p *PngCompressionStrategy) Compress(imageData []byte) []byte {
	fmt.Println("Сжатие изображения в формате PNG")
	return []byte("PNG compressed data")
}
type ImageProcessorContext struct {
	compressionStrategy ImageCompressionStrategy
}

func (c *ImageProcessorContext) SetCompressionStrategy(strategy ImageCompressionStrategy) {
	c.compressionStrategy = strategy
}

func (c *ImageProcessorContext) CompressImage(imageData []byte) []byte {
	return c.compressionStrategy.Compress(imageData)
}

func main() {
	jpegStrategy := &JpegCompressionStrategy{}
	pngStrategy := &PngCompressionStrategy{}
	imageProcessor := &ImageProcessorContext{}
	imageProcessor.SetCompressionStrategy(jpegStrategy)
	compressedData := imageProcessor.CompressImage([]byte("Image Data"))
	fmt.Println(string(compressedData))
	imageProcessor.SetCompressionStrategy(pngStrategy)
	compressedData = imageProcessor.CompressImage([]byte("Image Data"))
	fmt.Println(string(compressedData))
}