package simulator

import (
	"time"

	"github.com/nsf/termbox-go"
)

const renderingInterval = 100 * time.Millisecond

// Start listens to key events and renders the simulator.
func Start() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()

	eventQueue := make(chan termbox.Event)
	go func() {
		for {
			eventQueue <- termbox.PollEvent()
		}
	}()

	c := NewController()
	r := NewRenderer(c)
	r.render()

	for {
		select {
		case ev := <-eventQueue:
			if ev.Type == termbox.EventKey {
				switch {
				case ev.Key == termbox.KeyArrowUp:
					c.incrementNodes()

				case ev.Key == termbox.KeyArrowDown:
					c.decrementNodes()

				// Quit the program
				case ev.Ch == 'q' || ev.Key == termbox.KeyEsc || ev.Key == termbox.KeyCtrlC:
					return
				}
			}

		default:
			r.render()
			time.Sleep(renderingInterval)
		}
	}
}
