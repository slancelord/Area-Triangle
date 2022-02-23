package main

import (
	"fmt"
	"math"
	"strconv"
)

func input() (a, b, c, AB, AC, BC float64, err error) {
	var ai, bi, ci, ABi, ACi, BCi string

	fmt.Println("If the side or angle is unknown, skip by writing `NaN`")

	fmt.Print("a = ")
	fmt.Scan(&ai)
	a, err = strconv.ParseFloat(ai, 64)
	if err != nil {
		return
	}

	fmt.Print("b = ")
	fmt.Scan(&bi)
	b, err = strconv.ParseFloat(bi, 64)
	if err != nil {
		return
	}

	fmt.Print("c = ")
	fmt.Scan(&ci)
	c, err = strconv.ParseFloat(ci, 64)
	if err != nil {
		return
	}

	fmt.Print("AB = ")
	fmt.Scan(&ABi)
	AB, err = strconv.ParseFloat(ABi, 64)
	if err != nil {
		return
	}

	fmt.Print("AC = ")
	fmt.Scan(&ACi)
	AC, err = strconv.ParseFloat(ACi, 64)
	if err != nil {
		return
	}

	fmt.Print("BC = ")
	fmt.Scan(&BCi)
	BC, err = strconv.ParseFloat(BCi, 64)
	if err != nil {
		return
	}

	if a <= 0 || b <= 0 || c <= 0 || AB <= 0 || AC <= 0 || BC <= 0 || AB+AC+BC >= 360 {
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

	return

}

func areaAngle(AB, a, b float64) float64 {
	return (0.5 * a * b * math.Sin((math.Pi*AB)/180))
}

func areaTWOAngle(AB, AC, a float64) float64 {
	return a * a * math.Sin((math.Pi*AB)/180) * math.Sin((math.Pi*AC)/180) * 0.5 / math.Sin((math.Pi*(180-(AB+AC)))/180)
}

func area(a, b, c, AB, AC, BC float64) float64 {
	if !math.IsNaN(a) && !math.IsNaN(b) && !math.IsNaN(c) {
		p := (a + b + c) / 2
		return math.Sqrt(p * (p - a) * (p - b) * (p - c))
	}

	if !math.IsNaN(AB) && !math.IsNaN(a) && !math.IsNaN(b) {
		return areaAngle(AB, a, b)
	}
	if !math.IsNaN(AC) && !math.IsNaN(a) && !math.IsNaN(c) {
		return areaAngle(AC, a, c)
	}
	if !math.IsNaN(BC) && !math.IsNaN(b) && !math.IsNaN(c) {
		return areaAngle(BC, b, c)
	}

	if !math.IsNaN(AB) && !math.IsNaN(AC) && !math.IsNaN(a) {
		return areaTWOAngle(AB, AC, a)
	}
	if !math.IsNaN(AC) && !math.IsNaN(BC) && !math.IsNaN(c) {
		return areaTWOAngle(AC, BC, c)
	}
	if !math.IsNaN(AB) && !math.IsNaN(BC) && !math.IsNaN(b) {
		return areaTWOAngle(AB, BC, b)
	}

	return math.NaN()
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

	if !math.IsNaN(area(a, b, c, AB, AC, BC)) {
		fmt.Print(area(a, b, c, AB, AC, BC))
	} else {
		fmt.Print("Unable to calculate area insufficient data")
	}
}
