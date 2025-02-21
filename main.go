package main

import (
    "fmt"
)

func main() {

//1. Using the var keyword (Explicit type)
var name string = "Archit"
var age int = 30

//2. Using var without an initial value (Zero value)
var city string  // Zero value: ""
var score int    // Zero value: 0
var isActive bool // Zero value: false

//  Type inference (without specifying type)
var country = "India" // Compiler infers type as string
var number = 100      // Compiler infers type as int

// // 4. Short declaration using := (Inside functions only)
name := "Archit1" // Equivalent to: var name string = "Archit"
age := 30        // Equivalent to: var age int = 30

// 5. Multiple variable declaration
var x, y, z int = 1, 2, 3
// name, age := "Archit", 30

// 6. Grouped declaration
var (
    firstName string = "Archit"
    lastName  string = "Patel"
    isAdmin   bool   = true
)

    // Printing values
    fmt.Println("Name:", name)
    fmt.Println("Age:", age)
    fmt.Println("Is Admin:", isAdmin)
    fmt.Println("City:", city)
    fmt.Println("Score:", score)
    fmt.Println("Country:", country)
    fmt.Println("Is Active:", isActive)
    fmt.Println("Number:", number)
    fmt.Println("x, y, z:", x,y,z)
    fmt.Println("FirstName:", firstName)
    fmt.Println("LastName:", lastName)

}
// package main

// import (
// 	"fmt"
// 	"runtime"
// 	"time"
// )

// func printNumbers() {
// 	for i := 1; i <= 5; i++ {
// 		fmt.Println(i)
// 		time.Sleep(time.Millisecond * 500)
// 	}
// }

// func printLetters() {
// 	for ch := 'A'; ch <= 'E'; ch++ {
// 		fmt.Println(string(ch))
// 		time.Sleep(time.Millisecond * 500)
// 	}
// }

// func main() {
// 	runtime.GOMAXPROCS(runtime.NumCPU()) // Use all available CPU cores

// 	go printNumbers() // Run concurrently
// 	go printLetters() // Run concurrently

// 	time.Sleep(time.Second * 3) // Wait for goroutines to finish
// 	fmt.Println("Main function finished")
// }
