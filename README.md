# Declare the Variable

1. **Using the var keyword (Explicit type)** 

    var name string = "Archit"
    var age int = 30

2. **Using var without an initial value (Zero value)**

    var city string  //Zero value: ""
    var score int    // Zero value: 0
    var isActive bool // Zero value: false

3. **Type inference (without specifying type)**

    var country = "India" // Compiler infers type as string
    var number = 100      // Compiler infers type as int

4. **Short declaration using := (Inside functions only)**

    name := "Archit1" // Equivalent to: var name string = "Archit"
    age := 30        // Equivalent to: var age int = 30

5. **Multiple variable declaration**

    var x, y, z int = 1, 2, 3
    name, age := "Archit", 30

6. **Grouped declaration**

    var (
        firstName string = "Archit"
        lastName  string = "Patel"
        isAdmin   bool   = true
    )
```
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
```

# Looping concept in the GoLang
1. **Simple for loop.** 
```
package main
import "fmt"
func main() {
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }
}
```
2. **While loop** 
```
package main
import "fmt"
func main() {
    count := 0
    for count < 5 {
        fmt.Println(count)
        count++
    }
}
```

3. **Infinite Loop** 
```
package main
import "fmt"
func main() {
    for {
        fmt.Println("This will loop forever")
        break // To prevent infinite loop
    }
}
```

4. **Using for with range (For Arrays, Slices, Maps, and Strings)**
```
package main

import "fmt"

func main() {
    numbers := []int{10, 20, 30, 40, 50}

    for index, value := range numbers {
        fmt.Println("Index:", index, "Value:", value)
    }
}
```

    OR
```
package main

import "fmt"

func main() {
    numbers := []int{10, 20, 30, 40, 50}
    
    for _, value := range numbers {
        fmt.Println(value)
    }   
}
```
