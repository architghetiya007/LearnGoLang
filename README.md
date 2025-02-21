## Declare the Variable

1. **Using the var keyword (Explicit type)** /n
var name string = "Archit"
var age int = 30

2. **Using var without an initial value (Zero value)**
<br>
var city string  //Zero value: ""
var score int    // Zero value: 0
var isActive bool // Zero value: false

3. Type inference (without specifying type)
var country = "India" // Compiler infers type as string
var number = 100      // Compiler infers type as int

4. Short declaration using := (Inside functions only)
name := "Archit1" // Equivalent to: var name string = "Archit"
age := 30        // Equivalent to: var age int = 30

5. Multiple variable declaration
var x, y, z int = 1, 2, 3
// name, age := "Archit", 30

6. Grouped declaration
var (
    firstName string = "Archit"
    lastName  string = "Patel"
    isAdmin   bool   = true
)

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

