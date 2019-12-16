package utils

import (
	"github.com/shopspring/decimal"
)

// a + b
func Add(a, b float64) float64 {
	f, _ := decimal.NewFromFloat(a).Add(decimal.NewFromFloat(b)).Round(2).Float64()
	return f
}

// a - b
func Sub(a, b float64) float64 {
	f, _ := decimal.NewFromFloat(a).Sub(decimal.NewFromFloat(b)).Round(2).Float64()
	return f
}

// a * b
func Mul(a, b float64) float64 {
	f, _ := decimal.NewFromFloat(a).Mul(decimal.NewFromFloat(b)).Round(2).Float64()
	return f
}

// a / b
func Div(a, b float64) float64 {
	f, _ := decimal.NewFromFloat(a).Div(decimal.NewFromFloat(b)).Round(2).Float64()
	return f
}
