package main
import "math/rand"
func main() {
    const lowercaseLetters = "abcdefghijklmnopqrstuvwxyz"
    var randomCode string
    for i := 0; i < 10; i++ {
        randomCode += string(lowercaseLetters[rand.Intn(26)])
    }
    fmt.Println(randomCode)
}
