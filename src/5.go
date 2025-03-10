import "math/rand"
func main() {
   var a int
   var b int
   for i := 0; i < len(a); i++ {
      a = rand.Intn(50) + 1
   }
   for j := 0; j < len(b); j++ {
      b = rand.Intn(50) + 1
   }
   fmt.Println(a, b)
}
