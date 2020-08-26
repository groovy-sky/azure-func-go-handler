package main

import (
    "fmt"
    "http"
    "time"
)

func httpTriggerHandler(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	fmt.Println(t.Month())
	fmt.Println(t.Day())
	fmt.Println(t.Year())
	ua := r.Header.Get("User-Agent")
	invocationid := r.Header.Get("X-Azure-Functions-InvocationId")
	fmt.Printf("invocationid is: %s \n", invocationid)
    w.Write([]byte(ua))
}



func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/HttpTrigger", httpTriggerHandler)
}
