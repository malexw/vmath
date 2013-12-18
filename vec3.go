package vmath

import (
	"fmt"
	"math"
)

type Vec3 struct {
	X, Y, Z float64
}

func (lhs Vec3) Add(rhs Vec3) Vec3 {
	return Vec3{lhs.X+rhs.X, lhs.Y+rhs.Y, lhs.Z+rhs.Z}
}

func (lhs Vec3) Sub(rhs Vec3) Vec3 {
	return Vec3{lhs.X-rhs.X, lhs.Y-rhs.Y, lhs.Z-rhs.Z}
}

func (lhs Vec3) Scale(rhs float64) Vec3 {
	return Vec3{lhs.X*rhs, lhs.Y*rhs, lhs.Z*rhs}
}

func (lhs Vec3) Dot(rhs Vec3) float64 {
	return lhs.X*rhs.X + lhs.Y*rhs.Y + lhs.Z*rhs.Z
}

func (lhs Vec3) Cross(rhs Vec3) Vec3 {
	return Vec3{lhs.Y*rhs.Z - lhs.Z*rhs.Y, lhs.Z*rhs.X - lhs.X*rhs.Z, lhs.X*rhs.Y - lhs.Y*rhs.X}
}

func (v Vec3) LengthSquare() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

func (v Vec3) Length() float64 {
	return math.Sqrt(v.LengthSquare())
}

func (v Vec3) Normalize() Vec3 {
	len := 1/v.Length()
	return Vec3{v.X*len, v.Y*len, v.Z*len}
}

func (v Vec3) String() string {
	return fmt.Sprintf("v<%f, %f, %f>", v.X, v.Y, v.Z)
}