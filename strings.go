package main

import (
	"fmt"
	"strings"
)

func main() {
	// This is an example of a string
	fmt.Println("Hello")
	fmt.Println("-")
	var ourString string
	ourString = "This is a Test"
	fmt.Println(ourString)
	fmt.Println("-")
	//This is an example of a Rune
	var ourRune rune
	ourRune = 'g'
	fmt.Printf("Type: %T, \v Value: %v\n", ourRune, ourRune)
	fmt.Println("-")
	// Example of needing special character
	fmt.Println("Don't touch That, that's my food")
	fmt.Println("She said, \"Thats my food too\"")
	fmt.Println("-")
	// Example of string literal
	fmt.Println(`Hello my nmae is Gerard, 
	and I am the Liverpool legend`)

	//length function 
	howlong := len("bre ak")
	print(howlong)
	
	fmt.Println(len("hello world"))   //length=6 counts the space

	//String formatting interpullation
	name := "Samboo"
	age := 45
	fmt.Printf("Hello my name is %s and I am %d years old", name, age)
	fmt.Println("-")
	var names string 
	names = "Abebe"
	team := 'L'
	ages := 36
	fmt.Printf("Hello my name is %s and I am %d years old and I represent team %c\n", names, ages, team)

	//String methods
	// upper and lower letters
	fmt.Println(strings.ToUpper("liverpool fc"))
	fmt.Println(strings.ToLower("CONVERT ALL TO LOWERCASE"))
	// change first letters from lower to upper
	fmt.Println(strings.Title("her royal highness"))
	fmt.Println(strings.Title("loud noises"))	
	fmt.Println(strings.Title("bread"))
	// split
	s := "why are we here"
	sSplit := strings.Split(s, "f")
	fmt.Println(sSplit)

}
