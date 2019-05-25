package main

import (
	"erlink/paste"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo"
)

// 构架
// 短链接部分
// 图床部分
// 剪切板部分
var e *echo.Echo
func main() {
	e:=echo.New()
	db, err := gorm.Open("sqlite3", "gorm.db")
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}
	paste.Initialize(e, db)

}
