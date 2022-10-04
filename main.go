package main

exi

var point = [8][3]float64{{100, 100, 100}, {100, 400, 100}, {400, 400, 100}, {400, 100, 100}, {100, 100, 400}, {100, 400, 400}, {400, 400, 400}, {400, 100, 400}}

var cube = mesh{polygons: [64]polygon{{vertices: [64][3]float64{point[0], point[1], point[2], point[3]}, amount: 4, clr: color.RGBA{R: 255, A: 255}},
	{vertices: [64][3]float64{point[4], point[5], point[6], point[7]}, amount: 4, clr: color.RGBA{B: 255, A: 255}},
	{vertices: [64][3]float64{point[0], point[3], point[7], point[4]}, amount: 4, clr: color.RGBA{G: 255, A: 255}},
	{vertices: [64][3]float64{point[1], point[2], point[6], point[5]}, amount: 4, clr: color.RGBA{R: 255, G: 255, A: 255}},
	{vertices: [64][3]float64{point[0], point[1], point[5], point[4]}, amount: 4, clr: color.RGBA{G: 255, B: 255, A: 255}},
	{vertices: [64][3]float64{point[3], point[2], point[6], point[7]}, amount: 4, clr: color.RGBA{B: 255, R: 255, A: 255}}}, amount: 6}

func main() {
	donecheck := make(chan bool)
	rendertime := make(chan int)
	go func() {
		for i := 0; true; i++ {
			select {
			case msg := <-donecheck:
				if msg {
					rendertime <- i
					return
				}
			default:
				time.Sleep(1000000)
			}
		}
	}()
	
	for i := float64(0); i < 1000; i++ {
		img := CreateViewport(500, 500, 500)
		DrawMesh(img, cube)
		DrawWireframe(img, cube, color.White)
		var name string
		name = fmt.Sprintf("output/%v.png", i)
		SaveImage(img, name)
		cube = cube.rotate(0.1, 0)
		cube = cube.rotate(0.3, 1)
	}
	donecheck <- true
	msg := <-rendertime
	fmt.Printf("Done. %vms\n", msg)
	close(donecheck)
	close(rendertime)
}
