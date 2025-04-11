package main 
import "fmt"
func main(){
fmt.Print("enter your string : ")
var str string 
fmt.Scanln(&str)
fmt.Print("\n")
reverse(str)
fmt.Println("thank you for using our app")


}
func reverse (str string ){
var ans = []byte{}
for j := len(str)-1 ; j>=0 ; j--{
ans = append(ans , str[j])
}
fmt.Println(string(ans))

}
