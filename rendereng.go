package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

func CreateCanvas(x, y int) image.Image {
	img := image.NewRGBA(image.Rectangle{Min: image.Point{}, Max: image.Point{X: x, Y: y}})
	return img
}

func SaveImage(img image.Image) {
	f, _ := os.Create("img.png")
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}(f)
	err := png.Encode(f, img)
	if err != nil {
		panic(err)
		return
	}
}

func DrawCircle(img image.Image, r float64, relativeX, relativeY int, R, G, B, A uint8) {
	var x, y float64
	for i := 0.0; i < 2*(r*math.Pi); i += 0.1 {
		x = math.Cos(i) * r
		y = math.Sin(i) * r
		img.(*image.RGBA).Set(int(math.Ceil(x))+relativeX, int(math.Ceil(y))+relativeY, color.RGBA{R: R, G: G, B: B, A: A})
	}
}

func DrawLine(img image.Image, ax, ay, bx, by float64, R, G, B, A uint8) {
	incl := (by - ay) / (bx - ax)
	fmt.Printf("AX: %v, AY: %v, BX: %v, BY: %v\n", ax, ay, bx, by)
	fmt.Printf("(%v - %v) / (%v - %v) = %v\n", by, ay, bx, ax, incl)
	fmt.Printf("Inclination %v\n", incl)
	oct := 1
	cnt := 1
	if ax < bx {
		for i := ax; i < bx; i++ {
			img.(*image.RGBA).Set(int(math.Ceil(i)), int(float64(cnt)*incl+ay), color.RGBA{R: R, G: G, B: B, A: A})
			fmt.Printf("(clr: %v, %v, %v, %v) X: %v, Y: %v\n", R, G, B, A, math.Ceil(i), int(float64(cnt)*incl+ax))
			oct = cnt
			cnt++
			if ay < by {
				for j := float64(oct) * incl; j < float64(cnt)*incl; j += 0.1 {
					img.(*image.RGBA).Set(int(math.Ceil(i)), int(math.Ceil(j)+ay), color.RGBA{R: R, G: G, B: B, A: A})

				}
			} else {
				for j := float64(oct) * incl; j > float64(cnt)*incl; j -= 0.1 {
					img.(*image.RGBA).Set(int(math.Ceil(i)), int(math.Ceil(j)+ay), color.RGBA{R: R, G: G, B: B, A: A})

				}
			}
		}
	} else {
		for i := ax; i > bx; i-- {
			img.(*image.RGBA).Set(int(math.Ceil(i)), int(float64(cnt)*incl+ay), color.RGBA{R: 255, A: 255})
			fmt.Printf("(clr: %v, %v, %v, %v) X: %v, Y: %v\n", R, G, B, A, math.Ceil(i), int(float64(cnt)*incl+ax))
			oct = cnt
			cnt++
			if ay < by {
				for j := float64(oct) * incl; j < float64(cnt)*incl; j += 0.1 {
					img.(*image.RGBA).Set(int(math.Ceil(i)), int(math.Ceil(j)+ay), color.RGBA{R: R, G: G, B: B, A: A})

				}
			} else {
				for j := float64(oct) * incl; j > float64(cnt)*incl; j -= 0.1 {
					img.(*image.RGBA).Set(int(math.Ceil(i)), int(math.Ceil(j)+ay), color.RGBA{R: R, G: G, B: B, A: A})

				}
			}
		}
	}
}
