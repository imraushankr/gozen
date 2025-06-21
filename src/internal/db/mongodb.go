package db

import (
  "context"
  "fmt"
  "time"
  "github.com/imraushankr/gozen/src/internal/config"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
)

func initMongoDB(cfg config.DatabaseConfig) error {
  ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
  defer cancel()
  
  // MongoDB connection string
  uri := fmt.Sprintf("mongodb://%s:%d", cfg.Host, cfg.Port)
  if cfg.Username != "" && cfg.Password != "" {
    uri = fmt.Sprintf("mongodb://%s:%s@%s:%d",
      cfg.Username, cfg.Password,
      cfg.Host, cfg.Port)
  }
  
  clientOptions := options.Client().ApplyURI(uri)
  clientOptions.SetMaxPoolSize(uint64(cfg.MaxOpenConns))
  clientOptions.SetMaxConnIdleTime(cfg.ConnMaxLifetime)
  
  client, err := mongo.Connect(ctx, clientOptions)
  if err != nil {
    return err
  }
  
  // Ping to verify connection
  if err = client.Ping(ctx, nil); err != nil {
    return err
  }
  
  MongoInstance = client.Database(cfg.Database)
  fmt.Println("✅ Connected to MongoDB")
  return nil
}