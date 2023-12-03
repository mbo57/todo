package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
)

func NewDB() *gorm.DB {
	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
	}

	writer := fmt.Sprintf("%s:%s@%s/%s?charset=utf8&parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("WRITER_DB_HOST"),
		os.Getenv("DB_DB"),
	)
	reader := fmt.Sprintf("%s:%s@%s/%s?charset=utf8&parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("READER_DB_HOST"),
		os.Getenv("DB_DB"),
	)
	db, err := gorm.Open(mysql.Open(writer), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				ParameterizedQueries: true,
			},
		),
	})
	if err != nil {
		log.Fatalln(err)
	}
	// db.Logger = db.Logger.LogMode(logger.Info)
	err = db.Use(dbresolver.Register(dbresolver.Config{
		Replicas:          []gorm.Dialector{mysql.Open(reader)},
		Policy:            dbresolver.RandomPolicy{},
		TraceResolverMode: true,
	}))
	if err != nil {
		log.Fatalln(err)
	}
	return db
}

func CloseDB(db *gorm.DB) {
	sqlDB, _ := db.DB()
	if err := sqlDB.Close(); err != nil {
		log.Fatalln(err)
	}
	fmt.Println("データベースとの接続を切断しました")
}
