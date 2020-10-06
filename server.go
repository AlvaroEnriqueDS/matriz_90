package matriz_90

import (
	"github.com/labstack/echo"
	"log"
)

//Inicia el servicio
func Start(){
	e := echo.New()

	Route(e)

	log.Println("Servidor iniciando en http://127.0.0.1:8080")
	err := e.Start(":8080")
	if err != nil {
		log.Printf("Error en el servidor: %v\n", err)
	}
}
