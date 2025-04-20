package main
import(
"database/sql"
"fmt"
_ "github.com/lib/pq"
)
func connect(input string , output string) error{
db, err :=sql.Open("postgres","host=localhost port=5432 user=postgres dbname=requests sslmode=disable")
if err != nil{
return fmt.Errorf("error in accesing the database : %v",err)
}
defer db.Close()
if err :=db.Ping();err!=nil{fmt.Errorf("error while connecting to database" , err)}
_, err = db.Exec("INSERT INTO reversed_strings (input , output) VALUES ($1 ,$2)",input, output)
if err != nil{
return fmt.Errorf("error in inserting the values : %v" ,err)
}
return nil
}

