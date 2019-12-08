package main

import (
	"fmt"
	"math"
	"strings"
)

const (
	MaxIters      = 5
	SegmentLength = 16
	ScreenWidth   = 100
)

type Fractal struct {
	sb strings.Builder
}

func (f *Fractal) Draw(it int) {
	f.Padding(ScreenWidth)
	f.Newline()

	for i := MaxIters; i > it; i-- {
		n := f.Num(i)

		for y := 0; y < 2*SegmentLength/n; y++ {
			f.Padding(ScreenWidth)
			f.Newline()
		}
	}

	for i := it; i > 0; i-- {
		f.Iter(i)
	}
}

func (f *Fractal) Iter(it int) {
	n := f.Num(it)

	f.Branches(n, it, SegmentLength/n)
	f.Stem(n, SegmentLength/n)
}

func (f *Fractal) Num(it int) int {
	return int(math.Pow(float64(2), float64(it-1)))
}

func (f *Fractal) Branches(n, it, side int) {
	for y := 1; y <= side; y++ {
		f.Border()

		for i := 1; i <= n; i++ {
			f.Padding(side - 1)
			f.BranchLeft(y, side)
			f.Blank()
			f.BranchRight(y, side)
			f.Padding(side)
		}

		f.Border()
		f.Newline()
	}
}

func (f *Fractal) Stem(n, side int) {
	for y := 1; y <= side; y++ {
		f.Border()

		for i := 1; i <= n; i++ {
			f.Padding(side - 1)
			f.StemPart(side)
			f.Padding(side)
		}

		f.Border()
		f.Newline()
	}
}

func (f *Fractal) BranchLeft(y, side int) {
	for x := 1; x <= side; x++ {
		if x == y {
			f.Mark()
		} else {
			f.Blank()
		}
	}
}

func (f *Fractal) BranchRight(y, side int) {
	for x := 1; x <= side; x++ {
		if x == side-y+1 {
			f.Mark()
		} else {
			f.Blank()
		}
	}
}

func (f *Fractal) StemPart(side int) {
	for x := 1; x <= side; x++ {
		f.Blank()
	}

	f.Mark()

	for x := 1; x <= side; x++ {
		f.Blank()
	}
}

func (f *Fractal) Padding(width int) {
	for x := 0; x < width; x++ {
		f.Blank()
	}
}

func (f *Fractal) Border() {
	f.Padding(ScreenWidth - MaxIters*SegmentLength - 2)
}

func (f *Fractal) Blank() {
	f.sb.WriteString("_")
}

func (f *Fractal) Mark() {
	f.sb.WriteString("1")
}

func (f *Fractal) Newline() {
	f.sb.WriteString("\n")
}

func (f *Fractal) Print() {
	fmt.Print(f.sb.String())
}
