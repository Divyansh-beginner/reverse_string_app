package main 
import ("fmt"
"net/http"
"log"
)
func reverse (str string ) string{
var ans = []byte{}
for j := len(str)-1 ; j>=0 ; j--{
ans = append(ans , str[j])
}
return string(ans)
}
func response_sender_function(w http.ResponseWriter , r *http.Request){
fmt.Fprintf(w , reverse(r.FormValue("text")))
}
func main(){
http.HandleFunc("/reverse" , response_sender_function)
log.Println("this server is running at http://localhost:8080")
log.Fatal(http.ListenAndServe(":8080" , nil))
}

