//redis_utils.go
package main 

import(
"fmt"
"time"
"context"
"github.com/redis/go-redis/v9"
"encoding/json"
)

func settocache(key string ,entries []reversedentry , expiration time.Duration) error{

ctx , cancel := context.WithTimeout(context.Background(), 2*time.Second) 
defer cancel()
value, jsonerr := json.Marshal(entries)
if jsonerr != nil{
return fmt.Errorf("failed to convert the rows into json for redis , error : %w", jsonerr)
}

redisseterr := redisclient.Set(ctx , key ,value ,expiration).Err()
if redisseterr != nil {
return fmt.Errorf("failed to set key in redis , error : %w", redisseterr)
}
return nil 
}

func getfromcache(key string) ([]reversedentry , error){

ctx , cancel := context.WithTimeout(context.Background() , 2*time.Second)
defer cancel()
value , redisgeterr := redisclient.Get(ctx , key).Result()
if redisgeterr == redis.Nil{
return nil , nil
}
if redisgeterr != nil {
return nil ,fmt.Errorf("errror on getting the key , error : %w",redisgeterr)
}
var entries []reversedentry
unmarshalerr := json.Unmarshal([]byte(value), &entries)
if unmarshalerr != nil{
return nil , fmt.Errorf("error in converting back from json , error : %w",unmarshalerr)
}
return entries , nil
}
