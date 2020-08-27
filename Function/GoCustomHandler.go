package main

import (
    "fmt"
    "net/http"
    "log"
    "os"
    "strings"
)

func httpTriggerHandler(w http.ResponseWriter, r *http.Request) {
	clientIP := r.Header.Get("X-Forwarded-For")
    w.Write([]byte(strings.Split(clientIP,":")[0]))
}



func main() {
	httpInvokerPort, exists := os.LookupEnv("FUNCTIONS_HTTPWORKER_PORT")
	if exists {
		fmt.Println("FUNCTIONS_HTTPWORKER_PORT: " + httpInvokerPort)
	}
    mux := http.NewServeMux()
    mux.HandleFunc("/httptrigger", httpTriggerHandler)
	log.Println("Go server Listening...on httpInvokerPort:", httpInvokerPort)
	log.Fatal(http.ListenAndServe(":"+httpInvokerPort, mux))
}
