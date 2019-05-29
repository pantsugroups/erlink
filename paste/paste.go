package paste

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	"time"
)

var e *echo.Echo

func push(c echo.Context) error {
	if c.Request().Method == "GET" {
		return c.Render(http.StatusOK, "paste.get.html", nil)
	} else if c.Request().Method == "POST" {
		err := c.Request().ParseForm()
		if err != nil {
			return err
		}
		var newPaste paste
		newPaste.Title = c.Request().Form["title"][0]
		newPaste.Content = c.Request().Form["code"][0]
		newPaste.Language = c.Request().Form["language"][0]
		switch c.Request().Form["long"][0] {
		case "day":
			m, _ := time.ParseDuration("24h")
			newPaste.DeadLine = time.Now().Add(m)
		case "week":
			m, _ := time.ParseDuration("168h")
			newPaste.DeadLine = time.Now().Add(m)
		case "year":
			m, _ := time.ParseDuration("8760h")
			newPaste.DeadLine = time.Now().Add(m)
		case "unlimited":
			m, _ := time.ParseDuration("99999h")
			newPaste.DeadLine = time.Now().Add(m)
		default:
			m, _ := time.ParseDuration("24h")
			newPaste.DeadLine = time.Now().Add(m)
		}

		err = Create(&newPaste)
		if err != nil {
			return err
		}
		fmt.Println(newPaste)
	} else {
		return c.String(http.StatusOK, "Method is not Allow.")
	}
	return c.String(http.StatusOK, "Method is not Allow.")
}
func show(c echo.Context) error {
	hash := c.Param("hash")
	if hash == "" {
		return c.Render(http.StatusBadRequest, "error.html", nil)
	}
	var t template
	var p paste
	uhash, err := strconv.ParseUint(hash, 0, 32)
	if err != nil {
		return err
	}
	err = Get(uhash, &p)
	if err != nil {
		return err
	}
	if p == (paste{}) {
		return c.Render(http.StatusBadRequest, "error.html", nil)
	}
	t.Time = p.CreatedAt.String()
	t.Long = p.DeadLine.String()
	t.Title = p.Title
	t.Code = p.Content
	t.Language = p.Language
	return c.Render(http.StatusOK, "paste.show.html", t)

}

func Initialize(e2 *echo.Echo, db2 *gorm.DB) {
	e = e2
	db = db2
	CreateTables()

	e.GET("/paste/show/:hash", show)
	e.Match([]string{"GET", "POST"}, "/paste/push", push)
}
