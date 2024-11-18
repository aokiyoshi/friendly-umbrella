package render

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"os"

	"../camera"
	"../color"
	"../scene"
	"../shader"
)

func Render(camera camera.Camera, scene scene.Scene) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, camera.Width, camera.Height))
	fmt.Println("rendering...")
	for y := 0; y < camera.Height; y++ {
		for x := 0; x < camera.Width; x++ {
			ray := camera.GetRay(x, y)
			img.Set(
				x,
				y,
				color.Vec3ToColor(
					shader.ComputeColor(ray, scene),
				),
			)
		}
	}
	fmt.Println("done")
	return img
}

func Save(img image.Image, filename string) {
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	png.Encode(f, img)
}
