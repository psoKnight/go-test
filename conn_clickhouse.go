package main

import (
	"fmt"
	_ "github.com/ClickHouse/clickhouse-go"
	"github.com/jmoiron/sqlx"
)

func main() {
	source := "tcp://10.117.48.122:8123?username=default&password=megvii2022&database=default"
	connect, err := sqlx.Connect("clickhouse", source)
	if err != nil {
		fmt.Printf("clickhouse open err %s", err.Error())
	}

	fmt.Println()
	fmt.Println("connect:", connect)
}
