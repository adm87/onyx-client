package geom

import "math"

func ZeroVec2() Vec2 {
	return Vec2{0, 0}
}

func OneVec2() Vec2 {
	return Vec2{1, 1}
}

type Vec2 struct {
	X, Y float64
}

func (v Vec2) Add(other Vec2) Vec2 {
	return Vec2{
		X: v.X + other.X,
		Y: v.Y + other.Y,
	}
}

func (v Vec2) Sub(other Vec2) Vec2 {
	return Vec2{
		X: v.X - other.X,
		Y: v.Y - other.Y,
	}
}

func (v Vec2) Mul(scalar float64) Vec2 {
	return Vec2{
		X: v.X * scalar,
		Y: v.Y * scalar,
	}
}

func (v Vec2) Div(scalar float64) Vec2 {
	return Vec2{
		X: v.X / scalar,
		Y: v.Y / scalar,
	}
}

func (v Vec2) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Vec2) Normalize() Vec2 {
	length := v.Length()
	if length == 0 {
		return Vec2{0, 0}
	}
	return v.Div(length)
}

func (v Vec2) Dot(other Vec2) float64 {
	return v.X*other.X + v.Y*other.Y
}

func (v Vec2) Cross(other Vec2) float64 {
	return v.X*other.Y - v.Y*other.X
}
