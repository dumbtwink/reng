package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

type polygon struct {
	vertices [64][3]float64
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
	if math.Floor(bx)-math.Floor(ax) == 0 {
		if ay < by {
			for i := ay; i < by; i++ {
				img.(*image.RGBA).Set(int(math.Ceil(ax)), int(math.Ceil(i)), color.RGBA{R: R, G: G, B: B, A: A})
			}
		} else {
			for i := ay; i > by; i-- {
				img.(*image.RGBA).Set(int(math.Ceil(ax)), int(math.Ceil(i)), color.RGBA{R: R, G: G, B: B, A: A})
			}
		}
	}
	if ax < bx {
		for i := ax; i < bx; i++ {
			if ay < by {
				img.(*image.RGBA).Set(int(math.Ceil(i)), int(float64(cnt)*incl+ay), color.RGBA{R: R, G: G, B: B, A: A})
				oct = cnt
				cnt++
				for j := float64(oct) * incl; j < float64(cnt)*incl; j += 0.1 {
					img.(*image.RGBA).Set(int(math.Ceil(i)), int(math.Ceil(j)+ay), color.RGBA{R: R, G: G, B: B, A: A})
					if j+ay > by {
						break
					}
				}
			} else {
				img.(*image.RGBA).Set(int(math.Ceil(i)), int(float64(cnt)*incl+ay), color.RGBA{R: R, G: G, B: B, A: A})
				oct = cnt
				cnt++
				for j := float64(oct) * incl; j > float64(cnt)*incl; j -= 0.1 {
					img.(*image.RGBA).Set(int(math.Ceil(i)), int(math.Ceil(j)+ay), color.RGBA{R: R, G: G, B: B, A: A})
					if j+ay < by {
						break
					}
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
					if j+ay > by {
						break
					}
				}
			} else {
				img.(*image.RGBA).Set(int(math.Ceil(i)), int(float64(cnt)*incl+ay), color.RGBA{R: R, G: G, B: B, A: A})
				oct = cnt
				cnt--
				for j := float64(oct) * incl; j > float64(cnt)*incl; j -= 0.1 {
					img.(*image.RGBA).Set(int(math.Ceil(i)), int(math.Ceil(j)+ay), color.RGBA{R: R, G: G, B: B, A: A})
					if j+ay < by {
						break
					}
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

func (poly polygon) rotate(radians float64, axis int, override bool, x, y, z float64) polygon {
	var output polygon
	var xsum, ysum, avarageX, avarageY float64
	var slotx, sloty int
	switch axis {
	case 0:
		slotx, sloty = 0, 1
		avarageX = x
		avarageY = y
		break
	case 1:
		slotx, sloty = 0, 2
		avarageX = x
		avarageY = z
		break
	case 2:
		slotx, sloty = 1, 2
		avarageX = y
		avarageY = z
		break
	}
	if !override {
		for i := 0; i < poly.amount; i++ {
			xsum += poly.vertices[i][slotx]
			ysum += poly.vertices[i][sloty]
		}
		avarageX = (xsum) / float64(poly.amount)
		avarageY = (ysum) / float64(poly.amount)
	}
	for i := 0; i < poly.amount; i++ {
		tmpangle := math.Atan((poly.vertices[i][sloty] - avarageY) / (poly.vertices[i][slotx] - avarageX))
		tmpradius := (poly.vertices[i][slotx] - avarageX) / math.Cos(tmpangle)
		relativeX := math.Cos(radians+tmpangle) * tmpradius
		relativeY := math.Sin(radians+tmpangle) * tmpradius
		switch axis {
		case 0:
			output.vertices[i] = [3]float64{relativeX + avarageX, relativeY + avarageY, poly.vertices[i][2]}
			break
		case 1:
			output.vertices[i] = [3]float64{relativeX + avarageX, poly.vertices[i][1], relativeY + avarageY}
			break
		case 2:
			output.vertices[i] = [3]float64{poly.vertices[i][0], relativeX + avarageX, relativeY + avarageY}
			break
		}
	}
	output.amount = poly.amount
	return output
}
