package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"log"

	"github.com/xyproto/plates"
)

// Constants
const (
	version = "1.0.6"
	usage   = `Usage: puppyart INPUT_IMAGE OUTPUT_IMAGE

Options:
  --help      Show this message and exit.
  --version   Show the program version and exit.`
)

// convert takes an image file and modifies it based on a specified threshold and two given colors.
// It returns a processed image.
func convert(infilename string, thresh uint8, color1, color2 color.RGBA) image.Image {
	// Combine the two colors to create a mixed color
	mixcolor := plates.PaintMix(color1, color2)

	// Load the input image
	img, err := plates.Read(infilename)
	if err != nil {
		log.Fatalln(err)
	}

	// Separate the image into three based on the specified threshold and intensity: color1, mixcolor, and color2.
	color1image, mixcolorimage, color2image := plates.Separate3(img, color1, mixcolor, color2, thresh)

	// Combine the color1 image and mixcolor image to create a new image plate with color1
	color1plate := plates.AddToAs(color1image, mixcolorimage, color1)

	// Combine the color1plate and color2image to create an allplate image with color2 color
	// And then combine the resulting image with the mixcolorimage to get the final output image.
	allplate := plates.AddToAs(color1plate, color2image, color2)
	allplate = plates.AddToAs(allplate, mixcolorimage, mixcolor)

	return allplate
}

func main() {
	helpFlag := flag.Bool("help", false, usage)
	versionFlag := flag.Bool("version", false, "Display version info")

	flag.Parse()

	// Display help info
	if *helpFlag {
		fmt.Println(usage)
		return
	}

	// Display version info
	if *versionFlag {
		fmt.Printf("Version: %s\n", version)
		return
	}

	if flag.NArg() < 2 {
		log.Fatalln("Please provide both input and output image names.")
	}

	infile := flag.Arg(0)
	outfile := flag.Arg(1)

	image1 := convert(infile, 255, color.RGBA{0, 0, 255, 255}, color.RGBA{255, 255, 255, 255})
	image2 := convert(infile, 255, color.RGBA{16, 63, 255, 255}, color.RGBA{0, 0, 0, 255})
	image3 := convert(infile, 255, color.RGBA{255, 0, 0, 255}, color.RGBA{255, 255, 0, 255})
	image4 := convert(infile, 255, color.RGBA{16, 255, 255, 255}, color.RGBA{255, 0, 255, 255})

	// Create a new image to contain the 4 processed images.
	width := image1.Bounds().Dx() * 2
	height := image2.Bounds().Dy() * 2
	newimage := image.NewRGBA(image.Rect(0, 0, width, height))

	// Arrange the 4 processed images on the new image.
	draw.Draw(newimage, image1.Bounds(), image1, newimage.Bounds().Min, draw.Src)
	draw.Draw(newimage, image.Rect(image1.Bounds().Dx(), 0, width, image1.Bounds().Dy()), image2, newimage.Bounds().Min, draw.Src)
	draw.Draw(newimage, image.Rect(0, image1.Bounds().Dy(), image1.Bounds().Dx(), height), image3, newimage.Bounds().Min, draw.Src)
	draw.Draw(newimage, image.Rect(image1.Bounds().Dx(), image1.Bounds().Dy(), width, height), image4, newimage.Bounds().Min, draw.Src)

	// Write the final image to file.
	if err := plates.Write(outfile, newimage); err != nil {
		log.Fatalln(err)
	}
}
