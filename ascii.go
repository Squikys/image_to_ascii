package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

func main() {
	args := os.Args
	imageFile, err := os.Open(string(args[1]))
	if err != nil {
		panic(err)
	}
	defer imageFile.Close()
	var loadedImage image.Image
	if strings.Contains(string(args[1]), ".jpeg") || strings.Contains(string(args[1]), ".jpg") {
		loadedImage, err = jpeg.Decode(imageFile)
		if err != nil {
			panic(err)
		}
	} else if strings.Contains(string(args[1]), ".png") {
		loadedImage, err = png.Decode(imageFile)
		if err != nil {
			panic(err)
		}
	}
	bounds := loadedImage.Bounds()
	height := bounds.Dy()
	width := bounds.Dx()

	x_temp := 0
	y_temp := 0
	x_ratio := 50
	y_ratio := x_ratio * 1 / 2 * height / width

	for y := 0; y < y_ratio; y++ {
		for x := 0; x < x_ratio; x++ {
			average := 0
			size := 0
			for j := y_temp; j < y_temp+height/y_ratio; j++ {
				for i := x_temp; i < x_temp+width/x_ratio; i++ {
					pixel := loadedImage.At(i, j)

					originalcolor := color.RGBAModel.Convert(pixel).(color.RGBA)
					grey := ((float64(originalcolor.R) + float64(originalcolor.G) + float64(originalcolor.B)) / 3)
					average = average + int(grey)
					size++
				}
			}
			x_temp = x_temp + width/x_ratio
			total := average

			total = total / size
			mainpix := total
			if mainpix <= 32 {
				fmt.Printf(" ")
			} else if mainpix < 64 && mainpix > 32 {
				fmt.Printf("_")
			} else if mainpix < 96 && mainpix > 64 {
				fmt.Printf("+")
			} else if mainpix < 128 && mainpix > 96 {
				fmt.Printf("!")
			} else if mainpix < 160 && mainpix > 128 {
				fmt.Printf("?")
			} else if mainpix < 192 && mainpix > 160 {
				fmt.Printf("#")
			} else if mainpix < 224 && mainpix > 192 {
				fmt.Printf("$")
			} else if mainpix < 256 && mainpix > 224 {
				fmt.Printf("@")
			}

		}
		y_temp = y_temp + height/y_ratio
		x_temp = 0
		fmt.Print("\n")
	}

}
