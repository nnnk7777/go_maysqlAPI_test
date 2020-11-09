// 2020/11/3
// MAMPで起動したMySQLを操作するテスト

package main

import (
    "fmt"

    _ "github.com/go-sql-driver/mysql"
    "github.com/jinzhu/gorm"
)

func main() {
    // db接続
    db, err := sqlConnect()
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()

    // error := db.Create(&Dogs{
    //   ImgPath: "test",
    //   Description: "https://dummyimage.com/600x400/000/fff",
    // }).Error
    // if error != nil {
    //     fmt.Println(error)
    // } else {
    //     fmt.Println("データ追加成功")
    // }

    result := []*Dogs{}
    error := db.Find(&result).Error
    if error != nil || len(result) == 0 {
        return
    }
    for _, dog := range result {
        fmt.Println(dog.Description)
    }
}

// SQLConnect DB接続
func sqlConnect() (database *gorm.DB, err error) {
    DBMS := "mysql"
    USER := "root"
    PASS := "root"
    PROTOCOL := "tcp(localhost:8889)"
    DBNAME := "test_db"

    CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
    return gorm.Open(DBMS, CONNECT)
}


// テーブル情報
type Dogs struct {
  ImgPath string `json: "img_path"`
  Description string `json: "description"`
}
