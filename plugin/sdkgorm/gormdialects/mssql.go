package gormdialects

import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

// Ex: sqlserver://username:password@localhost:1433?database=db_name
func MSSqlDB(uri string) (db *gorm.DB, err error) {
	return gorm.Open(sqlserver.Open(uri))
}
