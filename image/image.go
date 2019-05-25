package image

import "github.com/labstack/echo"

var e *echo.Echo

func upload(e echo.Context) error {
	return nil
}
func show(e echo.Context) error {
	return nil
}
func Initialize(e2 *echo.Echo) {
	e = e2
	e.POST("/upload", upload)
	e.GET("/upload", upload)
	e.POST("/show", show)
	e.GET("/show", show)

}
