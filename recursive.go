package main
import "fmt"
func main()  {
  var n int
  fmt.Print("Input your number : ")
  fmt.Scanf("%d",&n)
  fmt.Println(fac(n))
}
func fac(num int)  int{
if num == 0{
  return 1
}
return num*fac(num-1)
}
