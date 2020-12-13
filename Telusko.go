// You can edit this code!
// Click here and start typing.
package main // so that othrs can also use go

// import "fmt" // here we are using the fmt package

// for multiple package
import (
	"fmt"
	"math"
	"time"
)

// Function
func add(x int, y int) (int, int) { // if both variables are of same type, "func add(x,y int)int{"
	out1 := x + y
	out2 := x - y
	return out1, out2
}

func calc(x int, y int) (out1, out2 int) { // if both variables are of same type, "func add(x,y int)int{"
	out1 = x + y
	out2 = x - y
	return
}

func main() { // this way we write function
	fmt.Println("Hello, GO")

	// Variables
	var num1 int // type compulsory when defining without default value
	var num2 = 3 // type not compulsory

	var num3, num4 int // multiple variables same line of same type

	result1 := 9 // var result1=9 // shorthand..

	// var num5 int, num6 int // not allowed

	num3 = 4
	num4 = 4

	num1 = 2

	var result = num1 + num2 + num3 + num4 + result1

	// Constant
	const num7 = 9 // constant
	// num7=1 // is not possible now

	fmt.Print(result, num7)

	// Function call
	x := 1
	y := 2

	z, w := add(x, y)
	al, bl := calc(x, y)

	fmt.Print(z, w)
	fmt.Print(al, bl)

	// Loop
	/*
		for{ // infinite loop
			fmt.Println("Hello")
		}
	*/

	// Like while loop
	i := 1
	for i <= 5 {
		fmt.Println("Akash")
		i++
	}

	// Like normal for loop
	count := 7
	for j := 0; j < count; j++ {
		fmt.Println("aksh", j)
	}

	//Exported Name
	findLarge()
	FindLarge()

	// Variable Scope
	scopeDemo()
	scopeDemo1()
	fmt.Println("in main", vs)

	var fsq float64 = 26

	// Math Package
	sqVal := math.Sqrt(fsq) // only type float64 is allowed
	fmt.Println("Sqrt of ", fsq, " = ", sqVal)

	// Formating
	fmt.Printf("The final answer is %g ,Thank you", sqVal)
	fmt.Println()
	// fmt.Printf("The final answer is %.2 ,Thank you",sqVal) // warning
	// fmt.Println()
	fmt.Printf("The final answer is %.2f ,Thank you", sqVal)
	fmt.Println()

	var CeilSqVal = math.Ceil(sqVal)
	var FloorSqVal = math.Floor(sqVal)

	fmt.Println("Ceil : ", CeilSqVal, " Floor : ", FloorSqVal)

	// If-Else
	numifelse := 7
	if numifelse <= 5 {
		fmt.Println("Hi ", numifelse)
	} else if numifelse <= 8 {
		fmt.Println("Middle ", numifelse)
	} else {
		fmt.Println("Bye ", numifelse)
	}

	// Switch, no break is required
	var checkSwitch = 9
	switch checkSwitch {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	default:
		fmt.Println("Default")
	}

	// Time package

	today := time.Now().Weekday()
	hola := today + 1
	fmt.Println(today+0, hola)
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today")
	case today + 6:
		fmt.Println("Saturday")
	case today + 2:
		fmt.Println("Too Far")
	default:
		fmt.Println("Too soon")
	}

	// Defer Function Call

	for i := 0; i < 10; i++ {
		defer fmt.Println(i) // defer function executes after the current function gets over and it stores them in a stack, LIFO manner
	}

	fmt.Println("Good Night")

	// Structure
	structGo()
}

// Exported Name
func findLarge() { // only accessible to inside package, not available to outside package
	fmt.Println("Small find large")
}

// FindLarge is an exported function
func FindLarge() { // accessible to inside as well as outside package
	fmt.Println("LArge find large")
}

var vs = 7 // package level
func scopeDemo() {
	fmt.Println("in scopeDemo package level", vs)
}

func scopeDemo1() {
	var vs = 8 // function level, function level>package level
	fmt.Println("in scopeDemo1 function level", vs)
}

// user defined type struct
type student struct {
	rollno int
	name   string
	marks  int
}

func structGo() {
	var s1 student = student{101, "Akash", 91}
	fmt.Println(s1)
	fmt.Println(s1.rollno, s1.name, s1.marks)

	var s2 student = student{rollno: 102, marks: 92}
	fmt.Println(s2)
	fmt.Println(s2.rollno, s2.name, s2.marks)
}

// NOTES
/*
// This is not allowed!.
func main()
{

}
*/

// In GO if we define a var or package and don't use it, then it will give error

// Don't use starting letter as capital letter for any thing in Go

/*
// Output:>go run Telusko.go
Hello, GO
22 93 -13 -1Akash
Akash
Akash
Akash
Akash
aksh 0
aksh 1
aksh 2
aksh 3
aksh 4
aksh 5
aksh 6
Small find large
LArge find large
in scopeDemo package level 7
in scopeDemo1 function level 8
in main 7
Sqrt of  26  =  5.0990195135927845
The final answer is 5.0990195135927845 ,Thank you
The final answer is 5.10 ,Thank you
Ceil :  6  Floor :  5
Middle  7
Default
Sunday Monday
Saturday
Good Night
{101 Akash 91}
101 Akash 91
{102  92}
102  92
9
8
7
6
5
4
3
2
1
0
*/

/*
// Tutorial Playlist Link:
https://www.youtube.com/playlist?list=PLsyeobzWxl7pJ9Gy1iHRKjUTE5xPhJ18b
*/
