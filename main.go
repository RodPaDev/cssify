package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"math"
	"os"

	tColor "github.com/TwiN/go-color"
	"github.com/cheggaaa/pb/v3"
	"github.com/julvo/htmlgo"
	attr "github.com/julvo/htmlgo/attributes"
)

type CssImage struct {
	colors []string
	width  int
	height int
	name   string
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

	loadedImage := imageLoader(path)

	img, _, err := image.Decode(loadedImage)
	if err != nil {
		log.Fatal(err)
	}

	dst := image.NewRGBA(img.Bounds())
	draw.Draw(dst, img.Bounds(), img, img.Bounds().Min, draw.Src)
	img = nil

	cssImage := CssImage{
		colors: getCssColors(dst),
		width:  dst.Bounds().Dx(),
		height: dst.Bounds().Dy(),
		name:   "cssify",
	}

	cssify(&cssImage)

}

func getCssColors(img image.Image) []string {

	bounds := img.Bounds()

	size := bounds.Dx() * bounds.Dy()
	prettyPrint(fmt.Sprintf("CSS-ifying Image: %vx%v", bounds.Dx(), bounds.Dy()), true)

	var hexSlice = make([]string, 0, size)
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

func rgbaToCssColor(rgba *color.RGBA, isRgb bool) string {
	if isRgb {
		alpha := math.Floor((float64(rgba.A)/255)*10) / 10
		return fmt.Sprintf("rgba(%v, %v, %v, %v)", rgba.R, rgba.G, rgba.B, alpha)
	}
	return fmt.Sprintf("#%02x%02x%02x%02x", rgba.R, rgba.G, rgba.B, rgba.A)
}

func cssify(cssImage *CssImage) {

	sliceLength := len(cssImage.colors)

	bar := progressBar("CSS-ifying image:", sliceLength)

	file := createFile("./out/index.html")
	defer file.Close()

	writer := bufio.NewWriter(file)

	for _, cssColor := range cssImage.colors {
		style := fmt.Sprintf("background-color: %v;", cssColor)
		div := htmlgo.Div(htmlgo.Attr(attr.Style_(style), attr.Class_("pixel")))

		writer.WriteString(string(div))
		writer.Flush()
		bar.Increment()
	}

	bar.Finish()
	prettyPrint("HTML Generated", true)
}

func imageLoader(path string) *os.File {
	file, err := os.Open(path)

	if err != nil {
		fmt.Println("Error: File could not be opened")
		fmt.Println(err)
		os.Exit(1)
	}
	return file
}

func createFile(path string) *os.File {
	prettyPrint("HTML File Created", true)

	err := os.MkdirAll("./out/", 0700)
	if err != nil {
		fmt.Println("Error: Directory could not be created")
		fmt.Println(err)
		os.Exit(1)
	}

	file, err := os.Create(path)

	if err != nil {
		fmt.Println("Error: index.html could not be created")
		fmt.Println(err)
		os.Exit(1)
	}
	return file
}

func progressBar(str string, total int) *pb.ProgressBar {
	tmpl := fmt.Sprintf(`{{ cyan "%v" }} {{ bar . (white "")  "█" "▓" "░" "░" (white "")}} {{percent . }} {{rtime .}}`, str)
	bar := pb.ProgressBarTemplate(tmpl).Start(total)
	return bar
}

func prettyPrint(str string, newLine bool) {
	fmt.Print("\033[2K\r")
	if newLine {
		fmt.Println("\033[32m" + str)
		return
	}
	fmt.Print("\033[36m" + str)
}
