package bridge

import "fmt"

type iDrawShape interface {
	drawShape(x [5]float32, y [5]float32)
}

type drawShape struct{}

func (d drawShape) drawShape(x [5]float32, y [5]float32) {
	fmt.Println("Drawing Shape")
}

type iContour interface {
	drawContour(x [5]float32, y [5]float32)
	resizeByFactor(factor int)
}

type drawContour struct {
	x      [5]float32
	y      [5]float32
	shape  drawShape
	factor int
}

func (contour drawContour) drawContour(x [5]float32, y [5]float32) {
	fmt.Println("Drawing Contour")
	contour.shape.drawShape(contour.x, contour.y)
}

func (contour drawContour) resizeByFactor(factor int) {
	contour.factor = factor
}

func bridgeDesign() {
	var x = [5]float32{1, 2, 3, 4, 5}
	var y = [5]float32{1, 2, 3, 4, 5}
	var contour iContour = drawContour{x, y, drawShape{}, 2}
	contour.drawContour(x, y)
	contour.resizeByFactor(2)
}
