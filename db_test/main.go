package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// MySQL 数据库连接信息
	username := "root"
	password := "root_pwd"
	host := "db.dev.zenwell.cn"
	port := "6447"
	database := "jjj_shop_multi"

	//username := "user_test"
	//password := "175ac29d62931825950a9716e8fa2f18"
	//host := "120.27.225.90"
	//port := "3306"
	//database := "jjj_shop_multi"

	// 构建数据库连接字符串
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, database)

	// 记录连接开始时间
	startConnectTime := time.Now()

	// 连接数据库
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 记录连接结束时间
	endConnectTime := time.Now()

	// 计算连接花费的时间
	connectTime := endConnectTime.Sub(startConnectTime)

	fmt.Printf("连接时间 %s\n", connectTime)

	// 记录执行多条 SQL 语句开始时间
	startExecutionTime := time.Now()

	// 在这里可以执行多条数据库查询操作
	queries := []string{
		"SELECT * FROM jjj_shop_multi.jjjshop_employer_charge_record",
		//"INSERT INTO jjj_shop_multi.jjjshop_employer_charge_record (id, employer_id) VALUES (FLOOR(1111 + RAND() * 100), '1');",
		//"UPDATE jjj_shop_multi.jjjshop_employer_charge_record SET employer_id = '1' WHERE id = 1;",
	}

	for _, query := range queries {
		_, err := db.Exec(query)
		if err != nil {
			log.Fatal(err)
		}
	}

	// 记录执行多条 SQL 语句结束时间
	endExecutionTime := time.Now()

	// 计算执行多条 SQL 语句的总时间
	executionTime := endExecutionTime.Sub(startExecutionTime)

	fmt.Printf("sql执行时间 executed successfully in %s\n", executionTime)
}
