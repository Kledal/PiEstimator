package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	total := int64(0)
	circle := int64(0)

	r := float64(200)
	r2 := r * r

	var x float64
	var y float64

	var dist float64
	var newPi float64

	for {

		for i := 0; i < 10000000; i++ {
			x = randFloat(-r, r)
			y = randFloat(-r, r)

			dist = x*x + y*y
			if dist < r2 {
				circle += 1
			}

			total += 1

			newPi = float64(4 * (float64(circle) / float64(total)))
		}

		fmt.Println(fmt.Sprintf("4 * %v / %v = %v", circle, total, newPi))
	}
}

func randFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}
