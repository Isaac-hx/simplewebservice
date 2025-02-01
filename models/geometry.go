package models

import (
	"math"
)

type Shape2D interface {
	Area() float64
	Perimeter() float64
}
type Shape3D interface {
	Shape2D
	Volume() float64
}

type Cube struct {
	Side float64 `json:"side"`
}

func (c Cube) Area() float64 {
	return 6 * math.Pow(c.Side, 2)
}

func (c Cube) Perimeter() float64 {
	return 12 * c.Side
}

func (c Cube) Volume() float64 {
	return math.Pow(c.Side, 3)
}

type Transformer interface {
	Rotate(degrees float64)
}

// implementasi interface Transformer pada struct Cube
func (c *Cube) Rotate(degrees float64) {
	//Penambahan efek dari gaya rotasi
	c.Side = c.Side * math.Cos(degrees)
}

func CalculateRotateShape(s Transformer, degrees float64) map[string]interface{} {
	s.Rotate(degrees)
	return map[string]interface{}{
		"side": s.(*Cube).Side,
	}

}
func CalculateShape(s Shape2D) map[string]interface{} {
	var area, perimeter, volume interface{}
	area = s.Area()
	perimeter = s.Perimeter()
	ok, err := s.(Shape3D)
	if !err {
		volume = "Tidak memiliki nilai volume"
		return map[string]interface{}{
			"area":      area,
			"perimeter": perimeter,
			"volume":    volume,
		}
	}
	volume = ok.Volume()
	return map[string]interface{}{
		"area":      area,
		"perimeter": perimeter,
		"volume":    volume,
	}
}
