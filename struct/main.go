package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

func circleArea(c Circle) float64 {
	return math.Pi * c.r * c.r
}

func circleArea2(c *Circle) float64 {
	fmt.Println(c)

	//c is a Pointer, meaning it holds the
	//memory address of the Struct being
	//passed in (call by reference)
	//To access the content of c we use *c
	//To access whats in *c we write (*c).field
	fmt.Println((*c).x, (*c).y, (*c).r)

	//But to make life easier golang allows for
	//a short hand
	//fmt.Println(c.x,c.y,c.r)
	//is equal to
	//fmt.Println((*c).x,(*c).y,(*c).r)
	fmt.Println(c.x, c.y, c.r)

	return math.Pi * c.r * c.r
}

type EmptyStruct struct{}
type MyMap map[string]EmptyStruct

type Any interface{}

type MyList []Any

func (l *MyList) printLC() {
	fmt.Println(len(*l), cap(*l))
}

type Rectangle struct {
	Width, Height float64
}

//
// Pointer receiver
func (r *Rectangle) SetWidth(width float64) *Rectangle {
	r.Width = width
	return r
}

func (r *Rectangle) SetHeight(height float64) *Rectangle {
	r.Height = height
	return r
}

//Value Receiver
func (r Rectangle) Scale(factor float64) {
    r.Width *= factor
    r.Height *= factor
}

func main() {

	//1. b := new(Circle) is as good as saying
	//2. b := &Circle{} which is as good as saying
	//3. a := Circle{}
	//	 b := &a
	//All 3 return pointers to a struct, aka the b's hold the address of the struct
	//For now we will use
	b := new(Circle)

	//This will print the memory address of the pointer b
	fmt.Println(&b)

	//This will print the contents of the pointer b,
	//which is the Memory address of the struct
	fmt.Println(b)

	//This will print the contents of the
	//memory address being pointed by b
	//which is the actual Struct content
	fmt.Println(*b)

	c := Circle{x: 0, y: 0, r: 5}
	fmt.Println(circleArea(c))
	fmt.Println(circleArea2(&c))


	set := new(MyMap) // new always returns a pointer; it allocates memory for the specified type
	// In every situation it also setsup Zero values EXCEPT Maps, Slices and Channels
	fmt.Println()
	fmt.Println(set)
	*set = make(MyMap)
	fmt.Println(set)
	(*set)["apple"] = EmptyStruct{}
	(*set)["banana"] = EmptyStruct{}
	
	// This below is the standard way of creating a map
	// set := map[string]EmptyStruct{}
	// set["apple"] = EmptyStruct{}
	// set["banana"] = EmptyStruct{}

	l := MyList{1, "two", 3}
	fmt.Println(l, set)
	(&l).printLC() // is equal to l.printLC()

	rect := new(Rectangle) //returns pointer
	rect.SetWidth(10).SetHeight(5) // Same as (&rect).SetWidth(10).SetHeight(5)
	fmt.Println("Rectangle:", rect)

	rect2 := Rectangle{}
	rect2.SetWidth(10).SetHeight(5) // Chaining methods

	(&rect2).Scale(2) // SAME as rect2.Scale(2)
	fmt.Println(rect2)
	
	mychan := make(chan int)
	fmt.Println()
	fmt.Println(mychan)
}
