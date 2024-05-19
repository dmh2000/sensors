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
	dy        float64
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
	y         float64
	x         float64
	epsilon   float64
	state     int
}

type wave interface {
	step() (t float64, y float64, w wave)
}

// =======================================
// SIN WAVE
// =======================================

func newSin(amplitude float64, frequency float64, dt float64) wave {
	var w sin = sin{
		amplitude: amplitude,
		frequency: frequency * twopi,
		t:         0.0 - dt, // offset for start
		dt:        dt,
		y:         0.0,
	}
	return w
}

func (w sin) step() (t float64, y float64, wx wave) {
	// advance for next step
	w.t += w.dt

	// avoid overflow and loss of precision
	w.y = w.amplitude * math.Sin(w.t*w.frequency)

	return w.t, w.y, w
}

// =======================================
// TRIANGLE WAVE
// =======================================
func newTriangle(amplitude float64, frequency float64, dt float64) wave {

	var w triangle = triangle{
		amplitude: amplitude,
		frequency: frequency * twopi,
		t:         0.0 - dt,
		dt:        dt,
		// dy = (aplitude * 4 legs) * frequency * dt
		dy:    amplitude * 4.0 * frequency * dt,
		y:     0.0,
		state: triangleInit,
	}

	return w
}

const (
	triangleInit = iota
	triangleUp
	triangleDown
)

func (w triangle) step() (t float64, y float64, wx wave) {
	w.t += w.dt
	switch w.state {
	case triangleInit:
		w.y = 0.0
		w.state = triangleUp
	case triangleUp:
		w.y += w.dy
		if w.y >= w.amplitude {
			w.y = w.amplitude
			w.state = triangleDown
		}
	case triangleDown:
		w.y -= w.dy
		if w.y <= -w.amplitude {
			w.y = -w.amplitude

			w.state = triangleUp
		}
	default:
		panic("invalid state")
	}
	return w.t, w.y, w
}

// =======================================
// SQUARE WAVE
// =======================================
func newSquare(amplitude float64, frequency float64, dt float64) wave {
	var w square = square{
		amplitude: amplitude,
		frequency: frequency * twopi,
		period:    1.0 / frequency,
		t:         0.0 - dt,
		dt:        dt,
		dx:        dt / frequency,
		x:         0.0,
		y:         0.0,
		epsilon:   dt / (frequency * 2),
		state:     squareInit,
	}

	return w
}

const (
	squareInit = iota
	squareUp
	squareDown
)

func (w square) step() (t float64, y float64, wx wave) {
	w.t += w.dt
	w.x += w.dt
	switch w.state {
	case squareInit:
		w.y = w.amplitude
		w.x = 0.0
		w.state = squareUp
	case squareUp:
		if w.x >= (w.period/2.0)-w.epsilon {
			w.x = 0
			w.y = -w.amplitude
			w.state = squareDown
		}
	case squareDown:
		if w.x >= (w.period/2.0)-w.epsilon {
			w.x = 0.0
			w.y = w.amplitude
			w.state = squareUp
		}
	default:
		panic("invalid state")
	}
	return w.t, w.y, w
}

// dispatch
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
