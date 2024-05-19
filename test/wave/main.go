package main

import (
	"fmt"
)

func main() {
	// w, err := newWave("sin", 1.5, 0.1, 0.2)
	// w, err := newWave("triangle", 1.0, 1.0, 0.05)
	w, err := newWave("square", 2.0, 2.0, 0.05)
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Println(math.Mod(twopi-0.1, twopi),
	// 	math.Mod(twopi-0.01, twopi),
	// 	math.Mod(twopi-0.001, twopi),
	// 	math.Mod(twopi-0.0001, twopi),
	// 	math.Mod(twopi-0.00001, twopi),
	// 	math.Mod(twopi-0.000001, twopi))

	for i := 0; i < 30; i++ {
		var t float64
		var y float64
		t, y, w = w.step()
		fmt.Printf("%f,%f\n", t, y)
	}
}
