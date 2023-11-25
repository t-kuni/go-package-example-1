package main

import "github.com/t-kuni/go-package-example-2/lib"

func main() {
	println("Hello, World!")
	i := lib.GetInfo()
	println(i.Name)
}
