package webserver

import (
	"rentcar/veiculo"

	"github.com/gin-gonic/gin"
)

//external function
func New() *gin.Engine {
	return startServer(gin.New())

}

//internal function
func startServer(r *gin.Engine) *gin.Engine {
	stgVeiculo := CreateDB()
	//agroupa os endpoints
	v1 := r.Group("api/v1")
	//configurar endpoints do veiculo
	handler := veiculo.NewVeiculo(stgVeiculo)
	v1.GET("/veiculos", handler.Get)
	v1.POST("/veiculos", handler.Create)
	v1.PUT("/veiculos", handler.Update)
	//passagem de parametros url = http://localhost:8080/api/v1/veiculos/1 => ID
	v1.DELETE("/veiculos", handler.Delete)

	return r
}
func CreateDB() veiculo.MySQLStorage {
	conn := "root: (127.0.0.0.1)/veiculos"
	return veiculo.NewStorage(conn)
}
