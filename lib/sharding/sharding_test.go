package sharding

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"testing"
)

var (
	db       *gorm.DB
	dbConfig postgres.Config
	dbURL    = "postgres://localhost:5432/sharding-test?sslmode=disable"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dbURL = os.Getenv("DB_URL")
	dbConfig = postgres.Config{DSN: dbURL, PreferSimpleProtocol: true}
	db, _ = gorm.Open(postgres.New(dbConfig), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	db, err := gorm.Open(mysql.Open("db1_dsn"), &gorm.Config{})

	db.Use(Register(Config{
		NumberOfShards: 2,
		ShardingKey:    "id",
		ShardingAlgorithm: func(columnValue any) (suffix string, err error) {
			return fmt.Sprintf("%02d", columnValue.(int64)%2), nil
		},
	}, "sys_user"))
}

func TestMigrate(t *testing.T) {

}
