package main

import (
	"erlink/paste"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"html/template"
	"io"
	"log"
)

// 构架
// 短链接部分
// 图床部分
// 剪切板部分
var e *echo.Echo

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()
	db, err := gorm.Open("sqlite3", "gorm.db")
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}
	paste.Initialize(e, db)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}
	e.Renderer = renderer
	e.Static("/static", "templates/static")
	log.Fatal(e.Start(":8000"))
}
