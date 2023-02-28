package mutation

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gestor-gastos/app/models"
	"github.com/gin-gonic/gin"
)

func InsertIngresos(ctx *gin.Context) {

	var resp, listIngresos []models.ListIngresosType

	err := ctx.ShouldBindJSON(&resp)
	if err != nil {
		log.Printf("Error al obtener el JSON %s\n", err)
		ctx.JSON(http.StatusNotAcceptable, gin.H{
			"Error": err,
		})
	}

	archivo, err := os.ReadFile("./JSONS/ingresos.json")
	if err != nil {
		log.Printf("Error al leer el archivo %s\n", err)
		ctx.JSON(http.StatusNotFound, gin.H{
			"Error": err,
		})
	}

	err = json.Unmarshal(archivo, &listIngresos)
	if err != nil {
		log.Printf("Error al extrar la informacion %s\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": err,
		})
	}

	listIngresos = append(listIngresos, resp...)

	datos, err := json.Marshal(listIngresos)
	if err != nil {
		log.Printf("Error generando el string json encoding %s\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": err,
		})
	}

	err = os.WriteFile("./JSONS/ingresos.json", datos, 0222)
	if err != nil {
		log.Printf("Error al escribir el archivo de ingresos %s\n", err)
		ctx.JSON(http.StatusNotFound, gin.H{
			"Error": err,
		})
	}

	ctx.JSON(http.StatusOK, listIngresos)

}
