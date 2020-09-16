package main
import "fmt"
var gvariable int=200 //Define cont gobal
func main()  {
  lvariable:=30
  fmt.Println("gobal =",gvariable)
  fmt.Println("local =",lvariable)
  anotherFunction()
}
func anotherFunction()  {
  lvariable:=60
  fmt.Println("gobal =",gvariable)
  fmt.Println("local =",lvariable)

}
