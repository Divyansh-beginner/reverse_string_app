//reverse_string.go
package main
import (
"net/http"
"log"
)

func main(){
http.HandleFunc("/reverse" , response_sender_function)
http.HandleFunc("/check" , check_function)
log.Println("this server is running at http://localhost:8080")
log.Fatal(http.ListenAndServe(":8080" , nil))
}

