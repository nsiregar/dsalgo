package factory

type Base struct {
	Type   string
	Width  float64
	Height float64
}

type Shape interface {
	ShapeType() string
}

type Circle struct {
	Radius float64
}

func (c Circle) ShapeType() string {
	return "circle"
}

func newCircle(dia float64) Circle {
	return Circle{
		Radius: dia / 2,
	}
}

type DummyShape struct {
	Width  float64
	Height float64
}

func (ds DummyShape) ShapeType() string {
	return "dummyshape"
}

func newDummyShape(w, h float64) DummyShape {
	return DummyShape{
		Width:  w,
		Height: h,
	}
}

func Factory(b Base) Shape {
	switch b.Type {
	case "circle":
		return newCircle(b.Width)
	default:
		return newDummyShape(b.Width, b.Height)
	}
}
