package datastores

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"properly.jp/properlyweb-backend/adapters/gateways"
)

type SQLHandler struct {
	db *gorm.DB
}

func InitSQLHandler() gateways.SQLHandler {
	user := "root"               // os.Getenv("DB_USER")
	pass := "password"           // os.Getenv("DB_PASSWORD")
	containerName := "localhost" // os.Getenv("DB_CONTAINER_NAME")
	port := "3306"               // os.Getenv("DB_PORT")
	dbName := "development"      // os.Getenv("DB_NAME")

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, containerName, port, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
	sqlHandler := new(SQLHandler)
	sqlHandler.db = db
	return sqlHandler
}

func (handler *SQLHandler) First(out interface{}, where ...interface{}) *gorm.DB {
	return handler.db.First(out, where...)
}

func (handler *SQLHandler) Take(out interface{}, where ...interface{}) *gorm.DB {
	return handler.db.Take(out, where...)
}

func (handler *SQLHandler) Last(out interface{}, where ...interface{}) *gorm.DB {
	return handler.db.Last(out, where...)
}

func (handler *SQLHandler) Find(out interface{}, where ...interface{}) *gorm.DB {
	return handler.db.Find(out, where...)
}
