//db.go
package main
import(
"database/sql"
"fmt"
"time"
_ "github.com/lib/pq"
)
type reversedentry struct{
Input string
Output string
Created_at time.Time
}

var db *sql.DB

func connectandcheckdbconnection() error{
    var err error
    db, err =sql.Open("postgres","host=localhost port=5432 user=postgres dbname=requests sslmode=disable")
    if err != nil{
    return fmt.Errorf("error in accesing the database : %v",err)
    }
    if pingerr :=db.Ping();pingerr!=nil{
    return fmt.Errorf("error while connecting to database" , pingerr)
    }
    return nil
}

func inserttodb(input string , output string) error{
_, execerr := db.Exec("INSERT INTO reversed_strings (input , output) VALUES ($1 ,$2)",input, output)
if execerr != nil{
return fmt.Errorf("error in inserting the values : %v" ,execerr)
}
return nil
}

func checkindb(input string)(bool , []reversedentry , error){
query := "select input, output, created_at from reversed_strings where input = $1;"
rows, queryerr :=db.Query(query,input)
if queryerr != nil {
 return false , nil , fmt.Errorf("check in db failed , error : %w",queryerr)
}
defer rows.Close()

var results []reversedentry
for rows.Next() {
 var entry reversedentry
 rowerr := rows.Scan(&entry.Input, &entry.Output, &entry.Created_at)
 if rowerr != nil {
  return false , nil , fmt.Errorf("rows exists but error on scanning the rows, error : %w",rowerr)
 }
 results = append(results, entry)
}
iterationerr := rows.Err()
if iterationerr != nil {
 return false , nil , fmt.Errorf("rows exists but some internal error : %w",iterationerr)
}

return true , results , nil
}


