//reverse_handler.go
package main
import(
"fmt"
"net/http"
"log"
)
func response_sender_function(w http.ResponseWriter , r *http.Request){

if r.Method != "POST" {http.Error(w , "invalid request method , must use POST ", http.StatusMethodNotAllowed)
return }

parseerr :=r.ParseForm()
 if parseerr !=nil {
http.Error(w , "cant parse the string " , http.StatusBadRequest)
return
}
input :=r.FormValue("text")
if input ==""{http.Error(w,"missing the string to reverse",http.StatusBadRequest)
return}

output :=reverse(input)
fmt.Fprintln(w , output)

connectionerr := connectandcheckdbconnection()
if connectionerr != nil {
log.Println(connectionerr)
http.Error(w , connectionerr.Error() , http.StatusInternalServerError)
return
}
defer db.Close()
inserterr := inserttodb(input , output)
if inserterr != nil{
log.Println("database connection has errors :" , inserterr)
http.Error(w, inserterr.Error(),http.StatusInternalServerError)
return
}
}
func reverse(str string ) string{
var ans = []byte{}
for j := len(str)-1 ; j>=0 ; j--{
ans = append(ans , str[j])
}
return string(ans)
}
