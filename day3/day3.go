package main

import (
	"fmt"
	"github.com/tomp/aoc-2017-go/util"
)

const (
	INPUT = 368078
)

type Position struct {
	loc  int // sequential location along the spiral
	ring int // ring number
	x    int // x coordinate
	y    int // y coordinate
}

var (
	ORIGIN = Position{1, 0, 0, 0}
)

type Direction struct {
	dx int // x increment
	dy int // y increment
}

var (
	RIGHT = Direction{1, 0}
	LEFT  = Direction{-1, 0}
	UP    = Direction{0, 1}
	DOWN  = Direction{0, -1}
)

func ringEnd(ring int) int {
	return ((2 * ring) + 1) * ((2 * ring) + 1)
}

func positionStream() (pStream chan Position) {
	pStream = make(chan Position)
	go func() {
		pos := ORIGIN
		end := ringEnd(pos.ring)
		dir := RIGHT
		for {
			pStream <- pos
			pos.loc += 1
			if pos.loc > end {
				pos.ring += 1
				end = ringEnd(pos.ring)
				pos.x += dir.dx
				dir = UP
			} else {
				if util.IntAbs(pos.x*dir.dx) == pos.ring || util.IntAbs(pos.y*dir.dy) == pos.ring {
					if dir == RIGHT {
						dir = UP
					} else if dir == UP {
						dir = LEFT
					} else if dir == LEFT {
						dir = DOWN
					} else if dir == DOWN {
						dir = RIGHT
					}
				}
				pos.x += dir.dx
				pos.y += dir.dy
			}
		}
	}()
	return
}

func distanceToOrigin(start int) (dist int) {
	for p := range positionStream() {
		if p.loc == start {
			dist = util.IntAbs(p.x) + util.IntAbs(p.y)
			break
		}
	}
	return
}

type locValue struct {
	loc, value int
}

type xyCoord struct {
	x, y int
}

func valueStream() (vStream chan locValue) {
	c := make(chan locValue)
	prev := make(map[xyCoord]int)

	go func() {
		for pos := range positionStream() {
			value := 0
			for xv := pos.x - 1; xv < pos.x+2; xv += 1 {
				for yv := pos.y - 1; yv < pos.y+2; yv += 1 {
					if !(xv == pos.x && yv == pos.y) {
						value += prev[xyCoord{xv, yv}]
					}
				}
			}
			if value == 0 {
				value = 1
			}
			prev[xyCoord{pos.x, pos.y}] = value
			c <- locValue{pos.loc, value}
		}
	}()

	return c
}

func solve(start int) int {
	return distanceToOrigin(start)
}

func solve2(n int) int {
	for v := range valueStream() {
		if v.value > n {
			return v.value
		}
	}
	return 0
}

func main() {

	// Part 1
	result := solve(INPUT)
	fmt.Printf("1) result is %d\n", result)

	// Part 2
	result = solve2(INPUT)
	fmt.Printf("2) result is %d\n", result)
}
