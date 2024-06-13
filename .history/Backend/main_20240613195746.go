package main

import (
    "fmt"
    "log"
    "net/http"
)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, Aditya World!")
}

func main() {
    http.HandleFunc("/", helloWorldHandler)

    fmt.Println("Starting server at port 8000")
    serverAddr := ":8000"
    err := http.ListenAndServe(serverAddr, nil)
    if err != nil {
        log.Fatalf("server error: %s", err)
    }
    fmt.Printf("Server is listening on %s\n", serverAddr)
}
