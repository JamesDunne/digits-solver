package main

import "fmt"

var opFunc = []func(a, b int) int{
	func(a, b int) int { return a + b },
	func(a, b int) int { return a - b },
	func(a, b int) int { return a * b },
	func(a, b int) int { return a / b },
}
var opValid = []func(a, b int) bool{
	func(a, b int) bool { return true },
	func(a, b int) bool { return a >= b },
	func(a, b int) bool { return true },
	func(a, b int) bool { return a%b == 0 },
}

type calc struct {
	a  int
	b  int
	op int
}

func (c calc) String() string {
	switch c.op {
	case 0:
		return fmt.Sprintf("(%d + %d)", c.a, c.b)
	case 1:
		return fmt.Sprintf("(%d - %d)", c.a, c.b)
	case 2:
		return fmt.Sprintf("(%d * %d)", c.a, c.b)
	case 3:
		return fmt.Sprintf("(%d / %d)", c.a, c.b)
	default:
		panic(fmt.Errorf("unknown operation"))
	}
}

func main() {
	target := 144
	digits := [6]int{
		2,
		7,
		9,
		10,
		11,
		25,
	}

	sln, solved := solve(target, digits, nil)
	if !solved {
		fmt.Println("failed to solve")
		return
	}

	for _, c := range sln {
		fmt.Println(c)
	}
}

func solve(target int, digits [6]int, hist []calc) (sln []calc, solved bool) {
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

			for o := 0; o < 4; o++ {
				if !opValid[o](a, b) {
					continue
				}

				v := opFunc[o](a, b)
				if v == target {
					sln = make([]calc, len(hist), len(hist)+1)
					copy(sln, hist)
					sln = append(sln, calc{a, b, o})
					solved = true
					return
				}

				var newdigits [6]int
				copy(newdigits[:], digits[:])
				newdigits[n] = -1
				newdigits[i] = v

				newhist := make([]calc, len(hist), len(hist)+1)
				copy(newhist, hist)
				newhist = append(newhist, calc{a, b, o})
				if sln, solved = solve(target, newdigits, newhist); solved {
					return
				}
			}
		}
	}

	return nil, false
}
