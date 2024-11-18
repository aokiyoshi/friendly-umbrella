package intersection

import (
	"../material"
	"../vec3"
)

type Intersection struct {
	T float64
	P vec3.Vec3
	N vec3.Vec3
	V vec3.Vec3
	M material.Material
}
