package main
import (
    "fmt"
    "log"
    "net/http"
)



 func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, Aditya World!")
}
func main(){
	http.HandleFunc("/", helloWorldHandler)

	 res, err := http.ListenAndServe(":8000")
	if err != nil {
		log.Fatalf("server error")
	}
	fmt.Println("Starting server at port 8000")

}