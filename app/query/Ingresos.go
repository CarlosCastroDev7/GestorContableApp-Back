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
	json.Unmarshal(archivo, &listIngresos)

	listIngresos = append(listIngresos, models.ListIngresosType{
		Fecha:    "95-52-2022",
		Concepto: "nuevo",
		Valor:    1254,
	})

	ctx.JSON(http.StatusOK, listIngresos)
}
