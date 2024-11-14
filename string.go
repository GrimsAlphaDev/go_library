package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Hello world", "world"))
	fmt.Println(strings.Split("a,b,c,d,e", ","))
	fmt.Println(strings.ToLower("HELLO WORLD"))
	fmt.Println(strings.ToUpper("hello world"))
	fmt.Println(strings.Trim("   Hello world   ", " "))
	fmt.Println(strings.ReplaceAll("Hello world", "world", "Go"))
}
