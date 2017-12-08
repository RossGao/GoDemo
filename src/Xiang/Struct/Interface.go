package main

type Shape interface {
	Area() float64
}

func TotalArea(shapes ...Shape) float64 {
	total := float64(0)
	for _, shape := range shapes {
		total += shape.Area()
	}

	return total
}
