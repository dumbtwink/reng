package main

import (
	"fmt"
)

var point = [8][3]float64{{100, 100, 100}, {100, 400, 100}, {400, 400, 100}, {400, 100, 100}, {100, 100, 400}, {100, 400, 400}, {400, 400, 400}, {400, 100, 400}}
var face1 = polygon{[64][3]float64{point[0], point[1], point[2], point[3]}, 4}
var face2 = polygon{[64][3]float64{point[4], point[5], point[6], point[7]}, 4}
var face3 = polygon{[64][3]float64{point[0], point[3], point[7], point[4]}, 4}
var face4 = polygon{[64][3]float64{point[1], point[2], point[6], point[5]}, 4}
var face5 = polygon{[64][3]float64{point[0], point[1], point[5], point[4]}, 4}
var face6 = polygon{[64][3]float64{point[3], point[2], point[6], point[7]}, 4}

func main() {
	for i := float64(0); i < 1000; i++ {
		img := CreateCanvas(500, 500)
		DrawPolygon(img, face1, 255, 0, 0, 255)
		DrawPolygon(img, face2, 255, 255, 0, 255)
		DrawPolygon(img, face3, 0, 255, 255, 255)
		DrawPolygon(img, face4, 0, 0, 255, 255)
		DrawPolygon(img, face5, 255, 0, 255, 255)
		DrawPolygon(img, face6, 0, 255, 0, 255)
		var name string
		name = fmt.Sprintf("output/%v.png", i)
		SaveImage(img, name)
		face1 = face1.rotate(000.1, 1, true, 250, 250, 250)
		face2 = face2.rotate(000.1, 1, true, 250, 250, 250)
		face3 = face3.rotate(000.1, 1, true, 250, 250, 250)
		face4 = face4.rotate(000.1, 1, true, 250, 250, 250)
		face5 = face5.rotate(000.1, 1, true, 250, 250, 250)
		face6 = face6.rotate(000.1, 1, true, 250, 250, 250)
		face1 = face1.rotate(0000.1, 2, true, 250, 250, 250)
		face2 = face2.rotate(0000.1, 2, true, 250, 250, 250)
		face3 = face3.rotate(0000.1, 2, true, 250, 250, 250)
		face4 = face4.rotate(0000.1, 2, true, 250, 250, 250)
		face5 = face5.rotate(0000.1, 2, true, 250, 250, 250)
		face6 = face6.rotate(0000.1, 2, true, 250, 250, 250)
	}
}
