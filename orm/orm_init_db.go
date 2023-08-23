package orm

import (
	"fmt"
	"github.com/lantaris/rest-go-sdk/fmtlog"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	MYSQL_DEF_PORT    = "3306"
	POSTGRES_DEF_PORT = "5432"
)

// ***********************************************************************
func (SELF *TORM) Init(DBtype string, Host string, Port string, Database string, Username string, Password string) error {
	var (
		err error = nil
	)

	// Connect to database
	err = SELF.Connect(DBtype, Host, Port, Database, Username, Password)
	if err != nil {
		return err
	}

	return err
}

// ***********************************************************************
func (SELF *TORM) Connect(DBtype string, Host string, Port string, Database string, Username string, Password string) error {
	var (
		err error = nil
		db_host, db_username,
		db_password, db_database,
		dsn string
		DBPort string
	)

	// Get database configuration
	db_host = Host
	db_database = Database
	db_username = Username
	db_password = Password

	if DBtype == "mysql" {
		fmtlog.Infoln("Initialization MySQL database")
		if Port == "" {
			DBPort = MYSQL_DEF_PORT
		} else {
			DBPort = Port
		}
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			db_username, db_password, db_host, DBPort, db_database)
		SELF.ORM, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			fmtlog.Errorln("Error connect to database")
			return err
		}
	} else if DBtype == "postgres" {
		fmtlog.Infoln("Initialization PostgreSQL database")
		if Port == "" {
			DBPort = POSTGRES_DEF_PORT
		} else {
			DBPort = Port
		}
		dsn = fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
			db_username, db_password, db_host, DBPort, db_database)
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
