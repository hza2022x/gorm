package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/mysql"
	"strconv"
)

// User 结构体:用户信息
type User struct {
	gorm.Model
	Username string `gorm:"not null;unique"`
	Password string `gorm:"not null"`
	Age      int    `gorm:"null"`
}

func main() {
	query()

	//updateAge()
}

func updateAge() {
	db := getDb()
	db.Exec("Update user set age=? where id=?", 10, 1)
}

//getDb 获取DB对象
func getDb() *gorm.DB {
	var err error
	var db *gorm.DB
	dsn := "root:root@tcp(localhost:13306)/test?parseTime=True&loc=Local"
	// set timeout
	dsn += "&timeout=10s&readTimeout=30s&writeTimeout=30s&parseTime=true"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func queryOneByRaw() User {
	var result User
	db := getDb()
	db.Raw("SELECT id, name, age FROM user WHERE age = ?", 10).Scan(&result)
	fmt.Printf("result:%s", result.Username)
	return result
}

func query() {
	db, sql := buildFindOneSql()

	//生成SQL
	sql = db.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Where("name like ?", "Trump%").Limit(50).Order("age desc").Find(&[]User{})
	})
	fmt.Println(sql)

	// 原生 SQL,多行
	rows, err := db.Raw(sql).Rows()
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	//将 sql.Rows 扫描至 model
	rows, err = db.Raw(sql).Rows()
	defer rows.Close()
	var user User
	for rows.Next() {
		//scan至变量
		db.ScanRows(rows, &user)
		//模拟业务逻辑
		fmt.Println(":" + user.Username + ":" + strconv.Itoa(user.Age))
	}
}

func buildFindOneSql() (*gorm.DB, string) {
	db := getDb()

	//生成sql
	sql := db.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Model(&User{}).Where("id = ?", 100).Limit(10).Order("age desc").Find(&[]User{})
	})
	fmt.Println(sql)
	return db, sql
}
