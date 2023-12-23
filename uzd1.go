package main

import (
	"fmt"
	"math"
)

type Quadratic struct {
	a, b, c      float64
	Diskriminant float64
}

type QuadraticError struct {
	error        string
	Diskriminant float64
	A            float64
	B            float64
	C            float64
}

func NewQuadratic(a, b, c float64) Quadratic {
	return Quadratic{a: a, b: b, c: c}
}

func (q Quadratic) Roots() (arr []float64, e error) {
	arr = make([]float64, 0)
	Diskriminantus := math.Pow(q.b, 2) - (4 * q.a * q.c)
	err := &QuadraticError{Diskriminant: Diskriminantus, A: q.a, B: q.b, C: q.c}
	message := ""
	if err.InvalidDiskriminant() {
		message = "Invalid Diskriminantus"
	}
	if err.InvalidKoefA() {
		if message != "" {
			message += ", "
		}
		message = message + "Invalid A"
	}

	var x1, x2 float64
	if Diskriminantus == 0 {
		x1 = (-q.b + math.Sqrt(Diskriminantus)) / 2 * q.a
		arr = append(arr, x1)
	}
	if Diskriminantus > 0 {
		x1 = (-q.b + math.Sqrt(Diskriminantus)) / 2 * q.a
		x2 = (-q.b - math.Sqrt(Diskriminantus)) / 2 * q.a
		arr = append(arr, x1)
		arr = append(arr, x2)
	}

	if message != "" {
		err.error = message
		e = err
	} else {
		e = nil
	}

	return arr, e
}

func (q QuadraticError) Error() string {
	return fmt.Sprintln("Can not solve this stuff because of the ", q.Diskriminant, "and koefs ", q.A, q.B, q.C, q.error)
}

func (q QuadraticError) InvalidKoefA() bool {
	return q.A == 0
}

func (q QuadraticError) InvalidDiskriminant() bool {
	return q.Diskriminant < 0
}

func main() {
	q := NewQuadratic(1, 0, 4)
	root, err := q.Roots()
	if err != nil {
		fmt.Println(err)
		if e, ok := err.(*QuadraticError); ok {
			if e.InvalidDiskriminant() {
				fmt.Printf("error: diskriminant %0.2f is invalid\n", e.Diskriminant)
			}
			if e.InvalidKoefA() {
				fmt.Printf("error: koef A %0.2f is invalid\n", e.A)
			}
		} else {
			fmt.Println("Generating error")
		}
		return
	}
	fmt.Println("roots is", root)
}
