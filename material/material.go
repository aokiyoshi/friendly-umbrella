package material

import (
	"math"

	"../vec3"
)

type Material struct {
	Color     vec3.Vec3
	Roughness float64
	Metallic  float64
	IOR       float64
}

func GetFresnelTerm(n1, n2, cosTheta float64) float64 {
	r0 := math.Pow((n1-n2)/(n1+n2), 2)
	return r0 + (1-r0)*((1+cosTheta)*(1+cosTheta))*(math.Pow(1-cosTheta, 5))
}

func GetFresnelTermFromIOR(ior, cosTheta float64) float64 {
	return ior - (1-ior)*((1+cosTheta)*(1+cosTheta))*(math.Pow(1-cosTheta, 5))
}
