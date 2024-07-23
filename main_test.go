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
