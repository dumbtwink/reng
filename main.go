package main

import "math"

var poly = [6]float64{100, 100, 200, 500, 300, 300}

func main() {
	img := CreateCanvas(500, 500)
	DrawLine(img, poly[0], poly[1], poly[2], poly[3], 255, 0, 0, 255)
	DrawLine(img, poly[2], poly[3], poly[4], poly[5], 255, 0, 255, 255)
	DrawLine(img, poly[0], poly[1], poly[4], poly[5], 255, 255, 0, 255)
	DrawCircle(img, 5, int(math.Ceil(poly[0])), int(math.Ceil(poly[1])), 0, 255, 0, 255)
	DrawCircle(img, 5, int(math.Ceil(poly[2])), int(math.Ceil(poly[3])), 0, 255, 0, 255)
	DrawCircle(img, 5, int(math.Ceil(poly[4])), int(math.Ceil(poly[5])), 0, 255, 0, 255)
	SaveImage(img)
}
