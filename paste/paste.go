package paste

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"net/http"
)

var e *echo.Echo

func get(e echo.Context) error {
	if e.Request().Method == "GET"{

	}else if e.Request().Method == "POST"{

	}else{
		return e.String(http.StatusOK,"Method is not Allow.")
	}
	return e.String(http.StatusOK,"Method is not Allow.")
}
func show(e echo.Context) error {
	return nil
}

func Initialize(e2 *echo.Echo, db2 *gorm.DB) {
	e = e2
	db = db2
	CreateTables()
	var p paste
	p.Title="aaa"
	p.Content="aaa"
	Create(&p)
	fmt.Println(p.ID,p.Hash)
	e.Any("/paste/",show)
	e.Any("/paste/show/:hash",get)
}
