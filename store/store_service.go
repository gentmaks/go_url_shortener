package store 

import "context" 
import "fmt" 
import "github.com/go-redis/redis"
import "time" 

// Struct wrapper around raw Redis client 
type StorageService struct {
  redisClient *redis.Client 
}

var (
  storeService = &StorageService{}
  ctx = context.Background()
)

const CacheDuration = 6 * time.Hour 

func InitializeStore() *StorageService {
  redisClient := redis.NewClient(&redis.Options{
    Addr: "localhost:6379",
    Password: "", 
    DB: 0
  })
  pont, err := redisClient.Ping(ctx).Result() 
  if err != nil {
    panic(fmt.Sprintf("Erro init Redis: %v", err))
  }
  fmt.Printf("\nRedis started successfully: pong message = {%s}", pong)
  storeService.redisClient = redisClient 
  return storeService
}

func SaveUrlMapping (shortUrl string, originalUrl string, userId string) {
  err := storeService.redisClient.SET(ctx, shortUrl, originalUrl, CacheDuration).Err()
  if err != nil {
    panic(fmt.Sprintf("Failed saving key url | Error: %v - shortUrl: %s - originalUrl: %s\n", err, shortUrl, originalUrl)) 
  }
}

func RetreiveInitialUrl(shortUrl string) {
  result, err := storeService.redisClient.GET(ctx, shortUrl).Result()
  if err != nil {
    panic(fmt.Sprintf("Failed RetreiveInitialUrl | Error: %v, shortUrl: %s", err, shortUrl))
  }
  return result
}
