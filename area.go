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
		return
	}
	return

}

func area(a, b, c, AB, AC, BC float64) (S float64) {
	if a == a && b == b && c == c {
		p := (a + b + c) / 2
		S = math.Sqrt(p * (p - a) * (p - b) * (p - c))
	}
	return
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

	fmt.Print(area(a, b, c, AB, AC, BC))

}
