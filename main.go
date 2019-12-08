package main

import (
	"flag"
	"fmt"
	"log"
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
