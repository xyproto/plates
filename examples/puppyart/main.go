package main

import (
	"fmt"
	"github.com/xyproto/imagelib"
	"image/color"
)

// Convert in.png to blue and yellow colorplates (mixed to green)
// Higher threshold includes more colors.
func convert(infilename string, bluefilename string, yellowfilename string, allfilename string, thresh uint8, t float64) {
	// Colors
	yellow := color.RGBA{255, 255, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	//red := imagelib.RGBAColor{255, 0, 0, 255}

	green := imagelib.PaintMix(yellow, blue) //imagelib.RGBAColor{0, 255, 0, 255}

	fmt.Println("(", green.R, green.G, green.B, ")")

	// Input image
	img := imagelib.ReadPNG(infilename)

	// Separate the image to blue, mixcolor (green) and yellow, with given threshold
	images := imagelib.Separate(img, blue, green, yellow, thresh, t)
	blueimage, greenimage, yellowimage := images[0], images[1], images[2]

	// Add the green image to the blue image with cyan color
	// blueplate = blueimage + greenimage in blue
	blueplate := imagelib.AddToAs(blueimage, greenimage, blue)

	// Add the green image to the yellow image with yellow color
	// yellowplate = yellowimage + greenimage in yellow
	yellowplate := imagelib.AddToAs(yellowimage, yellowimage, yellow)
	yellowplate = imagelib.AddToAs(yellowimage, greenimage, yellow)

	// Add the yellow to the cyan and the green to the rest
	allplate := imagelib.AddToAs(blueplate, yellowimage, yellow)
	allplate = imagelib.AddToAs(allplate, greenimage, green)

	// Write out all the images
	imagelib.WritePNG(blueplate, bluefilename)
	imagelib.WritePNG(yellowplate, yellowfilename)
	imagelib.WritePNG(allplate, allfilename)
}

func main() {
	convert("puppy.png", "outblue.png", "outyellow.png", "outall.png", 255, 1.0)
}
