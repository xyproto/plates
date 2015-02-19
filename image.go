package image

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

var verbose = true

// Don't write information to stdout
func Quiet() {
	verbose = false
}

// Write information to stdout (on by default)
func Verbose() {
	verbose = true
}

// Read a PNG file into an Image structure
func ReadPNG(filename string) image.Image {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	// m is image
	if verbose {
		fmt.Printf("Reading %s...", filename)
	}
	m, err := png.Decode(f)
	if err != nil {
		panic(err)
	}
	if verbose {
		fmt.Println("done")
	}
	return m
}

// Write an Image structure to a PNG file
func WritePNG(m image.Image, filename string) {
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	if verbose {
		fmt.Printf("Writing %s...", filename)
	}
	if err := png.Encode(f, m); err != nil {
		panic(err)
	}
	if verbose {
		fmt.Println("done")
	}
}

// Pick out only the red colors from an image
func RedImage(m image.Image) image.Image {
	if verbose {
		fmt.Println(m.Bounds()) // image.Rectangle; .Min, .Max; .X, .Y
	}
	var (
		rect image.Rectangle = m.Bounds()
		c    color.Color
		cr   color.RGBA
	)
	newImage := image.NewRGBA(image.Rect(0, 0, rect.Max.X-rect.Min.X, rect.Max.Y-rect.Min.Y))
	for y := rect.Min.Y; y < rect.Max.Y; y++ {
		for x := rect.Min.X; x < rect.Max.X; x++ {
			c = m.At(x, y)      // c is RGBAColor, which implements Color
			cr = c.(color.RGBA) // this is needed
			c = color.RGBA{cr.R, 0, 0, cr.A}
			newImage.Set(x-rect.Min.X, y-rect.Min.Y, c)
		}
	}
	return newImage
}

// Pick out only the green colors from an image
func GreenImage(m image.Image) image.Image {
	if verbose {
		fmt.Println(m.Bounds()) // image.Rectangle; .Min, .Max; .X, .Y
	}
	var (
		rect image.Rectangle = m.Bounds()
		c    color.Color
		cr   color.RGBA
	)
	newImage := image.NewRGBA(image.Rect(0, 0, rect.Max.X-rect.Min.X, rect.Max.Y-rect.Min.Y))
	for y := rect.Min.Y; y < rect.Max.Y; y++ {
		for x := rect.Min.X; x < rect.Max.X; x++ {
			c = m.At(x, y)      // c is RGBAColor, which implements Color
			cr = c.(color.RGBA) // this is needed
			c = color.RGBA{0, cr.G, 0, cr.A}
			newImage.Set(x-rect.Min.X, y-rect.Min.Y, c)
		}
	}
	return newImage
}

// Pick out only the blue colors from an image
func BlueImage(m image.Image) image.Image {
	if verbose {
		fmt.Println(m.Bounds()) // image.Rectangle; .Min, .Max; .X, .Y
	}
	var (
		rect image.Rectangle = m.Bounds()
		c    color.Color
		cr   color.RGBA
	)
	newImage := image.NewRGBA(image.Rect(0, 0, rect.Max.X-rect.Min.X, rect.Max.Y-rect.Min.Y))
	for y := rect.Min.Y; y < rect.Max.Y; y++ {
		for x := rect.Min.X; x < rect.Max.X; x++ {
			c = m.At(x, y)      // c is RGBAColor, which implements Color
			cr = c.(color.RGBA) // this is needed
			c = color.RGBA{0, 0, cr.B, cr.A}
			newImage.Set(x-rect.Min.X, y-rect.Min.Y, c)
		}
	}
	return newImage
}

// Absolute value
func Abs(a uint8) uint8 {
	if a < 0 {
		return -a
	}
	return a
}

// Absolute value
func Fabs(a float64) float64 {
	if a < 0 {
		return -a
	}
	return a
}

// Pick out only the colors close to the given color,
// within a given threshold
func CloseTo1(m image.Image, target color.RGBA, thresh uint8) image.Image {
	if verbose {
		fmt.Println(m.Bounds()) // image.Rectangle; .Min, .Max; .X, .Y
	}
	var (
		rect       image.Rectangle = m.Bounds()
		c          color.Color
		cr         color.RGBA
		r, g, b, a uint8
	)
	newImage := image.NewRGBA(image.Rect(0, 0, rect.Max.X-rect.Min.X, rect.Max.Y-rect.Min.Y))
	for y := rect.Min.Y; y < rect.Max.Y; y++ {
		for x := rect.Min.X; x < rect.Max.X; x++ {
			c = m.At(x, y)      // c is RGBAColor, which implements Color
			cr = c.(color.RGBA) // this is needed
			r = 0
			g = 0
			b = 0
			a = cr.A
			if Abs(target.R-cr.R) < thresh {
				r = target.R
			}
			if Abs(target.G-cr.G) < thresh {
				g = target.G
			}
			if Abs(target.B-cr.B) < thresh {
				b = target.B
			}
			c = color.RGBA{r, g, b, a}
			newImage.Set(x-rect.Min.X, y-rect.Min.Y, c)
		}
	}
	return newImage
}

// Pick out only the colors close to the given color,
// within a given threshold. Make it uniform.
// Zero alpha to unused pixels in returned image.
func CloseTo2(m image.Image, target color.RGBA, thresh uint8) image.Image {
	if verbose {
		fmt.Println(m.Bounds()) // image.Rectangle; .Min, .Max; .X, .Y
	}
	var (
		rect       image.Rectangle = m.Bounds()
		c          color.Color
		cr         color.RGBA
		r, g, b, a uint8
	)
	newImage := image.NewRGBA(image.Rect(0, 0, rect.Max.X-rect.Min.X, rect.Max.Y-rect.Min.Y))
	for y := rect.Min.Y; y < rect.Max.Y; y++ {
		for x := rect.Min.X; x < rect.Max.X; x++ {
			c = m.At(x, y)      // c is RGBAColor, which implements Color
			cr = c.(color.RGBA) // this is needed
			r = 0
			g = 0
			b = 0
			a = 0
			if Abs(target.R-cr.R) < thresh || Abs(target.G-cr.G) < thresh || Abs(target.B-cr.B) < thresh {
				r = target.R
				g = target.G
				b = target.B
				a = cr.A
			}
			c = color.RGBA{r, g, b, a}
			newImage.Set(x-rect.Min.X, y-rect.Min.Y, c)
		}
	}
	return newImage
}

// Take orig, add the nontransparent colors from addimage, as addascolor
func AddToAs(orig image.Image, addimage image.Image, addcolor color.RGBA) image.Image {
	if verbose {
		fmt.Print("Adding one image to another...")
	}
	var rect image.Rectangle = addimage.Bounds()
	var c color.Color
	var cr, or color.RGBA
	var r, g, b, a uint8
	newImage := image.NewRGBA(image.Rect(0, 0, rect.Max.X-rect.Min.X, rect.Max.Y-rect.Min.Y))
	for y := rect.Min.Y; y < rect.Max.Y; y++ {
		for x := rect.Min.X; x < rect.Max.X; x++ {
			cr = addimage.At(x, y).(color.RGBA)
			or = orig.At(x, y).(color.RGBA)
			r = or.R
			g = or.G
			b = or.B
			a = or.A
			if cr.A > 0 {
				r = addcolor.R
				g = addcolor.G
				b = addcolor.B
				a = addcolor.A
			}
			c = color.RGBA{r, g, b, a}
			newImage.Set(x-rect.Min.X, y-rect.Min.Y, c)
		}
	}
	if verbose {
		fmt.Println("done")
	}
	return newImage
}

// Convert RGB to hue
func Hue(cr color.RGBA) float64 {
	r := float64(cr.R) / 255.0
	g := float64(cr.G) / 255.0
	b := float64(cr.B) / 255.0
	var h float64
	rgb_max := r
	if g > rgb_max {
		rgb_max = g
	}
	if b > rgb_max {
		rgb_max = b
	}
	if rgb_max == r {
		h = 60 * (g - b)
		if h < 0 {
			h += 360
		}
	} else if rgb_max == g {
		h = 120 + 60*(b-r)
	} else /* rgb_max == rgb.b */ {
		h = 240 + 60*(r-g)
	}
	return h
}

// Convert RGB to HSV
func HSV(cr color.RGBA) (uint8, uint8, uint8) {
	var hue, sat, val uint8
	rgb_min := Min(cr.R, cr.G, cr.B)
	rgb_max := Max(cr.R, cr.G, cr.B)

	val = rgb_max
	if val == 0 {
		hue = 0
		sat = 0
		return hue, sat, val
	}

	sat = 255 * (rgb_max - rgb_min) / val
	if sat == 0 {
		hue = 0
		return hue, sat, val
	}

	span := (rgb_max - rgb_min)
	if rgb_max == cr.R {
		hue = 43 * (cr.G - cr.B) / span
	} else if rgb_max == cr.G {
		hue = 85 + 43*(cr.B-cr.R)/span
	} else { /* rgb_max == cr.B */
		hue = 171 + 43*(cr.R-cr.G)/span
	}

	return hue, sat, val
}

// Separate an image into three colors, with a given threshold
func Separate(inImage image.Image, color1, color2, color3 color.RGBA, thresh uint8, t float64) []image.Image {
	if verbose {
		fmt.Print("Separating...")
	}
	var (
		rect       image.Rectangle = inImage.Bounds()
		cr         color.RGBA
		r, g, b, a uint8
		h, s       float64
	)
	images := make([]image.Image, 3) // 3 is the number of images
	newImage1 := image.NewRGBA(image.Rect(0, 0, rect.Max.X-rect.Min.X, rect.Max.Y-rect.Min.Y))
	newImage2 := image.NewRGBA(image.Rect(0, 0, rect.Max.X-rect.Min.X, rect.Max.Y-rect.Min.Y))
	newImage3 := image.NewRGBA(image.Rect(0, 0, rect.Max.X-rect.Min.X, rect.Max.Y-rect.Min.Y))
	hue1, _, s1 := HLS(float64(color1.R)/255.0, float64(color1.G)/255.0, float64(color1.B)/255.0)
	hue2, _, s2 := HLS(float64(color2.R)/255.0, float64(color2.G)/255.0, float64(color2.B)/255.0)
	hue3, _, s3 := HLS(float64(color3.R)/255.0, float64(color3.G)/255.0, float64(color3.B)/255.0)
	for y := rect.Min.Y; y < rect.Max.Y; y++ {
		for x := rect.Min.X; x < rect.Max.X; x++ {
			// get the rgba color
			// cr = inImage.At(x, y).(image.RGBAColor)
			cr = color.RGBAModel.Convert(inImage.At(x, y)).(color.RGBA)
			r = 0
			g = 0
			b = 0
			a = 255
			h, _, s = HLS(float64(cr.R)/255.0, float64(cr.G)/255.0, float64(cr.B)/255.0)
			// Find the closest color of the three, measured in hue and saturation
			if ((Fabs(h-hue1) < Fabs(h-hue2)) && (Fabs(h-hue1) < Fabs(h-hue3))) ||
				((Fabs(s-s1) < Fabs(s-s2)) && (Fabs(s-s1) < Fabs(s-s3))) {
				// Only add if the color is close enough
				if Abs(color1.R-cr.R) < thresh || Abs(color1.G-cr.G) < thresh || Abs(color1.B-cr.B) < thresh {
					r = color1.R
					g = color1.G
					b = color1.B
					newImage1.Set(x-rect.Min.X, y-rect.Min.Y, color.RGBA{r, g, b, a})
				}
			} else if ((Fabs(h-hue2) < Fabs(h-hue1)) && (Fabs(h-hue2) < Fabs(h-hue3))) ||
				((Fabs(s-s2) < Fabs(s-s1)) && (Fabs(s-s2) < Fabs(s-s3))) {
				// Only add if the color is close enough
				if Abs(color2.R-cr.R) < thresh || Abs(color2.G-cr.G) < thresh || Abs(color2.B-cr.B) < thresh {
					r = color2.R
					g = color2.G
					b = color2.B
					newImage2.Set(x-rect.Min.X, y-rect.Min.Y, color.RGBA{r, g, b, a})
				}
			} else if ((Fabs(h-hue3) < Fabs(h-hue1)) && (Fabs(h-hue3) < Fabs(h-hue2))) ||
				((Fabs(s-s3) < Fabs(s-s1)) && (Fabs(s-s3) < Fabs(s-s2))) {
				if Abs(color3.R-cr.R) < thresh || Abs(color3.G-cr.G) < thresh || Abs(color3.B-cr.B) < thresh {
					r = color3.R
					g = color3.G
					b = color3.B
					newImage3.Set(x-rect.Min.X, y-rect.Min.Y, color.RGBA{r, g, b, a})
				}
			}
		}
	}
	images[0] = newImage1
	images[1] = newImage2
	images[2] = newImage3
	if verbose {
		fmt.Println("done")
	}
	return images
}

// Smallest of three numbers
func Min(a, b, c uint8) uint8 {
	if (a < b) && (a < c) {
		return a
	} else if (b < a) && (b < c) {
		return b
	}
	return c
}

// Largest of three numbers
func Max(a, b, c uint8) uint8 {
	if (a >= b) && (a >= c) {
		return a
	} else if (b >= a) && (b >= c) {
		return b
	}
	return c
}

// Smallest of three floats
func Fmin(a, b, c float64) float64 {
	return math.Min(math.Min(a, b), c)
}

// Largest of three floats
func Fmax(a, b, c float64) float64 {
	return math.Max(math.Max(a, b), c)
}

// Convert RGB to HLS
func HLS(r, g, b float64) (float64, float64, float64) {
	// Ported from Python colorsys
	var h, l, s float64
	maxc := Fmax(r, g, b)
	minc := Fmin(r, g, b)
	l = (minc + maxc) / 2.0
	if minc == maxc {
		return 0.0, l, 0.0
	}
	span := (maxc - minc)
	if l <= 0.5 {
		s = span / (maxc + minc)
	} else {
		s = span / (2.0 - maxc - minc)
	}
	rc := (maxc - r) / span
	gc := (maxc - g) / span
	bc := (maxc - b) / span
	if r == maxc {
		h = bc - gc
	} else if g == maxc {
		h = 2.0 + rc - bc
	} else {
		h = 4.0 + gc - rc
	}
	h = math.Mod((h / 6.0), 1.0)
	return h, l, s
}

// Ported from Python colorsys
func _v(m1, m2, hue float64) float64 {
	ONE_SIXTH := 1.0 / 6.0
	TWO_THIRD := 2.0 / 3.0
	hue = math.Mod(hue, 1.0)
	if hue < ONE_SIXTH {
		return m1 + (m2-m1)*hue*6.0
	}
	if hue < 0.5 {
		return m2
	}
	if hue < TWO_THIRD {
		return m1 + (m2-m1)*(TWO_THIRD-hue)*6.0
	}
	return m1
}

// Convert a HLS color to RGB
func HLS_to_RGB(h, l, s float64) (float64, float64, float64) {
	// Ported from Python colorsys
	ONE_THIRD := 1.0 / 3.0
	if s == 0.0 {
		return l, l, l
	}
	var m2 float64
	if l <= 0.5 {
		m2 = l * (1.0 + s)
	} else {
		m2 = l + s - (l * s)
	}
	m1 := 2.0*l - m2
	return _v(m1, m2, h+ONE_THIRD), _v(m1, m2, h), _v(m1, m2, h-ONE_THIRD)
}

// Mix two RGB colors, a bit like how paint mixes
func PaintMix(c1, c2 color.RGBA) color.RGBA {
	// Thanks to Mark Ransom via stackoverflow

	// The less pi-presition, the greener the mix between blue and yellow
	// Using math.Pi gives a completely different result, for some reason
	//pi := math.Pi
	//pi := 3.141592653589793
	pi := 3.141592653
	//pi := 3.1415

	h1, l1, s1 := HLS(float64(c1.R)/255.0, float64(c1.G)/255.0, float64(c1.B)/255.0)
	h2, l2, s2 := HLS(float64(c2.R)/255.0, float64(c2.G)/255.0, float64(c2.B)/255.0)
	h := 0.0
	s := 0.5 * (s1 + s2)
	l := 0.5 * (l1 + l2)
	x := math.Cos(2.0*pi*h1) + math.Cos(2.0*pi*h2)
	y := math.Sin(2.0*pi*h1) + math.Sin(2.0*pi*h2)
	if (x != 0.0) || (y != 0.0) {
		h = math.Atan2(y, x) / (2.0 * pi)
	} else {
		s = 0.0
	}
	r, g, b := HLS_to_RGB(h, l, s)
	return color.RGBA{uint8(r * 255.0), uint8(g * 255.0), uint8(b * 255.0), 255}
}
