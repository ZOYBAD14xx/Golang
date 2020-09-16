package main
import "fmt"
func main()  {
fmt.Println("Input number")
var number int
  fmt.Scanf("%d",&number)
  switch number {
  case 0:fmt.Println("Zero")
  case 1:fmt.Println("One")
  case 2:fmt.Println("Two")
default:fmt.Println("Unknow")
  }
}
