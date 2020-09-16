package main
import "fmt"
func main()  {
 slice1:=[]int{1,2,3,4}
 slice2:=append(slice1,5,6) //append = karn tr thaiy
 fmt.Println(slice2)
 fmt.Println()
 another()
}
func another()  {
 s:=[]int{1,2,3,4}
 s2:=make([]int, 3) //make ( type, position )
 copy(s2,s) //copy (copy 1 ma wai 2)
 fmt.Println(s2)
}
