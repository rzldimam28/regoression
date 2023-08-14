package main

import (
	"fmt"

	"github.com/rzldimam28/regoression/model"
)

func main() {
	fmt.Println("Welcome to ReGOression")

	x := []float64{530.0, 300.0, 358.0, 510.0, 302.0, 300.0, 387.0, 527.0, 415.0, 512.0}
	y := []float64{89.0, 48.0, 56.0, 72.0, 54.0, 42.0, 60.0, 85.0, 63.0, 74.0}

	regr := model.NewSimpleRegression(x, y)

	regr.Fit()
}