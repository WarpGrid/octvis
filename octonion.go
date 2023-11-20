package main

import "math"

type Octonion [8]float32;

func (o Octonion) Scale(r float32) Octonion {
	var p Octonion
	for i := 0; i < 8; i++ { p[i] = o[i] * r }
	return p
}

func (l Octonion) Mul(r Octonion) Octonion {
	var o Octonion

	o[0] = l[0]*r[0] -l[1]*r[1] -l[2]*r[2] -l[3]*r[3] -l[4]*r[4] -l[5]*r[5] -l[6]*r[6] -l[7]*r[7]
	o[1] = l[0]*r[1] + l[1]*r[0] +l[2]*r[4] +l[3]*r[7] -l[4]*r[2] +l[5]*r[6] -l[6]*r[5] -l[7]*r[3]
	o[2] = l[0]*r[2] + l[2]*r[0] -l[1]*r[4] +l[3]*r[5] +l[4]*r[1] -l[5]*r[3] +l[6]*r[7] -l[7]*r[6]
	o[3] = l[0]*r[3] + l[3]*r[0] -l[1]*r[7] -l[2]*r[5] +l[4]*r[6] +l[5]*r[2] -l[6]*r[4] +l[7]*r[1]
	o[4] = l[0]*r[4] + l[4]*r[0] +l[1]*r[2] -l[2]*r[1] -l[3]*r[6] +l[5]*r[7] +l[6]*r[3] -l[7]*r[5]
	o[5] = l[0]*r[5] + l[5]*r[0] -l[1]*r[6] +l[2]*r[3] -l[3]*r[2] -l[4]*r[7] +l[6]*r[1] +l[7]*r[4]
	o[6] = l[0]*r[6] + l[6]*r[0] +l[1]*r[5] -l[2]*r[7] +l[3]*r[4] -l[4]*r[3] -l[5]*r[1] +l[7]*r[2]
	o[7] = l[0]*r[7] + l[7]*r[0] +l[1]*r[3] +l[2]*r[6] -l[3]*r[1] +l[4]*r[5] -l[5]*r[4] -l[6]*r[2]

	return o
}

func (o Octonion) NormSq() float32 {
	return o[0]*o[0] +
		o[1]*o[1] +
		o[2]*o[2] +
		o[3]*o[3] +
		o[4]*o[4] +
		o[5]*o[5] +
		o[6]*o[6] +
		o[7]*o[7]
}

func (o Octonion) Norm() float32 {
	return float32(math.Sqrt(float64(o.NormSq())))
}

func (o Octonion) Normalized() Octonion {
	return o.Scale(1.0/o.Norm())
}
