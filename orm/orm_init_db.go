package orm

import (
	"github.com/lantaris/rest-go-sdk/fmtlog"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ***********************************************************************
func (SELF *TORM) Init(DBtype string, Host string, Database string, Username string, Password string) error {
	var (
		err error = nil
	)

	// Connect to database
	err = SELF.Connect(DBtype, Host, Database, Username, Password)
	if err != nil {
		return err
	}

	return err
}

// ***********************************************************************
func (SELF *TORM) Connect(DBtype string, Host string, Database string, Username string, Password string) error {
	var (
		err error = nil
		db_host, db_username,
		db_password, db_database,
		dsn string
	)

	// Get database configuration
	db_host = Host
	db_database = Database
	db_username = Username
	db_password = Password

	if DBtype == "mysql" {
		fmtlog.Infoln("Initialization MySQL database")
		dsn = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", db_username, db_password, db_host, db_database)
		SELF.ORM, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			fmtlog.Errorln("Error connect to database")
			return err
		}
	} else if DBtype == "postgres" {
		fmtlog.Infoln("Initialization PostgreSQL database")
		dsn = fmt.Sprintf("user=%s password=%s host=%s dbname=%s sslmode=disable", db_username, db_password, db_host, db_database)
		SELF.ORM, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			fmtlog.Errorln("Error connect to database")
			return err
		}
	} else {
		fmtlog.Errorln("Error database type")
		return err
	}

	return err
}

// ***********************************************************************
func (SELF *TORM) Migrate(dst ...interface{}) error {
	var (
		err error = nil
	)
	err = SELF.ORM.AutoMigrate(dst...)
	if err != nil {
		fmtlog.Errorln("Error database migrate", err.Error())

	}

	return err

}
