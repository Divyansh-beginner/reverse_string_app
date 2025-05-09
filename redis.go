//redis.go
package main
import(
"fmt"
"time"
"log"
"github.com/redis/go-redis/v9"
"context"
)
var redisclient *redis.Client

func setupredisclient(){

options := &redis.Options{
Addr: "redis_service3_redis:6379",
Password: "",
DB: 0,
}

redisclient = redis.NewClient(options)

}


func checkredisconnection() error{

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	var lastErr error
	for i := 0; i < 5; i++ {
		_, err := redisclient.Ping(ctx).Result()
		if err == nil {
			fmt.Println("✅ Redis connected!")
			return nil
		}
		lastErr = err
		log.Printf("Waiting for Redis... attempt %d: %v\n", i+1, err)
		time.Sleep(2 * time.Second)
	}

	return fmt.Errorf("redis is not giving ping connection: %w", lastErr)
}
