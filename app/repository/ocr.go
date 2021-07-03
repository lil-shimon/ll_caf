package repository

import "image"

func GetImageData() image.Image {
	src := file.Open()
	image, _, _ := image.Decode(src)
	return image
}
