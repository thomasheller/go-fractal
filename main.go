package main

import (
	"flag"
	"fmt"
	"log"
)

const (
	MaxIters      = 5
	SegmentLength = 16
	ScreenWidth   = 100
)

func main() {
	i := flag.Int("i", MaxIters, fmt.Sprintf("number of iterations (between 1 and %d)", MaxIters))
	flag.Parse()

	if *i < 1 || *i > MaxIters {
		log.Fatalf("Iterations must be >= 1 and <= %d", MaxIters)
	}

	f := &Fractal{}
	f.Draw(*i)
	f.Print()
}
