package main
import "fmt"
func main()  {

// var x[5] int
// x[0]=5
// x[1]=10
// x[2]=15
// x[3]=20
// x[4]=25

x:=[5]float64{5,10,15,20,25}
var a float64
for _ ,n:= range x{
a+=n
}
fmt.Println(a)
}
