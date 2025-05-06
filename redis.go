//redis.go
package main
import(
"fmt"
"time"
"github.com/redis/go-redis/v9"
"context"
)
var redisclient *redis.Client

func setupredisclient(){

options := &redis.Options{
Addr: "localhost:6379",
Password: "",
DB: 0,
}

redisclient = redis.NewClient(options)

}


func checkredisconnection() error{

ctx , cancel := context.WithTimeout(context.Background(), 2*time.Second)

defer cancel()

pong, err := redisclient.Ping(ctx).Result()
if err != nil {
return fmt.Errorf("redis is not giving ping connection : %w",err)
}
fmt.Printf("redis connection successful : %v",pong)
return nil
}
