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

type wave interface {
	step() (float64, float64)
}

func newSin(amplitude float64, frequency float64, dt float64) wave {
	var s sin
	s.f = frequency
	s.a = amplitude
	s.t = 0.0
	s.dt = dt
	return &s
}

func newWave(wavetype string, amplitude float64, frequency float64, dt float64) wave {
	switch wavetype {
	case "sin":
		var s sin
		s.f = frequency
		s.a = amplitude
		s.t = 0.0
		s.dt = dt
		return &s
	default:
		panic("unknown wave type")
	}
}

func (s *sin) step() (float64, float64) {
	angle := math.Mod(2*math.Pi*s.f*s.t, 2*math.Pi)
	v := s.a * math.Sin(angle)
	s.t += s.dt
	return v, s.t
}
