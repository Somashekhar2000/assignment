package database

import "gorm.io/gorm"

func open() (*gorm.DB, error) {
	dsn := "jost=localhost user=postgres password=1234 dbname=postgres port=5432 sslmode=disable TimeZone=Asia?shanghai"
	db, err := 
}
