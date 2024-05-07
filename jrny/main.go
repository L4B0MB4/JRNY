package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type Model struct {
	Name string `json:"name" binding:"required"`
}
type InternalServerError struct {
	Error string `json:"error"`
}

func onRequest(c *gin.Context) {
	var model Model
	err := c.ShouldBindBodyWithJSON(&model)
	if err != nil {
		log.Err(err)
		iSE := InternalServerError{
			Error: err.Error(),
		}
		c.JSON(500, &iSE)
		return

	}
	c.Status(201)

	log.Debug().Str("method", c.Request.Method).Str("path", c.Request.RequestURI).Interface("body", model).Msg("On Request")
}

func main() {

	router := gin.Default()
	router.POST("/api/event", onRequest)
	router.Run(":8080")

}
