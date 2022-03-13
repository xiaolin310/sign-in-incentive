package main

import (
    "context"
    "encoding/json"
    "fmt"
    "github.com/go-redis/redis/v8"
)


func main() {
    rdb := redis.NewClient(&redis.Options{
        Addr: "127.0.0.1:6379",
    })
    ctx := context.Background()
    data := make(map[string]interface{})
    data["20220310"], _ = json.Marshal([]float64{11.01, 12.01})
    err := rdb.HSet(ctx, "88", data).Err()
    if err != nil {
        fmt.Printf("%v", err)
        return
    }
    value, err := rdb.HGetAll(ctx, "88").Result()
    b, ok := value["20220310"]
    if !ok {
        fmt.Printf("not found")
        return
    }
    var tmp []float64
    err = json.Unmarshal([]byte(b), &tmp)
    if err != nil {
        fmt.Printf("parse error: %v", err)
        return
    }
    fmt.Println(tmp)

}