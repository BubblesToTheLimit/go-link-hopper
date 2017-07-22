package storage

import (
	_ "github.com/go-sql-driver/mysql"
	//"github.com/go-xorm/xorm"
	//"log"
	"fmt"
)

// Init storage
func Init() {
	fmt.Print("Init storage")
	//var engine, err = xorm.NewEngine("mysql", "hopper:go@127.0.0.1:3306?charset=utf8")
	//if err != nil {
	//    log.Fatal(err)
	//}
}