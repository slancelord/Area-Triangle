// go mod init -> go.mod

// go test -cover -coverprofile=c.out && go tool cover -html=c.out -o cover.html

package main

import (
	"math"
	"testing"
)

func TestArea_Table(tst *testing.T) {
	tests := []struct {
		a, b, c, AB, AC, BC float64
		want                float64
	}{ /*
			{1, 2, 2, math.NaN(), math.NaN(), math.NaN(), 0.968245836551854255},
			{1, math.NaN(), math.NaN(), 45, 45, 90, 0.24999999999999994},
			{math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN(), math.NaN()},
			{1, 1, 1, 1, 1, 1, -1},
			{1, 2, 21, 1, 1, 1, -1},
			{1, 2, 2, 189, 1, 1, -1},
			{1, 2, 2, 90, 90, math.NaN(), -1},
			{1, 2, 2, math.NaN(), 189, math.NaN(), -1},*/
		{1, 2, 2, math.NaN(), 75, 20, 0.968246},
		//{3, math.NaN(), 2, 90, math.NaN(), math.NaN(), -1},
		//{math.NaN(), 2, 3, math.NaN(), 90, math.NaN(), -1},
		//{2, math.NaN(), 3, math.NaN(), math.NaN(), 90, -1},
	}

	for _, t := range tests {
		total, err := area(t.a, t.b, t.c, t.AB, t.AC, t.BC)
		if err == nil {
			if (total != t.want) || (!math.IsNaN(total) && math.IsNaN(t.want)) || (math.IsNaN(total) && !math.IsNaN(t.want)) {
				tst.Errorf("area(%f,%f,%f,%f,%f,%f): total=%f; want=%f", t.a, t.b, t.c, t.AB, t.AC, t.BC, total, t.want)
			} else {
				tst.Log("ok")
			}
		} else {
			tst.Log(err)
		}

	}

}
