package main
import "fmt"
func main()  {
fmt.Println("input number")
var a float64
fmt.Scanf("%f",&a)
b:= a>2
if b{
  fmt.Println("More than")
}else{
  fmt.Println("Less")
}
}
