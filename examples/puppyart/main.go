package main

import (
	"fmt"
	"github.com/xyproto/imagelib"
	"image"
	"image/color"
	"image/draw"
)

// Convert in.png to blue and yellow colorplates (mixed to green)
// Higher threshold includes more colors.
func convert(infilename string, thresh uint8, t float64, color1, color2 color.RGBA) image.Image {
	mixcolor := imagelib.PaintMix(color1, color2)

	fmt.Println("(", mixcolor.R, mixcolor.G, mixcolor.B, ")")

	// Input image
	img := imagelib.ReadPNG(infilename)

	// Separate the image to color1, mixcolor and color2, with given threshold
	images := imagelib.Separate(img, color1, mixcolor, color2, thresh, t)
	color1image, mixcolorimage, color2image := images[0], images[1], images[2]

	// Add the green image to the blue image with cyan color
	// pseudo: blueplate = blueimage + greenimage in blue
	color1plate := imagelib.AddToAs(color1image, mixcolorimage, color1)

	// Add the green image to the color2 image with color2 color
	// pseudo: color2plate = color2image + greenimage in color2
	//color2plate := imagelib.AddToAs(color2image, color2image, color2)
	//color2plate = imagelib.AddToAs(color2image, mixcolorimage, color2)

	// Add the color2 to the cyan and the green to the rest
	allplate := imagelib.AddToAs(color1plate, color2image, color2)
	allplate = imagelib.AddToAs(allplate, mixcolorimage, mixcolor)

	// Write out all the images
	//imagelib.WritePNG(color1plate, color1filename)
	//imagelib.WritePNG(color2plate, color2filename)
	//imagelib.WritePNG(allplate, outfilename)
	return allplate
}

func main() {
	color1 := color.RGBA{0, 0, 255, 255}     // blue
	color2 := color.RGBA{255, 255, 255, 255} // white
	image1 := convert("puppy.png", 255, 1.0, color1, color2)
	color1 = color.RGBA{16, 63, 255, 255} // greenblue
	color2 = color.RGBA{0, 0, 0, 255}     // black
	image2 := convert("puppy.png", 255, 1.0, color1, color2)
	color1 = color.RGBA{255, 0, 0, 255}   // red
	color2 = color.RGBA{255, 255, 0, 255} // yellow
	image3 := convert("puppy.png", 255, 1.0, color1, color2)
	color1 = color.RGBA{16, 255, 255, 255} // turqoise
	color2 = color.RGBA{255, 0, 255, 255}  // purple
	image4 := convert("puppy.png", 255, 1.0, color1, color2)

	width := image1.Bounds().Dx() * 2
	height := image2.Bounds().Dy() * 2

	newimage := image.NewRGBA(image.Rect(0, 0, width, height))

	destrect := image1.Bounds()
	zero := newimage.Bounds().Min
	draw.Draw(newimage, destrect, image1, zero, draw.Src)
	destrect = image.Rect(image1.Bounds().Dx(), 0, width, image1.Bounds().Dy())
	draw.Draw(newimage, destrect, image2, zero, draw.Src)
	destrect = image.Rect(0, image1.Bounds().Dy(), image1.Bounds().Dx(), height)
	draw.Draw(newimage, destrect, image3, zero, draw.Src)
	destrect = image.Rect(image1.Bounds().Dx(), image1.Bounds().Dy(), width, height)
	fmt.Println(destrect)
	draw.Draw(newimage, destrect, image4, zero, draw.Src)

	imagelib.WritePNG(newimage, "out.png")
}
