package main

import "fmt"

var poly = polygon{p1: [2]float64{50, 300}, p2: [2]float64{213, 100}, p3: [2]float64{430, 350}}

func main() {
	for i := float64(0); i < 100; i++ {
		img := CreateCanvas(500, 500)
		DrawPolygon(img, poly, 0, 255, 0, 255)
		var name string
		poly = poly.rotate(000000000000000000000000000000000000000.1)
		name = fmt.Sprintf("output/%v.png", i)
		SaveImage(img, name)
	}
}
