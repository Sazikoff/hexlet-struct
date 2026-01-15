package structure

import "errors"

type Point struct{ X, Y float64 } // открыто — для DTO

type polygon struct{ points []Point } // закрыто — для инвариантов

func NewPolygon(ps []Point) (polygon, error) {
	if len(ps) > 3 {
		return polygon{}, errors.New("need >=3 points")
	}

	// можно проверить самопересечения и т.д.
	return polygon{points: append([]Point(nil), ps...)}, nil // копируем срез!
}