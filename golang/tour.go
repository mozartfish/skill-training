package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"time"
)

// The following code is from the Google Tour of Go tutorial
type Vertex struct {
	// X int
	// Y int
	// Lat, Long float64
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func add1(x int, y int) int {
	return x + y
}
func add2(x, y int) int {
	return x + y
}
func swap(x, y string) (string, string) {
	return y, x
}
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}
func pow2(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

// Exercise: Loops and Function
func newtonMethod(x float64) float64 {
	z := 1.0
	for i := 1; i <= 10; i++ {
		fmt.Println(z)
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

// Slices Exercise
func Pic(dx, dy int) [][]uint8 {
	result := make([][]uint8, dy)
	for x := range result {
		result[x] = make([]uint8, dx)
		for y := range result[x] {
			result[x][y] = uint8((x + y) / 2)
			result[x][y] = uint8(x * y)
			//result[x][y] = uint8(x ^ y)
		}
	}
	return result
}

// Fibonacci Closure Exercise
// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	total, nextTotal := 0, 1
	return func() int {
		result := total
		total, nextTotal = nextTotal, nextTotal+result
		return result
	}
}

// var c, python, java bool
// var i, j int = 1, 2
// const Pi = 3.14

var (
	// Struct Literal
	// v1       = Vertex{1, 2}
	// v2       = Vertex{X: 1}
	// v3       = Vertex{} // X: 0, Y : 0
	// p1       = &Vertex{1, 2}
	powSlice = []int{1, 2, 4, 8, 16, 32, 64, 128}
	m        map[string]Vertex
)

func main() {
	// Hello World
	fmt.Println("Hello, 世界")
	// Packages
	fmt.Println("My favorite number is", rand.Intn(10))
	// Factored import and print verbs
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	// Exported Names
	fmt.Println(math.Pi)
	// Functions
	add1 := add1(42, 13)
	add2 := add2(35, 13)
	s1, s2 := swap("hello", "world")
	fmt.Println("function in go", add1)
	fmt.Println("shared type function:", add2)
	fmt.Println("multiple result function:", s1, s2)
	fmt.Printf("Named return values: ")
	fmt.Println(split(17))
	// Variables
	// var i int
	// fmt.Println("Variables using var:", i, c, python, java)
	// Initialized Variables
	// var c, python, java = true, false, "no!"
	// fmt.Println("Initialized variables:", i, j, c, python, java)
	// Short Variable Declarations
	// var i, j int = 1, 2
	// k := 3
	// c, python, java := true, false, "no!"
	// fmt.Println("short variable declaration:", i, j, k, c, python, java)
	// // Constants
	// fmt.Println("Constants:", Pi)

	// Loops
	// For
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("for loop demo:", sum)
	sum2 := 1
	for sum2 < 1000 {
		sum2 += sum2
	}
	fmt.Println("for loop without init and post statements:", sum2)
	// While
	sum3 := 1
	for sum3 < 1000 {
		sum3 += sum3
	}
	fmt.Println("while loop:", sum3)
	// Infinite
	// for {

	// }
	// Conditional statements
	fmt.Println("if statement demo:", sqrt(2), sqrt(-4))
	// switch
	fmt.Print("Go run on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		// freebsd, openbsd
		// plan9, windows
		fmt.Printf("%s.\n", os)
	}
	// Switch statement continued
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
	// switch no case
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	// Defer statements
	defer fmt.Println("defer printing:", "world")
	fmt.Println("defer printing Hello")

	// Pointers
	// var p *int
	// i := 42
	// p = &i
	// fmt.Println(*p) // read i through the pointer p
	// *p = 21         // set i through the pointer p

	// Structs
	// fmt.Println("This is a struct:", Vertex{1, 2})
	// vVertex := Vertex{1, 2}
	// vVertex.X = 14
	// fmt.Println("vVertex set x:", vVertex.X)

	// Pointers to structs
	// v := Vertex{1, 2}
	// p := &v
	// p.X = 1e9
	// fmt.Println("Pointer to struct golang:", p.X)

	// fmt.Println(v1, p1, v2, v3)

	// Slices - using make
	a := make([]int, 5)
	printSlice("a", a)

	// Range Loops
	for index, value := range powSlice {
		fmt.Printf("2**%d = %d\n", index, value)
	}

	// Maps
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	// Methods

}

// Packages
// - Every go program is made up of packages
// - Programs start running in package main
// - Convenvention the package name is the same as the last
// element of the import path

// Exported Names
// - In Go, a name is exported if it begins with a capital letter

// Named Return Values
// - Go's return values may be named. If so they are treated as variables defined
//  at the top of the function
//  - The names should be used to document the meaning of the return values
//  - A return statement without arguments returns the named return values
//  This is known as a naked return
//  - Naked return statements should be used only in short functions such as split()
//  - They can harm readbility for long functions

// Variables
// - The var statement declares a list of variables; as a in function arguments list
//  the type is last
// - A var statement can be at package or function level

// Short Variable Declarations
// - Inside a function the := syntax can be used in place of a var declaration
// with implicit type
// - Outside of a function every statement begins with a keyword(var, func, and so on)
// so := is not available

// Basic Types
// byte - alias for uint8
// rune - alias for int32 (represents a Unicode point)

// Defer
// - Defer statements defer execution of a function until the surrounding function returns
// - Defer function calls are pushed onto a stack and executed in LIFO order

// Pointers
// - Go has no pointer arithmetic

// Struct
// - A collection of fields
// Struct Literals
// - A struct literal denotes a newly allocated struct value by listing
// the values of its fields

// Slices
//  - Go version of a dynamic array
//  - Slices do not store data and can be thought of as a reference to an underlying array
//  - Slices can be created with the built in make function

// Methods
// - Go does not have classes. Can define methods on types
// - A method is a function with a special receiver argument
// - Receiver appears in it own argument list between func keyword and method name
