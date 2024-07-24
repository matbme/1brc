package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const maxCityName = 64

var filename = "./measurements.txt"

// var filename = "./create/measurements.txt"

type Station struct {
	min, max, sum int
	n             int
}

func newStation(val int) *Station {
	return &Station{
		min: val,
		max: val,
		n:   1,
		sum: val,
	}
}

var stations = map[[maxCityName]byte]*Station{}

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
	mf += float64(measure[dotPos+1]-'0') * 0.1
	if neg {
		mf = -mf
	}

	return mf
}

func parseIntMult10(measure []byte) int {
	val := int(measure[len(measure)-1] - '0')

	negative := false
	lim := 0
	if measure[0] == '-' {
		negative = true
		lim = 1
	}

	mult := 10
	for ni := len(measure) - 3; ni >= lim; ni-- {
		val += int(measure[ni]-'0') * mult
		mult *= 10
	}

	if negative {
		return -val
	}
	return val
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
		city := [maxCityName]byte{}
		copy(city[:], line[:sepPos])
		measure := line[sepPos+1:]

		mf := parseIntMult10(measure)
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
		fmt.Printf("%s;%.1f;%.1f;%.1f\n", city, float64(s.min)/10.0, (float64(s.sum)/10.0)/float64(s.n), float64(s.max)/10.0)
	}
}
