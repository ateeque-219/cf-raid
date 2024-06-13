package main
 import "fmt"




 func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, World!")
}
func main(){
	fmt.Println("abc")
 

	http.HandleFunc("/", helloWorldHandler)


}