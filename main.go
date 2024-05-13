package main

import "fmt"

func main() {
	var printValue  string   = "hello world"
printMe(printValue)
fmt.Println(intDivision(10,3))
}

func printMe(printValue string){
	fmt.Println(printValue)
}

func intDivision(n int , d int ) (int, int){
return n/d, n%d
}