package orm

import (
	"rest-go-sdk/logger"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ***********************************************************************
func Init(Host string, Database string, Username string, Password string) error {
	var (
		err error = nil
	)

	// Connect to database
	err = Connect(Host, Database, Username, Password)
	if err != nil {
		return err
	}

	// Migration if needded
	err = Migrate()
	if err != nil {
		return err
	}

	return err
}

// ***********************************************************************
func Connect(Host string, Database string, Username string, Password string) error {
	var (
		err error = nil
		db_host, db_username,
		db_password, db_database, dsn string
	)

	// Get database configuration
	db_host = Host
	db_database = Database
	db_username = Username
	db_password = Password

	dsn = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", db_username, db_password, db_host, db_database)
	DB.ORM, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Errorln("Error connect to database")
	}
	return err
}

// ***********************************************************************
func Migrate(dst ...interface{}) error {
	var (
		err error = nil
	)
	err = DB.ORM.AutoMigrate(dst...)
	if err != nil {
		logger.Errorln("Error database migrate", err.Error())

	}

	return err

}
