package main

import (
    "fmt"
    "net/http"
    "time"
    "log"
    "os"
)

func httpTriggerHandler(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	fmt.Println(t.Month())
	fmt.Println(t.Day())
	fmt.Println(t.Year())
	ua := r.Header.Get("User-Agent")
	invocationid := r.Header.Get("X-Azure-Functions-InvocationId")
	fmt.Printf("Invocationid is: %s \n", invocationid)
	fmt.Printf("User-Agent is: %s \n", ua)
    w.Write([]byte(ua))
}



func main() {
	httpInvokerPort, exists := os.LookupEnv("FUNCTIONS_HTTPWORKER_PORT")
	if exists {
		fmt.Println("FUNCTIONS_HTTPWORKER_PORT: " + httpInvokerPort)
	}
    mux := http.NewServeMux()
    mux.HandleFunc("/httptrigger", httpTriggerHandler)
    fmt.Println(mux)
	log.Println("Go server Listening...on httpInvokerPort:", httpInvokerPort)
	log.Fatal(http.ListenAndServe(":"+httpInvokerPort, mux))
}
