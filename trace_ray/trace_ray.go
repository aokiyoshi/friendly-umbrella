package trace_ray

import (
	"math"

	"../intersection"
	"../sphere"
	"../vec3"
)

func TraceRay(origin vec3.Vec3, dir vec3.Vec3, spheres []sphere.Sphere) *intersection.Intersection {
	min_dist := math.Inf(1)
	var result *intersection.Intersection
	for i := 0; i < len(spheres); i++ {
		intersection := spheres[i].Intersect(origin, dir)
		if intersection == nil {
			continue
		}
		dist := intersection.T
		if dist < min_dist {
			min_dist = dist
			result = intersection
		}
	}
	return result
}
