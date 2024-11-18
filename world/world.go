package world

import (
	"../color"

	"../vec3"
)

func GetSkyColor(dir vec3.Vec3) vec3.Vec3 {

	blue := vec3.Vec3{X: 0, Y: 0, Z: 255.0}
	white := vec3.Vec3{X: 255.0, Y: 255.0, Z: 255.0}

	k := 0.0

	if dir.Z > 0 {
		k = dir.Z
	}

	return color.BlendColor(white, blue, k)
}
