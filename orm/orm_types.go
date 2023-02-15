package orm

import "gorm.io/gorm"

type TORM struct {
	ORM *gorm.DB
}

var (
	DB TORM
)
