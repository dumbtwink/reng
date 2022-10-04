package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
	"sort"
)

type polygon struct {
	vertices [64][3]float64
	amount   int
	clr      color.Color
}

type mesh struct {
	polygons [64]polygon
	amount   int
}

type canvas struct {
	x, y, z int
}

var viewport canvas

func CreateViewport(x, y, z int) image.Image {
	img := image.NewRGBA(image.Rectangle{Min: image.Point{}, Max: image.Point{X: x, Y: y}})
	viewport.x = x
	viewport.y = y
	viewport.z = z
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

func DrawCircle(img image.Image, r, x, y float64, clr color.Color) {
	var relativex, relativey float64
	for i := 0.0; i < 2*(r*math.Pi); i += 0.1 {
		relativex = math.Cos(i) * r
		relativey = math.Sin(i) * r
		img.(*image.RGBA).Set(int(math.Ceil(relativex+x)), int(math.Ceil(relativey+y)), clr)
	}
}

func DrawLine(img image.Image, ax, ay, bx, by float64, clr color.Color) {
	incl := (by - ay) / (bx - ax)
	var oct, cnt int
	if math.Floor(bx)-math.Floor(ax) == 0 {
		if ay < by {
			for i := ay; i < by; i++ {
				img.(*image.RGBA).Set(int(math.Ceil(ax)), int(math.Ceil(i)), clr)
			}
		} else {
			for i := ay; i > by; i-- {
				img.(*image.RGBA).Set(int(math.Ceil(ax)), int(math.Ceil(i)), clr)
			}
		}
	}
	if ax < bx {
		for i := ax; i < bx; i++ {
			if ay < by {
				img.(*image.RGBA).Set(int(math.Ceil(i)), int(float64(cnt)*incl+ay), clr)
				oct = cnt
				cnt++
				for j := float64(oct) * incl; j < float64(cnt)*incl; j += 0.1 {
					img.(*image.RGBA).Set(int(math.Ceil(i)), int(math.Ceil(j)+ay), clr)
					if j+ay > by {
						break
					}
				}
			} else {
				img.(*image.RGBA).Set(int(math.Ceil(i)), int(float64(cnt)*incl+ay), clr)
				oct = cnt
				cnt++
				for j := float64(oct) * incl; j > float64(cnt)*incl; j -= 0.1 {
					img.(*image.RGBA).Set(int(math.Ceil(i)), int(math.Ceil(j)+ay), clr)
					if j+ay < by {
						break
					}
				}
			}
		}
	} else {
		for i := ax; i > bx; i-- {
			if ay < by {
				img.(*image.RGBA).Set(int(math.Ceil(i)), int(float64(cnt)*incl+ay), clr)
				oct = cnt
				cnt--
				for j := float64(oct) * incl; j < float64(cnt)*incl; j += 0.1 {
					img.(*image.RGBA).Set(int(math.Ceil(i)), int(math.Ceil(j)+ay), clr)
					if j+ay > by {
						break
					}
				}
			} else {
				img.(*image.RGBA).Set(int(math.Ceil(i)), int(float64(cnt)*incl+ay), clr)
				oct = cnt
				cnt--
				for j := float64(oct) * incl; j > float64(cnt)*incl; j -= 0.1 {
					img.(*image.RGBA).Set(int(math.Ceil(i)), int(math.Ceil(j)+ay), clr)
					if j+ay < by {
						break
					}
				}
			}
		}
	}
}

func Trace(ax, ay, bx, by float64) [][2]float64 {
	var coords [][2]float64
	incl := (by - ay) / (bx - ax)
	var oct, cnt int
	if math.Floor(bx)-math.Floor(ax) == 0 {
		if ay < by {
			for i := ay; i < by; i++ {
				coords = append(coords, [2]float64{ax, i})
			}
		} else {
			for i := ay; i > by; i-- {
				coords = append(coords, [2]float64{ax, i})
			}
		}
	}
	if ax < bx {
		for i := ax; i < bx; i++ {
			if ay < by {
				coords = append(coords, [2]float64{i, float64(cnt)*incl + ay})
				oct = cnt
				cnt++
				for j := float64(oct) * incl; j < float64(cnt)*incl; j += 0.1 {
					coords = append(coords, [2]float64{i, math.Ceil(j) + ay})
					if j+ay > by {
						break
					}
				}
			} else {
				coords = append(coords, [2]float64{i, float64(cnt)*incl + ay})
				oct = cnt
				cnt++
				for j := float64(oct) * incl; j > float64(cnt)*incl; j -= 0.1 {
					coords = append(coords, [2]float64{i, math.Ceil(j) + ay})
					if j+ay < by {
						break
					}
				}
			}
		}
	} else {
		for i := ax; i > bx; i-- {
			if ay < by {
				coords = append(coords, [2]float64{i, float64(cnt)*incl + ay})
				oct = cnt
				cnt--
				for j := float64(oct) * incl; j < float64(cnt)*incl; j += 0.1 {
					coords = append(coords, [2]float64{i, float64(cnt)*incl + ay})
					if j+ay > by {
						break
					}
				}
			} else {
				coords = append(coords, [2]float64{i, float64(cnt)*incl + ay})
				oct = cnt
				cnt--
				for j := float64(oct) * incl; j > float64(cnt)*incl; j -= 0.1 {
					coords = append(coords, [2]float64{i, float64(cnt)*incl + ay})
					if j+ay < by {
						break
					}
				}
			}
		}
	}
	return coords
}

func DrawPolygon(img image.Image, poly polygon, clr color.Color) {
	for i := 0; i < poly.amount-1; i++ {
		DrawLine(img, poly.vertices[i][0], poly.vertices[i][1], poly.vertices[i+1][0], poly.vertices[i+1][1], clr)
	}
	DrawLine(img, poly.vertices[0][0], poly.vertices[0][1], poly.vertices[poly.amount-1][0], poly.vertices[poly.amount-1][1], clr)
}

func DrawWireframe(img image.Image, obj mesh, clr color.Color) {
	newMesh := obj.SortByDepthmap()
	for i := 0; i < obj.amount; i++ {
		DrawPolygon(img, newMesh.polygons[i], clr)
	}
}

func DrawMesh(img image.Image, obj mesh) {
	newMesh := obj.SortByDepthmap()
	for i := obj.amount - 1; i >= 0; i -= 1{
		DrawFace(img, newMesh.polygons[i])
	}
}

func DrawFace(img image.Image, poly polygon) {
	for i := 0; i < poly.amount; i += 4 {
		invlineA := Trace(poly.vertices[i][0], poly.vertices[i][1], poly.vertices[i+1][0], poly.vertices[i+1][1])
		invlineB := Trace(poly.vertices[i+3][0], poly.vertices[i+3][1]+1, poly.vertices[i+2][0], poly.vertices[i+2][1])
		for j := 0; j < len(invlineA) && j < len(invlineB); j += 1 {
			DrawLine(img, invlineA[j][0], invlineA[j][1], invlineB[j][0], invlineB[j][1], poly.clr)
		}
	}

}

func (poly polygon) calcDepth() float64 {
	var cnt float64
	for i := 0; i < poly.amount; i++ {
		cnt += poly.vertices[i][2]
	}
	return cnt / float64(poly.amount)
}

func (obj mesh) genDepthmap() []int {
	var depthmap []int
	for i := 0; i < obj.amount; i++ {
		depthmap = append(depthmap, int(obj.polygons[i].calcDepth()))
	}
	return depthmap
}

func (obj mesh) SortByDepthmap() mesh {
	depthmap := obj.genDepthmap()
	var newMesh mesh
	sort.Ints(depthmap)
	for i := 0; i < obj.amount; i++ {
		newMesh.polygons[i].clr = obj.polygons[i].clr
		for j := viewport.z; j >= 0; j-- {
			if depthmap[i] == j {
				for k := 0; k < obj.amount; k++ {
					if int(obj.polygons[k].calcDepth()) == depthmap[i] {
						newMesh.polygons[i] = obj.polygons[k]
					}
				}
			}
		}
	}
	newMesh.amount = obj.amount
	return newMesh
}

func (poly polygon) rotate(radians float64, axis int, obj ...mesh) polygon {
	var output polygon
	var xsum, ysum, avarageX, avarageY float64
	var slotx, sloty int

	switch axis {
	case 0:
		slotx, sloty = 0, 1
		break
	case 1:
		slotx, sloty = 0, 2
		break
	case 2:
		slotx, sloty = 1, 2
		break
	}
	for i := 0; i < poly.amount; i++ {
		xsum += poly.vertices[i][slotx]
		ysum += poly.vertices[i][sloty]
	}
	avarageX = (xsum) / float64(poly.amount)
	avarageY = (ysum) / float64(poly.amount)
	if len(obj) != 0 {
		var meshPos = obj[0].position()
		avarageX = meshPos[slotx]
		avarageY = meshPos[sloty]
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
	output.clr = poly.clr
	output.amount = poly.amount
	return output
}

func (obj mesh) position() [3]float64 {
	var cord [3]float64
	for i := 0; i < 3; i++ {
		amount := 0
		for j := 0; j < obj.amount; j++ {
			for k := 0; k < obj.polygons[j].amount; k++ {
				cord[i] += obj.polygons[j].vertices[k][i]
				amount++
			}
		}
		cord[i] = cord[i] / float64(amount)
	}
	return [3]float64{cord[0], cord[1], cord[2]}
}

func (obj mesh) rotate(radians float64, axis int) mesh {
	var output mesh
	for i := 0; i < obj.amount; i++ {
		output.polygons[i] = obj.polygons[i].rotate(radians, axis, obj)
		output.polygons[i].clr = obj.polygons[i].clr
	}
	output.amount = obj.amount
	return output
}
