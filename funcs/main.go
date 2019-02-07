package main

import (
	"fmt"
)

type binFunc func(int, int) int

type op struct {
	name string
	fn   func(int, int) int
}

type a func(int, int) int

func add(x, y int) int {
	return x + y
}

func (f binFunc) Error() string {
	return "binFunc error"
}

func main() {
	/*
		rand.Seed(time.Now().Unix())
		x := 5
				fn := func() {
					fmt.Printf("hello, my age is :%d", x)
				}

		// fn()

		//create a slice of functions
		fns := []binFunc{
			func(x, y int) int { return x + y },
			func(x, y int) int { return x - y },
			func(x, y int) int { return x * y },
			func(x, y int) int { return x / y },
		}

		//pick one of those funcs at random
		fnP := fns[rand.Intn(len(fns))]
		//execute it
		x, y := 12, 5
		fmt.Println(fnP(x, y))

		ops := []op{
			{"add", func(x, y int) int { return x + y }},
			{"sub", func(x, y int) int { return x - y }},
			{"mul", func(x, y int) int { return x * y }},
			{"div", func(x, y int) int { return x / y }},
		}

		//pick one of those ops at random
		o := ops[rand.Intn(len(ops))]

		fmt.Printf("%s: %d", o.name, o.fn(x, y))
	*/
	var err error

	err = binFunc(a)
	fmt.Println(err)

}
