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

type settings struct {
	turns int
	lo    int
	hi    int
}

type randEff interface {
	randomR(int, int) int
}

type randIO struct{}

func (r *randIO) randomR(lo, hi int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(hi-lo) + lo
}

func getTurns(ctx consoleEff) int {
	for {
		turnsS := ctx.getLine("turns: ")
		turns, err := strconv.Atoi(turnsS)
		if err != nil {
			ctx.putStrLn(turnsS + " is not an int")
		} else if turns < 1 {
			ctx.putStrLn("turns must be > 0")
		} else {
			return turns
		}
	}
}

func getLow(ctx consoleEff) int {
	for {
		loS := ctx.getLine("low: ")
		lo, err := strconv.Atoi(loS)
		if err != nil {
			ctx.putStrLn(loS + " is not an int")
		} else if lo < 0 {
			ctx.putStrLn("low must be >= 0")
		} else {
			return lo
		}
	}
}

func getHigh(ctx consoleEff, lo int) int {
	for {
		hiS := ctx.getLine("high: ")
		hi, err := strconv.Atoi(hiS)
		if err != nil {
			ctx.putStrLn(hiS + " is not an int")
		} else if hi < lo {
			ctx.putStrLn("high must be >= low")
		} else {
			return hi
		}
	}
}

func intro(ctx consoleEff) *settings {
	ctx.putStrLn("Guessing Game")

	turns := getTurns(ctx)
	lo := getLow(ctx)
	hi := getHigh(ctx, lo)
	return &settings{turns, lo, hi}
}

type playEff interface {
	consoleEff
	randEff
}

func play(ctx playEff, s *settings) {
	n := ctx.randomR(s.lo, s.hi)
	i := s.turns

	for i > 0 {
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
			return
		}

		i -= 1
	}

	ctx.putStrLn("game over")
}

func main() {
	ctx := struct {
		*consoleIO
		*randIO
	}{&consoleIO{}, &randIO{}}

	s := intro(ctx)
	play(ctx, s)
}
