package main

import (
	"fmt"
	"math"
)

var poly = polygon{[64][2]float64{{50, 300}, {213, 100}, {430, 350}, {232, 459}}, 4}

func main() {
	for i := float64(0); i < 100; i++ {
		img := CreateCanvas(500, 500)
		poly = poly.rotate(000.1)
		DrawPolygon(img, poly, uint8(math.Ceil(255/(i/10))), 0, uint8(math.Ceil(2.55*i)), 255)
		var name string
		name = fmt.Sprintf("output/%v.png", i)
		SaveImage(img, name)
	}
}
