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

if r.Method != "POST" {http.Error(w , "invalid request method , must use POST ", http.StatusMethodNotAllowed)
return }

if err :=r.ParseForm() ; err !=nil{ http.Error(w , "cant parse the string " , http.StatusBadRequest)
return}
input :=r.FormValue("text")
if input ==""{http.Error(w,"missing the string to reverse",http.StatusBadRequest)
return}

output :=reverse(input)
err := connect(input , output)
fmt.Fprintf(w , output)
if err != nil{
log.Println("database connection has errors : %v" , err)
}
}
func main(){
http.HandleFunc("/reverse" , response_sender_function)
log.Println("this server is running at http://localhost:8080")
log.Fatal(http.ListenAndServe(":8080" , nil))
}

