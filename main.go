package main

import (
	"log"
	"math"
	"os"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	_ "image/jpeg"
	_ "image/png"
	tColor "github.com/TwiN/go-color"
)

type Cssify struct {
	Pixels []*Pixel
	Width  int
	Height int
}

type Pixel struct {
	Color string
}

func main() {

	if len(os.Args) <= 1 {
		fmt.Println(tColor.White + "CSSify is a tool to convert a image to HTML & CSS")
		fmt.Print("\nUsage:\n\n")
		fmt.Print("\tcssify <image path> [arguments]\n\n")
		fmt.Println("Flags:")
		fmt.Println("  -h, -hex\tHexadecimal with transparency DEFAULT")
		fmt.Println("  -r, -rgb\tRed Green Blue with transparency")
		fmt.Print("\n")
		os.Exit(0)
	}

	prettyPrint("---- CSSify v0.1 ----", true)
	path := os.Args[1]

	loadedImage := fileOpen(path)

	img, _, err := image.Decode(loadedImage)
	if err != nil {
		log.Fatal(err)
	}

	dst := image.NewRGBA(img.Bounds())
	draw.Draw(dst, img.Bounds(), img, img.Bounds().Min, draw.Src)
	img = nil

	cssImage := Cssify{
		Pixels: getCssColors(dst),
		Width:  dst.Bounds().Dx(),
		Height: dst.Bounds().Dy(),
	}

	cssify(&cssImage)

}

func getCssColors(img image.Image) []*Pixel {

	bounds := img.Bounds()

	size := bounds.Dx() * bounds.Dy()
	prettyPrint(fmt.Sprintf("CSS-ifying Image: %vx%v", bounds.Dx(), bounds.Dy()), true)

	var hexSlice = make([]*Pixel, 0, size)

	bar := progressBar("Extracting color:", size)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			rgba := img.At(x, y).(color.RGBA)
			hexSlice = append(hexSlice, rgbaToCssColor(&rgba, false))
			bar.Increment()
		}
	}

	bar.Finish()

	return hexSlice
}

func rgbaToCssColor(rgba *color.RGBA, isRgb bool) *Pixel {
	if isRgb {
		alpha := math.Floor((float64(rgba.A)/255)*10) / 10
		return &Pixel{Color: fmt.Sprintf("rgba(%v, %v, %v, %v)", rgba.R, rgba.G, rgba.B, alpha)}
	}

	return &Pixel{Color: fmt.Sprintf("#%02x%02x%02x%02x", rgba.R, rgba.G, rgba.B, rgba.A)}
}

func cssify(cssImage *Cssify) {
	file := createFile()
	defer file.Close()

	htmlGenerator(Cssify{Pixels: cssImage.Pixels, Width: cssImage.Width, Height: cssImage.Height})
}
