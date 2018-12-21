package main

import (
	"fmt"
	"time"
	"github.com/noragalvin/learning-go/stringutil"
	"math/rand"
	"math"
)

func add(x int, y int) int {
	return x + y;
}

func main() {
	fmt.Printf(stringutil.Reverse("oG olleH tseT") + "\n");

	fmt.Println("Welcome to the playground!")

	fmt.Println("The time is", time.Now())

	fmt.Println("My favorite number is", rand.Intn(10))

	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))

	fmt.Println("PI:", math.Pi)

	fmt.Printf("Sum of %d and %d equal %d.\n", 1, 2, add(1, 2))
}