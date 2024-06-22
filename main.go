package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var filename = "./measurements.txt"

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

func main() {
	file, err := os.Open(filename)
	if err != nil {
		panic("error opening file: " + err.Error())
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		city, measure, _ := strings.Cut(scanner.Text(), ";")

		mf, _ := strconv.ParseFloat(measure, 64)
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
