package vmath

type Matrix4 struct {
	M [16]float64
}

func Matrix4FromSqt(scale Vec3, rot Quaternion, trans Vec3) *Matrix4 {
	x := rot.X
	y := rot.Y
	z := rot.Z
	w := rot.W
	dx := x + x
	dy := y + y
	dz := z + z
	xy := x * dy
	xz := x * dz
	xw := dx * w
	yz := y * dz
	yw := dy * w
	zw := dz * w
	x2 := dx * x
	y2 := dy * y
	z2 := dz * z

	return &Matrix4{[16]float64{(1.0 - (y2 + z2))*scale.X, (xy + zw)*scale.X,         (xz - yw)*scale.X,         0.0,
								(xy - zw)*scale.Y,         (1.0 - (x2 + z2))*scale.Y, (yz + xw)*scale.Y,         0.0,
								(xz + yw)*scale.Z,         (yz - xw)*scale.Z,         (1.0 - (x2 + y2))*scale.Z, 0.0,
								trans.X,                   trans.Y,                   trans.Z,                   1.0}}
}

func (lhs *Matrix4) Transform(rhs Vec3) Vec3 {
	nx := (lhs.M[0] * rhs.X) + (lhs.M[4] * rhs.Y) + (lhs.M[8] * rhs.Z) + lhs.M[12]
	ny := (lhs.M[1] * rhs.X) + (lhs.M[5] * rhs.Y) + (lhs.M[9] * rhs.Z) + lhs.M[13]
	nz := (lhs.M[2] * rhs.X) + (lhs.M[6] * rhs.Y) + (lhs.M[10] * rhs.Z) + lhs.M[14]
	return Vec3{nx, ny, nz}
}

func (lhs *Matrix4) MultMatrix(rhs *Matrix4) *Matrix4 {
	ret := [16]float64{}

	for i := 0; i < 15; i += 4 {
		for j := 0; j < 4; j += 1 {
			ret[j+i] = lhs.M[i]*rhs.M[j] + lhs.M[i+1]*rhs.M[j+4] + lhs.M[i+2]*rhs.M[j+8] + lhs.M[i+3]*rhs.M[j+12]
		}
	}

	return &Matrix4{ret}
}
