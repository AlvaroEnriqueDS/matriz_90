package matriz_90

import (
	"github.com/labstack/echo"
)

// Route de matriz
func Route(e *echo.Echo) {
	r := e.Group("/api/v1/matriz")

	//implement
	s := store{}
	h := NewHandler(s)
	r.POST("", h.process)
}