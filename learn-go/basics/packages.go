package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	fmt.Println("My favorite number is", rand.Intn(10))

	Another()
}

func Another() {
	fmt.Println("another number is", rand.Intn(100))
}
