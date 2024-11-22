package main

import "fmt"

var a = 1

func display(x string) {
	var f float32 = 10.1
	var (
		a int32  = int32(f)
		_ string = "hello"
	)
	_ = "hello there"
	d := 100
	_, _ = a, d
	fmt.Println(a, x)
}

func main() {
	display("Hello There!")

	// fmt.Println(x == nil)
	// x := []int{1, 2, 3}
	// // var x []int
	// y := []int{4, 5, 6}
	// a := append(x, y...)
	// fmt.Println(a)

	// var x []int
	// fmt.Println(x, len(x), cap(x))
	// x = append(x, 10)
	// fmt.Println(x, len(x), cap(x))
	// x = append(x, 20)
	// fmt.Println(x, len(x), cap(x))
	// x = append(x, 30)
	// fmt.Println(x, len(x), cap(x))
	// x = append(x, 40)
	// fmt.Println(x, len(x), cap(x))
	// x = append(x, 50)
	// fmt.Println(x, len(x), cap(x))

	// x := []int{1, 2, 3, 4}
	// y := x[:2]
	// z := x[1:]
	// x[1] = 20
	// y[0] = 10
	// z[1] = 30
	// fmt.Println("x:", x)
	// fmt.Println("y:", y)
	// fmt.Println("z:", z)

	x := make([]int, 0, 5)
	x = append(x, 1, 2, 3, 4)
	y := x[:2]
	z := x[2:]
	fmt.Println(cap(x), cap(y), cap(z))
	y = append(y, 30, 40, 50)
	fmt.Println(x, cap(x), y, cap(y), z, cap(z))
	x = append(x, 60)
	fmt.Println(x, cap(x), y, cap(y), z, cap(z))
	z = append(z, 70)
	fmt.Println(x, cap(x), y, cap(y), z, cap(z))
	// fmt.Println("x:", x)
	// fmt.Println("y:", y)
	// fmt.Println("z:", z)

	// var nilMap map[string]int
	var test = map[string][]string{}
	// fmt.Println(nilMap)
	test["a"] = []string{"a", "b", "c"}
	fmt.Println(test["a"], test["c"])
}
