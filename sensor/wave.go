package main

import (
	"fmt"
	"math"
)

const twopi = 2.0 * math.Pi

const sinWave = "sin"
const triangleWave = "triangle"
const squareWave = "square"

type sin struct {
	amplitude float64
	frequency float64
	t         float64
	dt        float64
	y         float64
}

type triangle struct {
	amplitude float64
	frequency float64
	t         float64
	dt        float64
	dx        float64
	dy        float64
	x         float64
	y         float64
	state     int
}

type square struct {
	amplitude float64
	frequency float64
	period    float64
	t         float64
	dt        float64
	dx        float64
	x         float64
	y         float64
	xm        float64
	epsilon   float64
	state     int
}

type wave interface {
	step() (x float64, y float64)
}

func newSin(amplitude float64, frequency float64, dt float64) wave {
	var w sin = sin{
		amplitude: amplitude,
		frequency: frequency * twopi,
		t:         0.0,
		dt:        dt,
		y:         0.0,
	}
	return &w
}

func newTriangle(amplitude float64, frequency float64, dt float64) wave {
	var w triangle = triangle{
		amplitude: amplitude,
		frequency: frequency * twopi,
		t:         0.0,
		dt:        dt,
		dx:        dt * (1.0 / frequency),
		dy:        amplitude * (dt * 4.0),
		x:         0.0,
		y:         0.0,
		state:     triangleInit,
	}

	return &w
}

func newSquare(amplitude float64, frequency float64, dt float64) wave {
	var w square = square{
		amplitude: amplitude,
		frequency: frequency * twopi,
		period:    1.0 / frequency,
		t:         0.0,
		dt:        dt,
		dx:        dt * (1.0 / frequency),
		x:         0.0,
		y:         0.0,
		xm:        0.0,
		epsilon:   dt / 2.0,
		state:     squareInit,
	}

	return &w
}

var shapeFuncs = map[string]func(float64, float64, float64) wave{
	sinWave:      newSin,
	triangleWave: newTriangle,
	squareWave:   newSquare,
}

func newWave(shape string, amplitude float64, frequency float64, dt float64) (wave, error) {

	if f, ok := shapeFuncs[shape]; ok {
		return f(amplitude, frequency, dt), nil
	}
	return nil, fmt.Errorf("unknown shape: %s", shape)
}

func (w *sin) step() (x float64, y float64) {
	tx := w.t
	// avoid overflow and loss of precision
	if tx >= twopi {
		tx -= twopi
	}

	// advance for next step
	w.y = w.amplitude * math.Sin(tx*w.frequency)
	w.t += w.dt
	return tx, w.y
}

const (
	triangleInit = 0
	triangleUp   = 1
	triangleDown = 2
	triangleEnd  = 3
)

func (w *triangle) step() (x float64, y float64) {
	t := w.t
	switch w.state {
	case triangleInit:
		w.y = 0.0
		w.state = 1
	case triangleUp:
		w.y += w.dy
		if w.y >= w.amplitude {
			w.y = w.amplitude
			w.state = 2
		}
	case triangleDown:
		w.y -= w.dy
		if w.y <= -w.amplitude {
			w.y = -w.amplitude
			w.state = 3
		}
	case triangleEnd:
		w.y += w.dy
		if w.y >= 0.0 {
			w.state = 1
		}
	default:
		panic("invalid state")
	}
	// advance for next step
	w.t += w.dt
	w.x += w.dx
	return t, w.y
}

const (
	squareInit = 0
	squareUp   = 1
	squareDown = 2
)

func (w *square) step() (x float64, y float64) {
	t := w.t
	switch w.state {
	case squareInit:
		w.y = w.amplitude
		w.xm = 0.0
		w.state = 1
	case squareDown:
		if w.xm >= (w.period/2.0)-w.epsilon {
			w.y = -w.amplitude
			w.state = 2
		}
	case squareUp:
		w.y = -w.amplitude
		if w.xm >= w.period-w.epsilon {
			w.xm = 0.0
			w.y = w.amplitude
			w.state = 1
		}
	default:
		panic("invalid state")
	}
	// advance for next step
	w.t += w.dt
	w.x += w.dx
	w.xm += w.dt
	return t, w.y
}
