package main

import "fmt"

func main() {

	for i := 1; i <= 1000; i++ {
		fmt.Printf("%d  -  %E \n", i, fib64(i))
	}
}

func fib(N int) int {
	if N <= 2 {
		return 1
	}

	//N>= 3

	nn := 1  // n2
	n_1 := 1 // n1
	for i := 3; i <= N; i++ {
		last := nn
		nn = last + n_1
		n_1 = last
	}

	return nn
}

func fib64(N int) float64 {
    if N <= 2 {
        return 1
    }

    //N>= 3

    nn := float64(1)  // n2
    n_1 := float64(1) // n1
    for i := 3; i <= N; i++ {
        last := nn
        nn = last + n_1
        n_1 = last
    }

    return nn
}
