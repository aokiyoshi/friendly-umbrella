package main

import (
	"math"

	"./sphere"
	"./vec3"

	"./camera"

	"./light"
	"./material"
	"./render"
	"./scene"
)

func main() {

	// Spheres
	spheres := []sphere.Sphere{
		// sphere 1
		{
			Center: vec3.Vec3{
				X: 0,
				Y: 0,
				Z: 0,
			},
			Radius: 2.0,
			Material: material.Material{
				// 168, 50, 66
				Color:     vec3.Vec3{X: 168, Y: 50, Z: 66},
				Roughness: 1.0,
				Metallic:  0.99,
				IOR:       1.33,
			},
		},

		// sphere 2
		{
			Center: vec3.Vec3{
				X: -10,
				Y: 0,
				Z: 0,
			},
			Radius: 3.0,
			Material: material.Material{
				// 98, 196, 173
				Color:     vec3.Vec3{X: 98, Y: 196, Z: 173},
				Roughness: 0.0,
				Metallic:  0.0,
				IOR:       1.33,
			},
		},

		// sphere 3
		{
			Center: vec3.Vec3{
				X: 0,
				Y: -10,
				Z: 0,
			},
			Radius: 2.0,
			Material: material.Material{
				// 168, 50, 66
				Color:     vec3.Vec3{X: 168, Y: 50, Z: 66},
				Roughness: 0.9,
				Metallic:  0.0,
				IOR:       1.5,
			},
		},
	}

	// Lights
	lights := []light.Light{
		{
			Pos: vec3.Vec3{
				X: 0,
				Y: 0,
				Z: 10,
			},
			Intensity: 3.0,
		},
		{
			Pos: vec3.Vec3{
				X: 0,
				Y: -10,
				Z: -10,
			},
			Intensity: 10.0,
		},
	}

	// Scene
	scene := scene.Scene{
		Spheres: spheres,
		Lights:  lights,
		Ambient: 0.7,
	}

	// Camera
	camera := camera.Camera{
		Pos:    vec3.Vec3{X: 5, Y: 5, Z: 0},
		Dir:    vec3.Vec3{X: -1, Y: -1, Z: 0}.Normalize(),
		FOV:    math.Pi / 1.8,
		Focal:  1.25,
		Height: 1024,
		Width:  1024,
	}

	// Render
	img := render.Render(camera, scene)

	// Save image
	render.Save(img, "out.png")

}
