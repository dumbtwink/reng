package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

type polygon struct {
	vertices [64][2]float64
	amount   int
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
	for i := 0; i < poly.amount-1; i++ {
		DrawLine(img, poly.vertices[i][0], poly.vertices[i][1], poly.vertices[i+1][0], poly.vertices[i+1][1], R, G, B, A)
	}
	DrawLine(img, poly.vertices[0][0], poly.vertices[0][1], poly.vertices[poly.amount-1][0], poly.vertices[poly.amount-1][1], R, G, B, A)
}

func (poly polygon) rotate(radians float64) polygon {
	var output polygon
	var xsum, ysum float64
	for i := 0; i < poly.amount; i++ {
		xsum += poly.vertices[i][0]
		ysum += +poly.vertices[i][1]
	}
	avarageX := (xsum) / float64(poly.amount)
	avarageY := (ysum) / float64(poly.amount)
	for i := 0; i < poly.amount; i++ {
		tmpangle := math.Atan((poly.vertices[i][1] - avarageY) / (poly.vertices[i][0] - avarageX))
		tmpradius := (poly.vertices[i][0] - avarageX) / math.Cos(tmpangle)
		relativeX := math.Cos(radians+tmpangle) * tmpradius
		relativeY := math.Sin(radians+tmpangle) * tmpradius
		output.vertices[i] = [2]float64{relativeX + avarageX, relativeY + avarageY}
	}
	output.amount = poly.amount
	return output
}
