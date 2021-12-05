
package main 

import (
    "fmt"
    "net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "<h1>Selam Dunya</h1>")
}

func about(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "<h1>Hersey</h1>")
}

func main() {
  http.HandleFunc("/", index)
  http.HandleFunc("/hersey", about)
  
  fmt.Println("Server Basliyor....")
  http.ListenAndServe(":3000", nil)
}