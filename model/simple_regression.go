package model

import (
	"errors"
	"fmt"
)

type Coefficient struct {
	B0 float64
	B1 float64
}

type SimpleRegression struct {
	X []float64
	Y []float64
	Coefficient Coefficient
}

func NewSimpleRegression(x []float64, y []float64) SimpleRegression {
	return SimpleRegression{
		X: x,
		Y: y,
	}
}

func (ths *SimpleRegression) Fit() SimpleRegression {

	n := ths.length()
	xMultY, _ := ths.multiply()
	sumXMultY := ths.sumArr(xMultY)
	sumX := ths.sumArr(ths.X)
	sumY := ths.sumArr(ths.Y)
	xSquared := ths.squared(ths.X)
	sumXSquared := ths.sumArr(xSquared)
	squaredX := sumX * sumX

	// fmt.Println("n", n)
	// fmt.Println("sumXMultY", sumXMultY)
	// fmt.Println("sumX", sumX)
	// fmt.Println("sumY", sumY)
	// fmt.Println("sumXSquared", sumXSquared)
	// fmt.Println("squaredX", squaredX)

	b0 := ((sumY * sumXSquared) - (sumX * sumXMultY)) / ((float64(n) * sumXSquared) - squaredX)
	b1 := ((float64(n) * sumXMultY) - (sumX * sumY)) / ((float64(n) * sumXSquared) - squaredX)

	fmt.Println("------------ REGRESSION MODEL --------------")
	model := fmt.Sprintf("%f + %fX", b0, b1)
	fmt.Println(model)

	coef := Coefficient{
		B0: b0,
		B1: b1,
	}

	return SimpleRegression{
		X: ths.X,
		Y: ths.Y,
		Coefficient: coef,
	}
}

func (ths *SimpleRegression) length() int {
	return len(ths.X)
}

func (ths *SimpleRegression) sumArr(arr []float64) float64 {
	var sum float64
	for _, value := range arr {
		sum += value
	}
	return sum
}

func (ths *SimpleRegression) squared(arr []float64) []float64 {
	var squaredArr []float64
	for _, value := range arr {
		squaredArr = append(squaredArr, value*value)
	}
	return squaredArr
}

func (ths *SimpleRegression) multiply() ([]float64, error) {
	if len(ths.X) != len(ths.Y) {
		return []float64{}, errors.New("vector length are not the same")
	}

	result := make([]float64, len(ths.X))
	for i := 0; i < len(ths.X); i++ {
		result = append(result, ths.X[i] * ths.Y[i])
	}

	return result, nil
}
