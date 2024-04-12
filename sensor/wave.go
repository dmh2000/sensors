package main

import (
	"math"
)

type sin struct {
	amplitude float64
	frequency float64
	t         float64
	dt        float64
	dx        float64
	x         float64
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
	t         float64
	dt        float64
	dx        float64
	dy        float64
	x         float64
	y         float64
	state     int
}

type wave interface {
	step() (float64, float64, float64)
}

func newSin(amplitude float64, frequency float64, dt float64) wave {
	var w sin

	w.amplitude = amplitude
	w.frequency = frequency * (math.Pi * 2.0) // convert to radians/sec
	w.t = 0.0
	w.dt = dt
	w.dx = dt * frequency
	w.x = 0.0
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
	w.frequency = frequency
	w.amplitude = amplitude
	w.t = 0.0
	w.dt = dt

	w.state = 0
	rate := 1.0 / dt
	w.dx = rate / frequency
	w.dy = w.amplitude * w.dx * 2.0 * math.Pi
	w.x = 0.0
	w.y = 0.0
	return &w
}

func newWave(wavetype string, amplitude float64, frequency float64, dt float64) wave {
	switch wavetype {
	case "sin":
		return newSin(amplitude, frequency, dt)
	case "triangle":
		return newTriangle(amplitude, frequency, dt)
	default:
		panic("unknown wave type")
	}
}

func (w *sin) step() (float64, float64, float64) {
	t := w.t
	x := w.x
	w.y = w.amplitude * math.Sin(w.t*w.frequency)
	w.t += w.dt
	w.x += w.dx
	return x, w.y, t
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

func (s *square) step() (float64, float64, float64) {
	switch s.state {
	case 0:
		s.y = s.amplitude
		s.state = 1
	case 1:
		s.x += s.dx
		s.y += s.dy
		if s.x >= s.amplitude {
			s.y = s.amplitude
			s.state = 2
		}
	case 2:
		s.x += s.dx
		s.y -= s.dy
		if s.y <= -s.amplitude {
			s.y = -s.amplitude
			s.state = 3
		}
	case 3:
		s.x += s.dx
		s.y += s.dy
		if s.y >= 0 {
			s.state = 1
		}
	}
	s.t += s.dt
	return s.x, s.y, s.t
}
