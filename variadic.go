package main
import "fmt"
func main()  {
  sum(1,5,6,2,6)
}
func sum(args...int)  {
  var total int
  for _ ,n:=range args{
    total+=n
  }
  fmt.Println(total)
}
