package main
import "fmt"

func main()  {
  add:=func (x,y int) int{
    return x+y
  }
  fmt.Println(add(4,5)) //add(x,y)
  x:=0
  increment:=func()int{
    x++
    return x
  }
  fmt.Println(increment())
}
