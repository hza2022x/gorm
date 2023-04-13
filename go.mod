module gorm.io/gorm

go 1.18

require (
	github.com/jinzhu/inflection v1.0.0
	github.com/jinzhu/now v1.1.5
	gorm.io/dbpool v0.0.0-00010101000000-000000000000
)

replace gorm.io/dbpool => ../dbpool
