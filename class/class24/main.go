package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//badCode()
	//goodCode()
	connDB()

}

func connDB() {
	dbName := "root:@tcp(127.0.0.1:3306)/learn_go"
	learnDB, err := sql.Open("mysql", dbName)
	if err != nil {
		fmt.Println("链接数据库失败:", err)
	}
	defer learnDB.Close()
	if err = learnDB.Ping(); err != nil {
		fmt.Println("DB 测试失败:", err)
	}
	
}

// 坏代码，面条代码，场景一般是第一次写代码，快速实现业务的存在，耦合度很高，代码维护很困难
func badCode() {
	CheckSomething()
}

// 优化过的好代码，业务和技术分开，解耦，耦合度很低，代码维护很简单
func goodCode() {
	CheckSomething2(operations)
}
