package orm

import (
	"github.com/lantaris/rest-go-sdk/fmtlog"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ***********************************************************************
func (SELF *TORM) Init(Host string, Database string, Username string, Password string) error {
	var (
		err error = nil
	)

	// Connect to database
	err = SELF.Connect(Host, Database, Username, Password)
	if err != nil {
		return err
	}

	return err
}

// ***********************************************************************
func (SELF *TORM) Connect(Host string, Database string, Username string, Password string) error {
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
	SELF.ORM, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmtlog.Errorln("Error connect to database")
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
