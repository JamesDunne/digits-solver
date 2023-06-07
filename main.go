package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

var opFunc = []func(a, b int) int{
	func(a, b int) int { return a + b },
	func(a, b int) int { return a * b },
	func(a, b int) int { return a - b },
	func(a, b int) int { return a / b },
}
var opValid = []func(a, b int) bool{
	func(a, b int) bool { return true },
	func(a, b int) bool { return true },
	func(a, b int) bool { return a >= b },
	func(a, b int) bool { return a%b == 0 },
}

type calc struct {
	a  int
	b  int
	op int
}

func (c calc) String() string {
	var fmtStr string
	switch c.op {
	case 0:
		fmtStr = "%3d + %3d = %3d"
		break
	case 1:
		fmtStr = "%3d * %3d = %3d"
		break
	case 2:
		fmtStr = "%3d - %3d = %3d"
		break
	case 3:
		fmtStr = "%3d / %3d = %3d"
		break
	default:
		panic(fmt.Errorf("unknown operation"))
	}

	return fmt.Sprintf(fmtStr, c.a, c.b, opFunc[c.op](c.a, c.b))
}

func main() {
	if len(os.Args) <= 6 {
		fmt.Println("digits <solution> <digits 1..6>")
		return
	}

	var err error

	var target int
	target, err = strconv.Atoi(os.Args[1])
	digits := [6]int{}
	for i := 0; i < 6; i++ {
		digits[i], err = strconv.Atoi(os.Args[2+i])
	}
	_ = err

	fmt.Printf("solving for %d using %v\n", target, digits)
	solved := false
	solutions := make([][]calc, 0, 100)
	solution := func(sln []calc) {
		solved = true
		solutions = append(solutions, sln)
	}
	solve(target, digits, nil, solution)

	if !solved {
		fmt.Println("failed to solve")
		return
	}

	sort.Slice(solutions, func(i, j int) bool {
		return len(solutions[i]) > len(solutions[j])
	})

	for _, sln := range solutions {
		fmt.Println("solution:")
		for _, c := range sln {
			fmt.Printf("  %s\n", c)
		}
	}
}

func solve(target int, digits [6]int, hist []calc, solution func(sln []calc)) {
	calcs := make([]calc, 0, 6*6)

	// commutative operations (+, *):
	for o := 0; o < 2; o++ {
		for n := 0; n < 5; n++ {
			a := digits[n]
			if a <= 0 {
				continue
			}

			for i := n + 1; i < 6; i++ {
				b := digits[i]
				if b <= 0 {
					continue
				}

				if !opValid[o](a, b) {
					continue
				}

				calcs = append(calcs, calc{n, i, o})
			}
		}
	}

	// non-commutative operations (-, /):
	for o := 2; o < 4; o++ {
		for n := 0; n < 6; n++ {
			a := digits[n]
			if a <= 0 {
				continue
			}

			for i := 0; i < 6; i++ {
				if n == i {
					continue
				}

				b := digits[i]
				if b <= 0 {
					continue
				}

				if !opValid[o](a, b) {
					continue
				}

				calcs = append(calcs, calc{n, i, o})
			}
		}
	}

	for _, c := range calcs {
		n, i, o := c.a, c.b, c.op

		a, b := digits[n], digits[i]

		v := opFunc[o](a, b)

		if v == target {
			sln := make([]calc, len(hist), len(hist)+1)
			copy(sln, hist)
			sln = append(sln, calc{a, b, o})
			solution(sln)
			continue
		}

		var newdigits [6]int
		copy(newdigits[:], digits[:])
		newdigits[n] = -1
		newdigits[i] = v

		newhist := make([]calc, len(hist), len(hist)+1)
		copy(newhist, hist)
		newhist = append(newhist, calc{a, b, o})
		solve(target, newdigits, newhist, solution)
	}
}
