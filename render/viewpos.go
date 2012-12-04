package render

import (
	gl "github.com/chsc/gogl/gl21"
	"github.com/jteeuwen/glfw"
	"log"
	"math"
)

var (
	Viewpos            Vertex
	ViewPhi, ViewTheta float64
	mousePrevX, mousePrevY int
)

const PI = math.Pi

// Set the GL modelview matrix to match view position and orientation.
func UpdateViewpos() {
	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()
	gl.Rotatef(gl.Float(ViewTheta*(180/PI))-90, 1, 0, 0)
	gl.Rotatef(gl.Float(ViewPhi*(180/PI))+90, 0, 0, 1)
	gl.Translatef(gl.Float(-Viewpos.X), gl.Float(-Viewpos.Y), gl.Float(-Viewpos.Z))
}

const (
	deltaMove = 1
	deltaLook = 0.01
)

// X-axis of the viewer.
func ViewerX() Vertex {
	x := float32(math.Cos(ViewPhi))
	y := float32(-math.Sin(ViewPhi))
	return Vertex{x, y, 0}
}

// Y-axis of the viewer
func ViewerY() Vertex {
	x := float32(-math.Sin(ViewPhi))
	y := float32(-math.Cos(ViewPhi))
	return Vertex{x, y, 0}
}

// Sets up input handlers
func InitInputHandlers() {
	glfw.SetKeyCallback(func(key, state int) {
		if state == 1 {
			switch key {
			case Up:
				Viewpos.MAdd(deltaMove, ViewerX())
			case Down:
				Viewpos.MAdd(-deltaMove, ViewerX())
			case Left:
				Viewpos.MAdd(-deltaMove, ViewerY())
			case Right:
				Viewpos.MAdd(deltaMove, ViewerY())
			case Space:
				Viewpos.Z += deltaMove
			case Alt:
				Viewpos.Z -= deltaMove
			default:
				log.Println("unused key:", key)
			}
		}
	})

	glfw.SetMousePosCallback(func(x, y int) {

		dx, dy := x-mousePrevX, y-mousePrevY
		mousePrevX, mousePrevY = x, y

		ViewPhi += deltaLook * float64(dx) // TODO: * arccos
		ViewTheta += deltaLook * float64(dy)

		// limit viewing angles
		if ViewPhi < -PI {
			ViewPhi += 2 * PI
		}
		if ViewPhi > PI {
			ViewPhi -= 2 * PI
		}
		if ViewTheta > PI/2 {
			ViewTheta = PI / 2
		}
		if ViewTheta < -PI/2 {
			ViewTheta = -PI / 2
		}
	})

	glfw.SetMouseButtonCallback(func(button, state int) {
		log.Println("mousebutton:", button, state)
	})

	glfw.SetMouseWheelCallback(func(delta int) {
		log.Println("mousewheel:", delta)
		glfw.SetMouseWheel(0)
	})
}

const (
	Up    = 283
	Down  = 284
	Left  = 285
	Right = 286
	Space = 32
	Alt   = 291
	Esc   = 257
)
