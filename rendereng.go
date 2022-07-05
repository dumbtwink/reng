package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

type polygon struct {
	p1 [2]float64
	p2 [2]float64
	p3 [2]float64
}

func CreateCanvas(x, y int) image.Image {
	img := image.NewRGBA(image.Rectangle{Min: image.Point{}, Max: image.Point{X: x, Y: y}})
	return img
}

func SaveImage(img image.Image, dir string) {
	f, _ := os.Create(dir)
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

func DrawCircle(img image.Image, r, x, y float64, R, G, B, A uint8) {
	var relativex, relativey float64
	for i := 0.0; i < 2*(r*math.Pi); i += 0.1 {
		relativex = math.Cos(i) * r
		relativey = math.Sin(i) * r
		img.(*image.RGBA).Set(int(math.Ceil(relativex+x)), int(math.Ceil(relativey+y)), color.RGBA{R: R, G: G, B: B, A: A})
	}
}

func DrawLine(img image.Image, ax, ay, bx, by float64, R, G, B, A uint8) {
	incl := (by - ay) / (bx - ax)
	var oct, cnt int
	if ax < bx {
		for i := ax; i < bx; i++ {
			if ay < by {
				img.(*image.RGBA).Set(int(math.Ceil(i)), int(float64(cnt)*incl+ay), color.RGBA{R: R, G: G, B: B, A: A})
				oct = cnt
				cnt++
				for j := float64(oct) * incl; j < float64(cnt)*incl; j += 0.1 {
					img.(*image.RGBA).Set(int(math.Ceil(i)), int(math.Ceil(j)+ay), color.RGBA{R: R, G: G, B: B, A: A})

				}
			} else {
				img.(*image.RGBA).Set(int(math.Ceil(i)), int(float64(cnt)*incl+ay), color.RGBA{R: R, G: G, B: B, A: A})
				oct = cnt
				cnt++
				for j := float64(oct) * incl; j > float64(cnt)*incl; j -= 0.1 {
					img.(*image.RGBA).Set(int(math.Ceil(i)), int(math.Ceil(j)+ay), color.RGBA{R: R, G: G, B: B, A: A})

				}
			}
		}
	} else {
		for i := ax; i > bx; i-- {
			if ay < by {
				img.(*image.RGBA).Set(int(math.Ceil(i)), int(float64(cnt)*incl+ay), color.RGBA{R: R, G: G, B: B, A: A})
				oct = cnt
				cnt--
				for j := float64(oct) * incl; j < float64(cnt)*incl; j += 0.1 {
					img.(*image.RGBA).Set(int(math.Ceil(i)), int(math.Ceil(j)+ay), color.RGBA{R: R, G: G, B: B, A: A})

				}
			} else {
				img.(*image.RGBA).Set(int(math.Ceil(i)), int(float64(cnt)*incl+ay), color.RGBA{R: R, G: G, B: B, A: A})
				oct = cnt
				cnt--
				for j := float64(oct) * incl; j > float64(cnt)*incl; j -= 0.1 {
					img.(*image.RGBA).Set(int(math.Ceil(i)), int(math.Ceil(j)+ay), color.RGBA{R: R, G: G, B: B, A: A})

				}
			}
		}
	}
}

func DrawPolygon(img image.Image, poly polygon, R, G, B, A uint8) {
	DrawLine(img, poly.p1[0], poly.p1[1], poly.p2[0], poly.p2[1], R, G, B, A)
	DrawLine(img, poly.p2[0], poly.p2[1], poly.p3[0], poly.p3[1], R, G, B, A)
	DrawLine(img, poly.p1[0], poly.p1[1], poly.p3[0], poly.p3[1], R, G, B, A)
}

func (poly polygon) rotate(radians float64) polygon {
	var output polygon
	avarageX := (poly.p1[0] + poly.p2[0] + poly.p3[0]) / 3
	avarageY := (poly.p1[1] + poly.p2[1] + poly.p3[1]) / 3

	p1angle := math.Atan((poly.p1[1] - avarageY) / (poly.p1[0] - avarageX))
	p1radius := (poly.p1[0] - avarageX) / math.Cos(p1angle)
	relativeX := math.Cos(radians+p1angle) * p1radius
	relativeY := math.Sin(radians+p1angle) * p1radius
	output.p1 = [2]float64{relativeX + avarageX, relativeY + avarageY}

	p2angle := math.Atan((poly.p2[1] - avarageY) / (poly.p2[0] - avarageX))
	p2radius := (poly.p2[0] - avarageX) / math.Cos(p2angle)
	relativeX = math.Cos(radians+p2angle) * p2radius
	relativeY = math.Sin(radians+p2angle) * p2radius
	output.p2 = [2]float64{relativeX + avarageX, relativeY + avarageY}

	p3angle := math.Atan((poly.p3[1] - avarageY) / (poly.p3[0] - avarageX))
	p3radius := (poly.p3[0] - avarageX) / math.Cos(p3angle)
	relativeX = math.Cos(radians+p3angle) * p3radius
	relativeY = math.Sin(radians+p3angle) * p3radius
	output.p3 = [2]float64{relativeX + avarageX, relativeY + avarageY}
	return output
}
