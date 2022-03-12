package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

func n(s float64) float64 {
	if math.IsNaN(s) {
		return 0
	}
	return s
}

func inp(out string) (a float64, err error) {
	var ai string

	fmt.Print(out)
	fmt.Scan(&ai)
	return strconv.ParseFloat(ai, 64)
}

func input() (a, b, c, AB, AC, BC float64, err error) {
	k := 0.
	c, AB, AC, BC = math.NaN(), math.NaN(), math.NaN(), math.NaN()

	fmt.Println("If the side or angle is unknown, skip by writing `NaN`")

	a, err = inp("side a = ")
	if !math.IsNaN(a) {
		k++
	}
	if err != nil {
		return
	}

	b, err = inp("side b = ")
	if !math.IsNaN(b) {
		k++
	}
	if err != nil {
		return
	}

	c, err = inp("side c = ")
	if !math.IsNaN(c) {
		k++
	}
	if err != nil || k >= 3 {
		return
	}

	AB, err = inp("angle AB = ")
	if !math.IsNaN(AB) {
		k++
	}
	if err != nil || k >= 3 {
		return
	}

	AC, err = inp("angle AC = ")
	if !math.IsNaN(AC) {
		k++
	}
	if err != nil || k >= 3 {
		return
	}

	BC, err = inp("angle BC = ")
	if !math.IsNaN(BC) {
		k++
	}
	if err != nil || k >= 3 {
		return
	}

	return
}

func check(a, b, c, AB, AC, BC float64) error {
	if a <= 0 || b <= 0 || c <= 0 || AB <= 0 || AC <= 0 || BC <= 0 {
		return errors.New("value entered incorrectly")

	}

	if a+b <= c || a+c <= b || c+b <= a {
		return errors.New("no such triangle exists")
	}

	if !math.IsNaN(AB) && !math.IsNaN(AC) && !math.IsNaN(BC) {
		if AB+AC+BC != 180 {
			return errors.New("no such triangle exists")
		}
	} else if n(AB)+n(AC)+n(BC) >= 180 {
		return errors.New("no such triangle exists")

	}

	if AC >= 90 && (b <= a || b <= c) {
		return errors.New("no such triangle exists")
	} else if AB >= 90 && (c <= a || c <= b) {
		return errors.New("no such triangle exists")
	} else if BC >= 90 && (a <= c || a <= b) {
		return errors.New("no such triangle exists")
	}

	return nil
}

func areaA(AC, b, c, S float64) float64 {
	var h, x float64

	if !math.IsNaN(AC) && !math.IsNaN(b) && !math.IsNaN(c) {
		h = c * math.Sin((math.Pi*AC)/180)
		if AC <= 90 {
			x = math.Sqrt(c*c-h*h) + math.Sqrt(b*b-h*h)
			return (0.5 * x * h)
		}
		if AC > 90 {
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

func area(a, b, c, AB, AC, BC float64) (S float64, err error) {
	S = math.NaN()

	if !math.IsNaN(AB) && !math.IsNaN(AC) && math.IsNaN(BC) {
		BC = 180 - AB - AC
	} else if !math.IsNaN(AC) && !math.IsNaN(BC) && math.IsNaN(AB) {
		AB = 180 - AC - BC
	} else if !math.IsNaN(AB) && !math.IsNaN(BC) && math.IsNaN(AC) {
		AC = 180 - AB - BC
	}

	err = check(a, b, c, AB, AC, BC)
	if err != nil {
		return
	}

	if !math.IsNaN(a) && !math.IsNaN(b) && !math.IsNaN(c) {
		p := (a + b + c) / 2
		return math.Sqrt(p * (p - a) * (p - b) * (p - c)), nil
	}

	S = areaAngle(AB, a, b, S)
	S = areaAngle(AC, a, c, S)
	S = areaAngle(BC, b, c, S)

	S = areaTWOAngle(AB, AC, a, S)
	S = areaTWOAngle(AC, BC, c, S)
	S = areaTWOAngle(AB, BC, b, S)

	S = areaA(AC, b, c, S)
	S = areaA(AB, c, b, S)
	S = areaA(BC, a, c, S)
	S = areaA(AB, c, a, S)
	S = areaA(AC, b, a, S)
	S = areaA(BC, a, b, S)

	return
}

func main() {
	a, b, c, AB, AC, BC, err := input()

	exit := false
	for !exit {
		exit = true
		if err != nil {
			fmt.Println("ERROR:", err)
			fmt.Printf("\nTry it again\n\n")
			a, b, c, AB, AC, BC, err = input()
			exit = false
		}
	}

	S, err := area(a, b, c, AB, AC, BC)
	if err != nil {
		fmt.Println("ERROR:", err)
	} else if !math.IsNaN(S) {
		fmt.Printf("Area ≈ %.18f", S)
	} else {
		fmt.Print("Unable to calculate area insufficient data")
	}
}
