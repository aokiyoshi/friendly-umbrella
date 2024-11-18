package shader

import (
	"math"

	"../material"

	"../vec3"

	"../color"
	"../intersection"
	"../ray"
	"../scene"
	"../trace_ray"

	"../world"
)

func ComputeColor(ray ray.Ray, scene scene.Scene) vec3.Vec3 {
	// check if ray hits something
	// if not, return black

	// check intersection
	res := trace_ray.TraceRay(ray.Origin, ray.Dir, scene.Spheres)

	// compute color if there is an intersection
	if res != nil {
		illuminate := ComputeLight(scene, *res)
		return color.MultiplyColor(
			color.BlendColor(
				res.M.Color, ComputeReflection(scene, *res, 8), 1.0-res.M.Roughness,
			),
			illuminate,
		)
	}

	// otherwise return black
	return world.GetSkyColor(ray.Dir)
}

func saturate(v float64) float64 {
	if v > 1.0 {
		return 1.0
	}
	if v < 0.0 {
		return 0.0
	}
	return v
}

func ComputeLight(scene scene.Scene, intersection intersection.Intersection) float64 {

	i := scene.Ambient

	for _, light := range scene.Lights {

		// generate shadow ray
		shadowRay := ray.Ray{
			Origin: intersection.P.Add(intersection.N.Kmul(0.001)),
			Dir:    light.Pos.Sub(intersection.P).Normalize(),
		}

		// check shadow
		shadow_trace := trace_ray.TraceRay(shadowRay.Origin, shadowRay.Dir, scene.Spheres)
		if shadow_trace != nil {
			continue
		}

		// diffuse
		roughness := intersection.M.Roughness
		light_dist := light.Pos.Sub(intersection.P).Length()
		l := light.Pos.Sub(intersection.P).Normalize()
		k_d := saturate(l.Dot(intersection.N))
		i += roughness * light.Intensity * k_d / light_dist

		// specular (kinda works?)
		v := intersection.V
		h := l.Add(v).Normalize()
		k_s := saturate(h.Dot(intersection.N))
		i += (1.0 - roughness) * light.Intensity * math.Pow(k_s, 500.0) / light_dist

		// metallic (TODO)
		// look cook-torrance
	}
	return saturate(i)
}

func ComputeReflection(scene scene.Scene, intersection intersection.Intersection, depth int) vec3.Vec3 {

	if depth == 0 {
		return vec3.Vec3{}
	}

	// prepare some variables
	v := intersection.V
	v_dot_n := v.Dot(intersection.N)

	// generate reflection ray
	ray := ray.Ray{
		Origin: intersection.P.Add(intersection.N.Kmul(0.001)),
		Dir:    intersection.N.Kmul(2.0 * v_dot_n).Sub(v).Normalize(),
	}

	// fresnel
	fresnel := material.GetFresnelTermFromIOR(intersection.M.IOR, v_dot_n)

	// check reflection
	res := trace_ray.TraceRay(ray.Origin, ray.Dir, scene.Spheres)

	// return black if no reflection
	if res == nil {
		return world.GetSkyColor(ray.Dir).Kmul(1.0 - fresnel)
	}

	return color.BlendColor(
		res.M.Color.Kmul(1.0-fresnel),
		ComputeReflection(scene, *res, depth-1),
		(1.0-res.M.Roughness)*(1.0-fresnel),
	)
}
