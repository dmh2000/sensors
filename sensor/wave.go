package main

import (
	"math"
)

type sin struct {
	a  float64
	f  float64
	t  float64
	dt float64
}

type triangle struct {
	a  float64
	f  float64
	t  float64
	dt float64

	state int
	dx    float64
	dy    float64
	x     float64
	y     float64
	p1    float64
	p2    float64
	p3    float64
}

type wave interface {
	step() (float64, float64, float64)
}

func newSin(amplitude float64, frequency float64, dt float64) wave {
	var w sin
	w.f = frequency
	w.a = amplitude
	w.t = 0.0
	w.dt = dt
	return &w
}

func newTriangle(amplitude float64, frequency float64, dt float64) wave {
	var w triangle
	w.f = frequency
	w.a = amplitude
	w.t = 0.0
	w.dt = dt

	w.state = 0
	w.dx = w.dt * w.f
	w.dy = w.a / w.f
	w.x = 0.0
	w.y = 0.0
	w.p1 = w.f / 4.0
	w.p2 = w.f * 3.0 / 4.0
	w.p3 = w.f
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

func (s *sin) step() (float64, float64, float64) {
	t := math.Mod(s.t, s.f)
	x := t * s.f
	y := s.a * math.Sin(x)
	s.t += s.dt
	return x, y, s.t
}

func (s *triangle) step() (float64, float64, float64) {
	switch s.state {
	case 0:
		s.y = 0.0
		s.state = 1
	case 1:
		s.x += s.dx
		s.y += s.dy
		if s.y >= s.a {
			s.y = s.a
			s.state = 2
		}
	case 2:
		s.x += s.dx
		s.y -= s.dy
		if s.y <= -s.a {
			s.y = -s.a
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
