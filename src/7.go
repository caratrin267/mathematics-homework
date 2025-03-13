package main

import "math/rand"

func main() {
    // Generate a random number between 1 and 10
    num := rand.Intn(10) + 1

    // Print the result
    println("Random number:", num)
}
