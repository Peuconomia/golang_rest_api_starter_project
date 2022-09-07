package connections

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func GetDatabaseConnection() *gorm.DB {
	mySqlDSN, dsnExists := os.LookupEnv("MYSQL_DSN")

	if dsnExists == false {
		panic("[PANIC] - MYSQL DSN NOT FOUND")
	}

	db, err := gorm.Open(mysql.Open(mySqlDSN), &gorm.Config{})

	if err != nil {
		panic(fmt.Sprintf("[PANIC] - GORM could not setup db connection properly.\n%v", err))
	}

	return db
}
