package main
import "fmt"

func main()  {
  x:=make(map[string] string)
  x["LA"]="Laos"
  x["TH"]="Thailand"
  x["JP"]="Japan"
  fmt.Println(x["LA"])

another()
}
func another ()  {

    var n int
    fmt.Scanf("%d",&n)
    y:=make(map[int] string)
    y[1]="Laos"
    y[2]="Thailand"
    y[3]="Japan"

    fmt.Println(y[n])

  }
