package utils

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	if db == nil {
		fmt.Println("Creating new host db connection")
		dbStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASSWORD"),
			os.Getenv("POSTGRES_HOST"),
			os.Getenv("POSTGRES_PORT"),
			os.Getenv("POSTGRES_DB"),
		)

		gormDB, err := setupDBConnection(dbStr)
		if err != nil {
			fmt.Println("Error setting up db connection: ", err)
			panic(err)
		}

		db = gormDB
	}

	return db
}

func setupDBConnection(connectionString string) (*gorm.DB, error) {
	sqlDB, err := sql.Open("pgx", connectionString)
	if err != nil {
		fmt.Println("Error connecting to db: ", err)
		return nil, err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: sqlDB}), &gorm.Config{})
	if err != nil {
		fmt.Println("Error initializing db: ", err)
		return nil, err
	}

	// newRedisClient := GetRedisClient()
	// cache, err := cache.NewGorm2Cache(&config.CacheConfig{
	// 	CacheLevel:           config.CacheLevelAll,
	// 	CacheStorage:         config.CacheStorageRedis,
	// 	RedisConfig:          cache.NewRedisConfigWithClient(newRedisClient),
	// 	InvalidateWhenUpdate: true,   // when you create/update/delete objects, invalidate cache
	// 	CacheTTL:             100000, // 100s
	// 	CacheMaxItemCnt:      20,     // if length of objects retrieved one single time exceeds this number, then don't cache
	// })

	// if err != nil {
	// 	fmt.Println("Error creating caching layer: ", err)
	// 	return nil, err
	// }

	// err = gormDB.Use(cache)
	// if err != nil {
	// 	fmt.Println("Error using caching layer: ", err)
	// 	return nil, err
	// }

	return gormDB, nil
}
