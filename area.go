package main

import (
	"fmt"
	"math"
	"strconv"
)

func inp(out string) (a float64, err error) {
	var ai string

	fmt.Print(out)
	fmt.Scan(&ai)
	return strconv.ParseFloat(ai, 64)
}

func input() (a, b, c, AB, AC, BC float64, err error) {

	fmt.Println("If the side or angle is unknown, skip by writing `NaN`")

	a, err = inp("a = ")
	if err != nil {
		return
	}

	b, err = inp("b = ")
	if err != nil {
		return
	}

	c, err = inp("c = ")
	if err != nil {
		return
	}

	AB, err = inp("AB = ")
	if err != nil {
		return
	}

	AC, err = inp("AC = ")
	if err != nil {
		return
	}

	BC, err = inp("BC = ")
	if err != nil {
		return
	}

	if a <= 0 || b <= 0 || c <= 0 || AB <= 0 || AC <= 0 || BC <= 0 {
		err = fmt.Errorf("Value entered incorrectly")
	}

	if a+b <= c || a+c <= b || c+b <= a {
		err = fmt.Errorf("No such triangle exists")
	}

	if !math.IsNaN(AB) && !math.IsNaN(AC) && !math.IsNaN(BC) {
		if AB+AC+BC != 180 {
			err = fmt.Errorf("No such triangle exists")
		}
	} else if AB+AC+BC > 180 {
		err = fmt.Errorf("No such triangle exists")
	}

	if !math.IsNaN(AB) && !math.IsNaN(AC) {
		BC = 180 - AB - AC
	} else if !math.IsNaN(AC) && !math.IsNaN(BC) {
		AB = 180 - AC - BC
	} else if !math.IsNaN(AB) && !math.IsNaN(BC) {
		AC = 180 - AB - BC
	}

	if AC >= 90 && (b <= a || b <= c) {
		err = fmt.Errorf("No such triangle exists")
	} else if AB >= 90 && (c <= a || c <= b) {
		err = fmt.Errorf("No such triangle exists")
	} else if BC >= 90 && (a <= c || a <= b) {
		err = fmt.Errorf("No such triangle exists")
	}

	return
}

// ПО ПИФАГОРУ
//
func areaA(AC, b, c, S float64) float64 {
	var h, x float64

	if !math.IsNaN(AC) && !math.IsNaN(b) && !math.IsNaN(c) {
		h = c * math.Sin((math.Pi*AC)/180)
		if AC < 90 {
			x = math.Sqrt(c*c-h*h) + math.Sqrt(b*b-h*h)
			return (0.5 * x * h)
		}
		if AC >= 90 {
			x = math.Sqrt(b*b-h*h) - math.Sqrt(c*c-h*h)
			return (0.5 * x * h)
		}
	}

	return S
}

func areaAngle(AB, a, b, S float64) float64 {
	if !math.IsNaN(AB) && !math.IsNaN(a) && !math.IsNaN(b) {
		return (0.5 * a * b * math.Sin((math.Pi*AB)/180))
	}

	return S
}

func areaTWOAngle(AB, AC, a, S float64) float64 {
	if !math.IsNaN(AB) && !math.IsNaN(AC) && !math.IsNaN(a) {
		return a * a * math.Sin((math.Pi*AB)/180) * math.Sin((math.Pi*AC)/180) * 0.5 / math.Sin((math.Pi*(180-(AB+AC)))/180)
	}
	return S
}

func area(a, b, c, AB, AC, BC float64) (S float64) {
	S = math.NaN()
	if !math.IsNaN(a) && !math.IsNaN(b) && !math.IsNaN(c) {
		p := (a + b + c) / 2
		S = math.Sqrt(p * (p - a) * (p - b) * (p - c))
	}

	S = areaAngle(AB, a, b, S)
	S = areaAngle(AB, a, c, S)
	S = areaAngle(AB, b, c, S)

	S = areaTWOAngle(AB, AC, a, S)
	S = areaTWOAngle(AC, BC, c, S)
	S = areaTWOAngle(AB, BC, b, S)

	S = areaA(AC, b, c, S)
	S = areaA(AB, c, b, S)
	S = areaA(BC, a, c, S)
	S = areaA(AB, c, a, S)
	S = areaA(AC, b, a, S)
	S = areaA(BC, a, b, S)

	return S
}

func main() {
	a, b, c, AB, AC, BC, err := input()
	exit := false

	for !exit {
		exit = true
		if err != nil {
			fmt.Println("ERROR: ", err)
			fmt.Printf("\nTry it again\n\n")
			a, b, c, AB, AC, BC, err = input()
			exit = false
		}
	}
	S := area(a, b, c, AB, AC, BC)
	if !math.IsNaN(S) {
		fmt.Printf("Area ≈ %.18f", S)
	} else {
		fmt.Print("Unable to calculate area insufficient data")
	}
}
