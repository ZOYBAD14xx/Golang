package main
import (
  "fmt"
  "net/http"
)
func main()  {
  http.HandleFunc("/",local)

  http.ListenAndServe(":8080",nil)
}
func local (w http.ResponseWriter, r *http.Request)  {
fmt.Fprintln(w,"Khonepaserth")
}
