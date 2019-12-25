package util

import (
	"github.com/shopspring/decimal"
)

// a + b
func Add(a float64, b ...float64) float64 {
	m := decimal.NewFromFloat(a)
	for _, v := range b {
		m = m.Add(decimal.NewFromFloat(v))
	}
	f, _ := m.Round(2).Float64()
	//f, _ := decimal.NewFromFloat(a).Add(decimal.NewFromFloat(b)).Round(2).Float64()
	return f
}

// a - b
func Sub(a float64, b ...float64) float64 {
	m := decimal.NewFromFloat(a)
	for _, v := range b {
		m = m.Sub(decimal.NewFromFloat(v))
	}
	f, _ := m.Round(2).Float64()
	//f, _ := decimal.NewFromFloat(a).Sub(decimal.NewFromFloat(b)).Round(2).Float64()
	return f
}

// a * b
func Mul(a float64, b ...float64) float64 {
	m := decimal.NewFromFloat(a)
	for _, v := range b {
		m = m.Mul(decimal.NewFromFloat(v))
	}
	f, _ := m.Round(2).Float64()
	//f, _ := decimal.NewFromFloat(a).Mul(decimal.NewFromFloat(b)).Round(2).Float64()
	return f
}

// a / b
func Div(a float64, b float64) float64 {
	f, _ := decimal.NewFromFloat(a).Div(decimal.NewFromFloat(b)).Round(2).Float64()
	return f
}
