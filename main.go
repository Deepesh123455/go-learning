// package main

// import "fmt"

// func one() int {
// 	a := 10
// 	return a
// }

// func two() *int {
// 	b := 20
// 	return &b
// }

// func main() {
// 	a1 := one()
// 	a2 := two()
// 	fmt.Println(a1, *a2)

// }

// package main

// import "fmt"

// func one(a *int) {
// 	b := 10
// 	*a = b

// }

// func main() {
// 	a := 0
// 	one(&a)
// 	fmt.Println(a)
// }

package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println("The square root of 36 will be", math.Sqrt(36))
}
