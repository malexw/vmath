package vmath

import (
	"math"
)

type Quaternion struct {
	X, Y, Z, W float64
}

func QuaternionFromAxisAngle(axis Vec3, angle float64) Quaternion {
	angleTerm := angle * math.Pi / 360
	sinAngle := math.Sin(angleTerm)
	return Quaternion{axis.X * sinAngle, axis.Y * sinAngle, axis.Z * sinAngle, math.Cos(angleTerm)}
}

func (lhs Quaternion) Normalize() Quaternion {
	f := 1/math.Sqrt(lhs.X*lhs.X + lhs.Y*lhs.Y + lhs.Z*lhs.Z + lhs.W*lhs.W)
	return Quaternion{lhs.X*f, lhs.Y*f, lhs.Z*f, lhs.W*f}
}