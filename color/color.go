package color

import (
	"../vec3"

	imgcolor "image/color"
)

func SaturateColor(color vec3.Vec3) vec3.Vec3 {
	if color.X > 255 {
		color.X = 255
	}
	if color.Y > 255 {
		color.Y = 255
	}
	if color.Z > 255 {
		color.Z = 255
	}
	return color
}

func MultiplyColor(c vec3.Vec3, k float64) vec3.Vec3 {
	if k < 0 || k == 0 {
		return vec3.Vec3{}
	}
	return SaturateColor(
		vec3.Vec3{
			X: c.X * k,
			Y: c.Y * k,
			Z: c.Z * k,
		},
	)
}

func BlendColor(c1, c2 vec3.Vec3, k float64) vec3.Vec3 {
	if k == 0 {
		return c1
	}
	if k == 1 {
		return c2
	}
	return SaturateColor(
		vec3.Vec3{
			X: c1.X*(1-k) + c2.X*k,
			Y: c1.Y*(1-k) + c2.Y*k,
			Z: c1.Z*(1-k) + c2.Z*k,
		},
	)
}

func Vec3ToColor(c vec3.Vec3) imgcolor.Color {
	return imgcolor.RGBA{
		R: uint8(c.X),
		G: uint8(c.Y),
		B: uint8(c.Z),
		A: 255,
	}
}

func ColorToVec3(c imgcolor.Color) vec3.Vec3 {
	r, g, b, _ := c.RGBA()
	magic := 0.003891
	return vec3.Vec3{
		X: float64(r) * magic,
		Y: float64(g) * magic,
		Z: float64(b) * magic,
	}
}
