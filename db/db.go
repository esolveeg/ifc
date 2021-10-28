package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

var (
	DBConn *gorm.DB
)

func InitDatabase() error {
	var err error
	// conStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_HOST"), config.Config("DB_PORT"), config.Config("DB_NAME"))
	// DBConn, err := gorm.Open(oracle.Open("system/oracle@127.0.0.1:1521/XE"), &gorm.Config{})
	// DBConn, err = gorm.Open(oracle.Open("%s/%s@%s:%s/%s"), config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_HOST"), config.Config("DB_PORT"), config.Config("DB_NAME"))
	// DBConn, err = gorm.Open("mssql", "sqlserver://mcs:123@41.38.87.59:1433?database=stock_main")
	// DBConn, err = gorm.Open("mysql", conStr)
	if err != nil {
		fmt.Println("Failed to connect to database")
		panic(err)
	}
	fmt.Println("Connection Opened to External Database")
	return nil

}
