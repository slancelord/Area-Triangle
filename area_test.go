// go mod init -> go.mod

// go test -cover -coverprofile=c.out
// go tool cover -html=c.out -o cover.html

package main

import (
	"math"
	"testing"
)

func equals(numA, numB float64) bool {
	if (math.IsNaN(numA)) && (math.IsNaN(numB)) {
		return true
	}

	const eps = 1e-5

	return math.Abs(numA-numB) < eps
}
func TestArea_Table(tst *testing.T) {
	tests := []struct {
		a, b, c, AB, AC, BC float64
		want                float64
	}{
		{1, 2, 2, math.NaN(), math.NaN(), math.NaN(), 0.968245836551854255},
		{1, math.NaN(), math.NaN(), 45, 45, 90, 0.24999999999999994},
		{math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN()},
		{1, 1, 1, 1, 1, 1, -1},
		{1, 2, 21, 1, 1, 1, -1},
		{1, 2, 2, 189, 1, 1, -1},
		{1, 2, 2, 90, 90, math.NaN(), -1},
		{1, 2, 2, math.NaN(), 189, math.NaN(), -1},
		{math.NaN(), 2, 2, -90, math.NaN(), math.NaN(), -1},

		//-----------------------------------------------------------------------------------
		{math.NaN(), math.NaN(), 1, 1, 2, math.NaN(), 0.05232798522331311},
		{math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN()},
		{1, math.NaN(), 3, 75, math.NaN(), math.NaN(), 1.4967326577910303},
		{100, 2, math.NaN(), math.NaN(), math.NaN(), 100, 98.11965804559631},
		{1, 3, math.NaN(), 75, math.NaN(), math.NaN(), 1.4488886337076199},
		{math.NaN(), math.NaN(), 2, 30, 70, math.NaN(), 3.7016663135932935},
		{1, math.NaN(), math.NaN(), 2, 3, math.NaN(), 0.010478},
		{1, 2, 2, math.NaN(), 75, 20, 0.968245836551854255},

		{math.NaN(), math.NaN(), math.NaN(), 200, 20, math.NaN(), -1},
		{math.NaN(), math.NaN(), math.NaN(), 200, math.NaN(), math.NaN(), -1},
		{math.NaN(), math.NaN(), math.NaN(), 9, 1, 1, -1},
		{1, 2, 3, 75, 70, math.NaN(), -1},
		{1, 2, 3, 75, math.NaN(), 70, -1},
		{-1, -2, 2, math.NaN(), math.NaN(), math.NaN(), -1},
		{3, math.NaN(), 2, 90, math.NaN(), math.NaN(), -1},
		{math.NaN(), 2, 3, math.NaN(), 90, math.NaN(), -1},
		{2, math.NaN(), 3, math.NaN(), math.NaN(), 90, -1},
	}

	for _, t := range tests {
		total, err := area(t.a, t.b, t.c, t.AB, t.AC, t.BC)
		if err == nil {
			if !equals(total, t.want) {
				tst.Errorf("area(%f,%f,%f,%f,%f,%f): total=%f; want=%f", t.a, t.b, t.c, t.AB, t.AC, t.BC, total, t.want)
			} else {
				tst.Log("ok")
			}
		} else {
			tst.Log(err)
		}

	}

}
