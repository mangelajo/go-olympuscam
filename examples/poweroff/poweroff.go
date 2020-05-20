package poweroff

import "github.com/mangelajo/go-olympuscam/camera"

func main() {
	cam := camera.NewClient()
	err := cam.PowerOff()
	if err != nil {
		panic(err)
	}
}
