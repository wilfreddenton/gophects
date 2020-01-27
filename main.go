package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type consoleEff interface {
	putStr(string)
	putStrLn(string)
	getLine(string) string
}

type consoleIO struct{}

func (p *consoleIO) putStr(s string) {
	fmt.Print(s)
}

func (p *consoleIO) putStrLn(s string) {
	fmt.Println(s)
}

func (p *consoleIO) getLine(s string) string {
	fmt.Print(s)
	var in string
	fmt.Scanln(&in)
	return in
}

type randRange struct {
	lo int
	hi int
}

type randEff interface {
	randomR(*randRange) int
}

type randIO struct{}

func (r *randIO) randomR(rr *randRange) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(rr.hi-rr.lo) + rr.lo
}

func intro(ctx consoleEff) *randRange {
	ctx.putStrLn("Guessing Game")

	var loS, hiS string
	var lo, hi int
	var err error
	for {
		loS = ctx.getLine("low: ")
		lo, err = strconv.Atoi(loS)
		if err != nil {
			ctx.putStrLn(loS + " is not an int")
		} else if lo < 0 {
			ctx.putStrLn("low must be >= 0")
		} else {
			break
		}
	}

	for {
		hiS = ctx.getLine("high: ")
		hi, err = strconv.Atoi(hiS)
		if err != nil {
			ctx.putStrLn(hiS + " is not an int")
		} else if hi < lo {
			ctx.putStrLn("high must be >= low")
		} else {
			break
		}
	}

	return &randRange{lo, hi}
}

type playEff interface {
	consoleEff
	randEff
}

func play(ctx playEff, r *randRange) {
	n := ctx.randomR(r)

	for {
		gS := ctx.getLine("guess: ")
		g, err := strconv.Atoi(gS)
		if err != nil {
			ctx.putStrLn(gS + " is not an int")
			continue
		}

		if g < n {
			ctx.putStrLn("higher")
		} else if g > n {
			ctx.putStrLn("lower")
		} else {
			ctx.putStrLn(gS + " is correct")
			break
		}
	}
}

func main() {
	ctx := struct {
		*consoleIO
		*randIO
	}{&consoleIO{}, &randIO{}}

	r := intro(ctx)
	play(ctx, r)
}
