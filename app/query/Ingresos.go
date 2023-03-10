package query

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gestor-gastos/app/models"
	"github.com/gin-gonic/gin"
)

func ListIngresos(ctx *gin.Context) {

	archivo, err := os.ReadFile("./JSONS/ingresos.json")
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusNotFound, gin.H{
			"Error": err,
		})
	}

	var listIngresos []models.ListIngresosType
	err = json.Unmarshal(archivo, &listIngresos)
	if err != nil {
		log.Printf("Error leyendo fichero: %s\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": err,
		})
	}

	datos, err := json.Marshal(listIngresos)
	if err != nil {
		log.Printf("Error leyendo fichero: %s\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": err,
		})
	}

	err = os.WriteFile("./JSONS/ingresos.json", datos, 0222)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusNotFound, gin.H{
			"Error": err,
		})
	}

	ctx.JSON(http.StatusOK, listIngresos)
}
