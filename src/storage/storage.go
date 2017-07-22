package storage

import (
    "fmt"
    "time"
    "log"
    "github.com/go-xorm/core"
    "github.com/go-xorm/xorm"
    _ "github.com/go-sql-driver/mysql"
    configReader "config"
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
    var config = configReader.ForDatabase()

    engine, err := xorm.NewEngine("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8", config.User, config.Password, config.Host, config.Name))
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

func GetProxyByCountry(country string) (*Proxy, error) {
    engine, err := NewEngine()
    if err != nil {
        log.Fatal(err)
        return nil, err
    }

    var proxy= new(Proxy)
    engine.Where("country = ?", country).Get(proxy)

    return proxy, nil
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