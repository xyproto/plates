package plates

import (
	"image"
	"image/color"
	"testing"
)

func TestRed(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 1, 1))
	img.Set(0, 0, color.RGBA{255, 0, 0, 255})
	newImage := Red(img)

	r, _, _, _ := newImage.At(0, 0).RGBA()
	if r != 65535 {
		t.Errorf("Expected 65535, got %d", r)
	}
}

func TestGreen(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 1, 1))
	img.Set(0, 0, color.RGBA{0, 255, 0, 255})
	newImage := Green(img)

	_, g, _, _ := newImage.At(0, 0).RGBA()
	if g != 65535 {
		t.Errorf("Expected 65535, got %d", g)
	}
}

func TestBlue(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 1, 1))
	img.Set(0, 0, color.RGBA{0, 0, 255, 255})
	newImage := Blue(img)

	_, _, b, _ := newImage.At(0, 0).RGBA()
	if b != 65535 {
		t.Errorf("Expected 65535, got %d", b)
	}
}

func TestHue(t *testing.T) {
	c := color.RGBA{255, 0, 0, 255}
	h := Hue(c)
	if h != 0 {
		t.Errorf("Expected 0, got %f", h)
	}
}

func TestHSV(t *testing.T) {
	c := color.RGBA{255, 0, 0, 255}
	h, s, v := HSV(c)
	if h != 0 || s != 0 || v != 255 {
		t.Errorf("Expected (0, 0, 255), got (%d, %d, %d)", h, s, v)
	}
}
