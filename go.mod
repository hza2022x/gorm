module gorm.io/gorm

go 1.18

require (
	github.com/jinzhu/inflection v1.0.0
	github.com/jinzhu/now v1.1.5
)

require github.com/go-sql-driver/mysql v1.7.0

//dbpool => ../dbpool
replace dbpool => ../dbpool
