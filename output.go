package main

import "fmt"

func main() {
    // variables
    name := "Maria"
    fmt.Println("Hello " + name)

    // *** Data Types *** //
    // String //
    var ourString string
    ourString = "ssdfjsdf 46789"
    fmt.Printf("Type: %T\n", ourString)
    
    // integers //
    var ourInt int
    ourInt = 46789
    fmt.Printf("Type: %T\n", ourInt)
    
    // boolean //
    var ourBool bool
    ourBool = true
    fmt.Printf("Type: %T\n", ourBool)
    
    // Float //
    var ourFloat float32
    ourFloat = 3.45677
    fmt.Printf("Type: %T\n", ourFloat)
    
    // *** collection *** //
    // Slice //
    var ourSlice = []string{"hello", "goodbye", "gerrard"}
    fmt.Println(ourSlice)
    var ourIntSlice = []int{4, 788, 34454}
    fmt.Println(ourIntSlice)

    // rune byte
    var ourrune rune
    ourrune = 'r'
    fmt.Printf("Type: %T\n", ourrune)

}