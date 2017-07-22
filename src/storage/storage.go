package storage

import (
    "fmt"
    "time"
    "log"
    "github.com/go-xorm/core"
    "github.com/go-xorm/xorm"
    _ "github.com/go-sql-driver/mysql"
)

type Result struct {
    Id int `xorm:"pk autoincr"`
    ExternalId int
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

func NewEngine() (*xorm.Engine, error) {
    engine, err := xorm.NewEngine("mysql", "root:root@tcp(192.168.2.108:3306)/hopper?charset=utf8")
    if err != nil {
        return engine, err
    }

    engine.SetMapper(core.SameMapper{})

    return engine, nil
}

// Init storage
func Init() {
    fmt.Print("Init storage\n")

    var engine, err = NewEngine()
    if err != nil {
        log.Fatal(err)
        return
    }

    err = engine.Sync2(new(Result), new(Proxy))
    if err != nil {
        log.Fatal(err)
    }
    fmt.Print("Done\n")
}

func Save(object interface{}) int64 {
    engine, err := NewEngine()
    if err != nil {
        log.Fatal(err)
        return 0
    }

    cols, err := engine.Insert(object)
    if err != nil {
        log.Fatal(err)
        return 0
    }

    return cols
}

func Read() []Result {
    engine, err := NewEngine()
    if err != nil {
        log.Fatal(err)
        return nil
    }

    var results []Result
    err = engine.Table("Result").Select("*").
         Find(&results)
    return results
}