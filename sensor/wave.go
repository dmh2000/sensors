package main

import (
	"math"
)

const twopi = 2.0 * math.Pi

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
	step() (float64, float64, float64)
}

func newSin(amplitude float64, frequency float64, dt float64) wave {
	var w sin

	w.amplitude = amplitude
	w.frequency = frequency
	w.t = 0.0
	w.dt = dt
	w.y = 0.0
	return &w
}

func newTriangle(amplitude float64, frequency float64, dt float64) wave {
	var w triangle
	w.amplitude = amplitude
	w.frequency = frequency * (math.Pi * 2.0)
	w.t = 0.0
	w.dt = dt
	w.dx = dt * (1.0 / frequency)
	w.dy = amplitude * (dt * 4.0)
	w.x = 0.0
	w.y = 0.0
	w.state = 0
	return &w
}

func newSquare(amplitude float64, frequency float64, dt float64) wave {
	var w square
	w.amplitude = amplitude
	w.frequency = frequency * (math.Pi * 2.0)
	w.period = 1.0 / frequency
	w.t = 0.0
	w.dt = dt
	w.dx = dt * (1.0 / frequency)
	w.x = 0.0
	w.y = 0.0
	w.xm = 0.0
	w.epsilon = w.dt / 2.0
	w.state = 0
	return &w
}

func newWave(shape string, amplitude float64, frequency float64, dt float64) wave {
	switch shape {
	case "sin":
		return newSin(amplitude, frequency, dt)
	case "triangle":
		return newTriangle(amplitude, frequency, dt)
	case "square":
		return newSquare(amplitude, frequency, dt)
	default:
		panic("unknown wave type")
	}
}

func (w *sin) step() (dt float64, y float64, t float64) {
	tx := w.t
	// avoid overflow and loss of precision
	if tx >= twopi {
		tx -= twopi
	}

	w.y = w.amplitude * math.Sin(tx*w.frequency*twopi)
	w.t += w.dt
	return w.dt, w.y, w.t
}

func (w *triangle) step() (float64, float64, float64) {
	x := w.x
	t := w.t
	switch w.state {
	case 0:
		w.y = 0.0
		w.state = 1
	case 1:
		w.y += w.dy
		if w.y >= w.amplitude {
			w.y = w.amplitude
			w.state = 2
		}
	case 2:
		w.y -= w.dy
		if w.y <= -w.amplitude {
			w.y = -w.amplitude
			w.state = 3
		}
	case 3:
		w.y += w.dy
		if w.y >= 0.0 {
			w.state = 1
		}
	}
	w.t += w.dt
	w.x += w.dx
	return x, w.y, t
}

func (w *square) step() (float64, float64, float64) {
	x := w.x
	t := w.t
	switch w.state {
	case 0:
		w.y = w.amplitude
		w.xm = 0.0
		w.state = 1
	case 1:
		if w.xm >= (w.period/2.0)-w.epsilon {
			w.y = -w.amplitude
			w.state = 2
		}
	case 2:
		w.y = -w.amplitude
		if w.xm >= w.period-w.epsilon {
			w.xm = 0.0
			w.y = w.amplitude
			w.state = 1
		}
	}
	w.t += w.dt
	w.x += w.dx
	w.xm += w.dt
	return x, w.y, t
}
