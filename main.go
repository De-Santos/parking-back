package main

import (
	"fmt"
	"parking-back/initializers"
)

func init() {
	
	initializers.LoadEnvVariables()
}

func main() {
	fmt.Println("Hello")
}
