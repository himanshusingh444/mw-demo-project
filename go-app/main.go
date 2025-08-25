package main
import ("fmt"; "net/http")
func main() { http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, "OK") }); http.ListenAndServe(":8080", nil) }