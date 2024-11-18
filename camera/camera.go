package camera

import (
	"math"

	"../ray"
	"../vec3"
)

type Camera struct {
	Pos    vec3.Vec3
	Dir    vec3.Vec3
	FOV    float64
	Focal  float64
	Width  int
	Height int
}

func (c Camera) GetRay(u, v int) ray.Ray {

	// This camera has no roll rotation

	aspect := float64(c.Width) / float64(c.Height)

	halfWidth := c.Focal * math.Tan(c.FOV/2.0)
	halfHeight := halfWidth / aspect

	rightVec := c.Dir.Cross(vec3.Vec3{X: 0, Y: 0, Z: 1}).Kmul(-1).Normalize()

	upVec := rightVec.Cross(c.Dir).Normalize()

	x := halfWidth * float64(u-c.Width/2) / float64(c.Width)
	y := halfHeight * float64(v-c.Height/2) / float64(c.Height)

	screenPoint := c.Dir.Add(rightVec.Kmul(x)).Add(upVec.Kmul(y)).Normalize()
	return ray.Ray{Origin: c.Pos, Dir: screenPoint}
}
