package sphere

import (
	"math"

	"../intersection"
	"../material"
	"../vec3"
)

type Sphere struct {
	Center   vec3.Vec3
	Radius   float64
	Material material.Material
}

func (v Sphere) Intersect(origin vec3.Vec3, dir vec3.Vec3) *intersection.Intersection {
	oc := origin.Sub(v.Center)
	a := dir.Dot(dir)
	b := 2 * oc.Dot(dir)
	c := oc.Dot(oc) - v.Radius*v.Radius
	discriminant := b*b - 4*a*c
	if discriminant < 0 {
		return nil
	}
	t := (-b - math.Sqrt(discriminant)) / (2 * a)
	if t < 0 {
		return nil
	}
	p := origin.Add(dir.Kmul(t))
	viewDir := origin.Sub(v.Center).Normalize()
	return &intersection.Intersection{
		T: t,
		P: p,
		N: v.Normal(p),
		V: viewDir,
		M: v.Material,
	}
}

func (v Sphere) Normal(point vec3.Vec3) vec3.Vec3 {
	return point.Sub(v.Center).Normalize()
}
