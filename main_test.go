package main

import (
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
	p.putStr(s)
	in := p.in[0]
	p.in = p.in[1:]
	return in
}

type randPure struct {
	n int
}

func (r *randPure) randomR(rr *randRange) int {
	r.n = (rr.lo + rr.hi) / 2
	return r.n
}

func assert(t *testing.T, g, e interface{}) {
	if g != e {
		t.Error(g)
	}
}

func TestIntro(t *testing.T) {
	ctx := &consolePure{in: []string{"x", "-1", "1", "y", "0", "10"}}
	r := intro(ctx)
	assert(t, ctx.out, `Guessing Game
low: x is not an int
low: low must be >= 0
low: high: y is not an int
high: high must be >= low
high: `)
	assert(t, r.lo, 1)
	assert(t, r.hi, 10)
}

func TestPlay(t *testing.T) {
	ctx := struct {
		*consolePure
		*randPure
	}{&consolePure{in: []string{"x", "3", "7", "5"}}, &randPure{}}
	r := &randRange{1, 10}
	play(ctx, r)
	assert(t, ctx.n, 5)
	assert(t, ctx.out, `guess: x is not an int
guess: higher
guess: lower
guess: 5 is correct
`)
}