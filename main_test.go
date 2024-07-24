package main

import (
	"fmt"
	"testing"
	"time"
)

func TestChallenge(t *testing.T) {
	start := time.Now()
	main()
	elapsed := time.Since(start)

	fmt.Printf("Took %vs", elapsed.Seconds())
}

func TestParseIntMult10(t *testing.T) {
	validations := []struct {
		source   []byte
		expected int
	}{
		{[]byte("1.0"), 10},
		{[]byte("-1.1"), -11},
		{[]byte("102.5"), 1025},
	}

	for _, v := range validations {
		out := parseIntMult10(v.source)
		if v.expected != out {
			fmt.Printf("Validation failed. Expected %d, got %d\n", v.expected, out)
			t.Fail()
		}
	}
}
