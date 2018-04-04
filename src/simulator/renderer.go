package simulator

import (
	"strconv"

	"github.com/nsf/termbox-go"
)

const defaultBackgroundColor = termbox.ColorBlack
const defaultColor = termbox.ColorWhite

// Renderer struct contains all the controller state.
type Renderer struct {
	controller   *Controller
	screenHeigth int
	screenWidh   int
}

// NewRenderer returns a fully initialized renderer.
func NewRenderer(c *Controller) (r *Renderer) {
	r = new(Renderer)
	r.reset(c)
	return
}

func (r *Renderer) reset(c *Controller) {
	r.controller = c
	r.screenWidh, r.screenHeigth = termbox.Size()
}

func (r *Renderer) render() {
	termbox.Clear(defaultColor, defaultBackgroundColor)
	r.renderContainer()
	r.renderData()
	termbox.Flush()
}

func (r *Renderer) renderContainer() {
	termbox.SetCell(0, 0, '╔', defaultColor, defaultBackgroundColor)
	termbox.SetCell(r.screenWidh-1, 0, '╗', defaultColor, defaultBackgroundColor)
	termbox.SetCell(r.screenWidh-1, r.screenHeigth-1, '╝', defaultColor, defaultBackgroundColor)
	termbox.SetCell(0, r.screenHeigth-1, '╚', defaultColor, defaultBackgroundColor)
	for x := 1; x < r.screenWidh-1; x++ {
		termbox.SetCell(x, 0, '═', defaultColor, defaultBackgroundColor)
		if x != r.screenWidh-50 {
			termbox.SetCell(x, 2, '─', defaultColor, defaultBackgroundColor)
			termbox.SetCell(x, r.screenHeigth-1, '═', defaultColor, defaultBackgroundColor)
		} else {
			termbox.SetCell(x, 2, '┬', defaultColor, defaultBackgroundColor)
			termbox.SetCell(x, r.screenHeigth-1, '╧', defaultColor, defaultBackgroundColor)
		}
	}
	for y := 1; y < r.screenHeigth-1; y++ {
		termbox.SetCell(0, y, '║', defaultColor, defaultBackgroundColor)
		termbox.SetCell(r.screenWidh-1, y, '║', defaultColor, defaultBackgroundColor)

		if y >= 3 {
			termbox.SetCell(r.screenWidh-50, y, '│', defaultColor, defaultBackgroundColor)
		}
	}

	draw(2, 1, "ATHENA SIMULATOR", defaultColor, defaultBackgroundColor)

	draw(2, r.screenHeigth-3, "▲", termbox.ColorBlue, defaultBackgroundColor)
	draw(4, r.screenHeigth-3, "Increase Nodes", defaultColor, defaultBackgroundColor)
	draw(2, r.screenHeigth-2, "▼", termbox.ColorBlue, defaultBackgroundColor)
	draw(4, r.screenHeigth-2, "Decrease Nodes", defaultColor, defaultBackgroundColor)
}

func (r *Renderer) renderData() {
	draw(r.screenWidh-48, 3, "Nodes", defaultColor, defaultBackgroundColor)
	draw(r.screenWidh-38, 3, strconv.Itoa(r.controller.nodesLength()), termbox.ColorCyan, defaultBackgroundColor)
}

func draw(x, y int, text string, color, backgroundColor termbox.Attribute) {
	for _, char := range text {
		termbox.SetCell(x, y, char, color, backgroundColor)
		x++
	}
}
