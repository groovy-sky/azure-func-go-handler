package main

import (
    "fmt"
    "net/http"
    "log"
    "os"
    "strings"
)

func httpTriggerHandler(w http.ResponseWriter, r *http.Request) {
    if r.Header.Get("X-Forwarded-For") != "" {
        w.Write([]byte(strings.Split(r.Header.Get("X-Forwarded-For"),":")[0]))
    } else if r.Header.Get("Host") != "" {
        w.Write([]byte(strings.Split(r.Header.Get("Host"),":")[0]))
    } else if r.RemoteAddr != "" {
        w.Write([]byte(strings.Split(r.RemoteAddr,":")[0]))
    }
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
