package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	puerto = ":8080"
)

// estructura persona
type Persona struct {
	Nombre    string `json:"nombre"`
	Apellido  string `json:"apellido"`
	Edad      int    `json:"edad"`
	Direccion string `json:"direccion"`
	Telefono  string `json:"telefono"`
	Activo    bool   `json:"activo"`
}

func main() {
	jsonPersona := `{
		"nombre" : "Juan",
		"apeliido": "Perez",
		"edad":25,
		"direccion": "Av. 2345",
		"telefono":"123454567",
		"activo":true
	}`

	var persona Persona

	err := json.Unmarshal([]byte(jsonPersona), &persona)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(persona)

	engine := gin.Default()
	engine.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"mensaje": "pong",
		})
	})
	personaResponse := Persona{
		Nombre:    "aa",
		Apellido:  "ss",
		Edad:      22,
		Direccion: "ad",
		Telefono:  "56698",
		Activo:    true,
	}

	engine.GET("/persona", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"data": personaResponse,
		})

	})

	err = engine.Run(puerto)
	if err != nil {
		log.Fatal(err)
	}
}
