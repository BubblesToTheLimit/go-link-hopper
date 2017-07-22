package storage

import (
    _ "github.com/go-sql-driver/mysql"
    "fmt"
    "time"
    "github.com/go-xorm/xorm"
    "log"
)

type Result struct {
    Id int `xorm:"pk"`
    Url string
    Target string
    Trace []string
    CreatedAt time.Time `xorm:"created"`
    Error error
}

type Proxy struct {
    Country string `xorm:"pk varchar(2)"`
    IP string
    CreatedAt time.Time `orm:"created"`
}

// Init storage
func Init() {
    fmt.Print("Init storage\n")

    var engine, err = xorm.NewEngine("mysql", "root:root@127.0.0.1:3306/hopper?charset=utf8")
    if err != nil {
        log.Fatal(err)
        return
    }

    engine.Sync2(new(Result), new(Proxy))
    fmt.Print("Done\n")
}