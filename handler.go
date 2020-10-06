package matriz_90

import (
	"github.com/labstack/echo"
)

type handler struct {
	store service
}

func NewHandler(store service) *handler {
	return &handler{store}
}

//process volcado de data y retorno de consulta
func (h *handler) process(c echo.Context) error {
	data := model{}
	err := c.Bind(&data)
	if err != nil {
		return c.JSON(400, "NO SE PUEDE VOLCAR LA DATA")
	}

	result, err := h.store.processService(&data)
	if err != nil {
		return c.JSON(400, err.Error())
	}
	data.Matriz = result


	return c.JSON(200, data)
}