package main
import "fmt"
  type books struct{
    title string
    price float64
    author string
  }
func main()  {
  Book:=books{title:"KPS",author:"Zoybad",price:50}
  fmt.Println(Book)
}
