package main

import "fmt"
import "net/http"

func index(w http.ResponseWriter, r *http.Request){
  for i := 1; i <= 100; i++ {
  fmt.Fprintln(w, "Apa kabar temen!", i)
}
}

func main(){

  http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
  fmt.Fprintln(w, " Hello Motto")
})
  http.HandleFunc("/index", index)
  fmt.Println("Memulai web server pada localhost:8080")
  http.ListenAndServe(":8080", nil)
}
