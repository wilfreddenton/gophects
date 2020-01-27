package main

import (
	"fmt"
	"testing"
)

type consolePure struct {
	in  []string
	out string
}

func (p *consolePure) putStr(s string) {
	p.out += s
}

func (p *consolePure) putStrLn(s string) {
	p.out += s + "\n"
}

func (p *consolePure) getLine(s string) string {
	in := p.in[0]
	p.putStrLn(fmt.Sprintf("%s%v", s, in))
	p.in = p.in[1:]
	return in
}

type randPure struct {
	n int
}

func (r *randPure) randomR(lo, hi int) int {
	r.n = (lo + hi) / 2
	return r.n
}

func assert(t *testing.T, g, e interface{}) {
	if g != e {
		t.Error(g)
	}
}

func TestGetTurns(t *testing.T) {
	ctx := &consolePure{in: []string{"t", "0", "1"}}
	turns := getTurns(ctx)
	assert(t, turns, 1)
	assert(t, ctx.out, `turns: t
t is not an int
turns: 0
turns must be > 0
turns: 1
`)
}

func TestGetLow(t *testing.T) {
	ctx := &consolePure{in: []string{"x", "-1", "1"}}
	lo := getLow(ctx)
	assert(t, lo, 1)
	assert(t, ctx.out, `low: x
x is not an int
low: -1
low must be >= 0
low: 1
`)
}

func TestGetHigh(t *testing.T) {
	ctx := &consolePure{in: []string{"y", "0", "10"}}
	hi := getHigh(ctx, 1)
	assert(t, hi, 10)
	assert(t, ctx.out, `high: y
y is not an int
high: 0
high must be >= low
high: 10
`)
}

func TestIntro(t *testing.T) {
	ctx := &consolePure{in: []string{"5", "1", "10"}}
	r := intro(ctx)
	assert(t, r.turns, 5)
	assert(t, r.lo, 1)
	assert(t, r.hi, 10)
	assert(t, ctx.out, `Guessing Game
turns: 5
low: 1
high: 10
`)
}

func TestPlayWin(t *testing.T) {
	ctx := struct {
		*consolePure
		*randPure
	}{&consolePure{in: []string{"x", "3", "7", "5"}}, &randPure{}}
	r := &settings{3, 1, 10}
	play(ctx, r)
	assert(t, ctx.n, 5)
	assert(t, ctx.out, `guess: x
x is not an int
guess: 3
higher
guess: 7
lower
guess: 5
5 is correct
`)
}

func TestPlayLoss(t *testing.T) {
	ctx := struct {
		*consolePure
		*randPure
	}{&consolePure{in: []string{"x", "3", "7"}}, &randPure{}}
	r := &settings{1, 1, 10}
	play(ctx, r)
	assert(t, ctx.out, `guess: x
x is not an int
guess: 3
higher
game over
`)
}
