package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var filename = "./measurements.txt"

// var filename = "./create/measurements.txt"

type Station struct {
	min, max, sum float64
	n             int
}

func newStation(val float64) *Station {
	return &Station{
		min: val,
		max: val,
		n:   1,
		sum: val,
	}
}

var stations = map[string]*Station{}

func parseFloat(measure []byte) float64 {
	var dotPos int
	for i, b := range measure {
		if b == '.' {
			dotPos = i
			break
		}
	}

	mf := 0.0
	i := 0
	neg := false
	if measure[0] == '-' {
		neg = true
		i++
	}

	for d := dotPos - i - 1; d >= 0; d-- {
		mf += float64(measure[i]-'0') * math.Pow10(d)
		i++
	}
	mf += float64(measure[dotPos+1]-'0') * math.Pow10(-1)
	if neg {
		mf = -mf
	}

	return mf
}

func main() {
	file, err := os.Open(filename)
	if err != nil {
		panic("error opening file: " + err.Error())
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Bytes()
		var sepPos int
		for i, b := range line {
			if b == ';' {
				sepPos = i
				break
			}
		}
		city := string(line[:sepPos])
		measure := line[sepPos+1:]

		mf := parseFloat(measure)
		if s, ok := stations[city]; ok {
			if mf > s.max {
				s.max = mf
			}
			if mf < s.min {
				s.min = mf
			}
			s.sum += mf
			s.n++
		} else {
			stations[city] = newStation(mf)
		}
	}

	for city, s := range stations {
		fmt.Printf("%s;%.1f;%.1f;%.1f\n", city, s.min, s.sum/float64(s.n), s.max)
	}
}
