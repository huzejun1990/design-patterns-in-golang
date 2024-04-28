// @Author huzejun 2024/4/28 19:07:00
package main

import (
	"design-patterns-in-golang/Decorator"
	"log"
	"os"
)

func main() {
	/*	fmt.Println("main")
		fmt.Println(Decorator.Pi(5000))
		fmt.Println(Decorator.Pi(10000))
		fmt.Println(Decorator.Pi(50000))*/

	f := Decorator.WrapLogger(Decorator.Pi, log.New(os.Stdout, "test", 1))
	f(100000)
}
