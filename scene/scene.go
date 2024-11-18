package scene

import (
	"../sphere"

	"../light"
)

type Scene struct {
	Spheres []sphere.Sphere
	Lights  []light.Light
	Ambient float64
}
