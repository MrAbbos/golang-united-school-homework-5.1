package square

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (s Square) End() (p Point) {
	p.x = s.start.x + int(s.a)
	p.y = s.start.y + int(s.a)
	return
}

func (s Square) Area() uint {
	return uint(s.a*s.a)
}

func (s Square) Perimeter() uint {
	return uint((int(s.a))*4)
}
