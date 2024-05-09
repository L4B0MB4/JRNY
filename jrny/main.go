package main

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/L4B0MB4/JRNY/jrny/models"
	"github.com/L4B0MB4/JRNY/jrny/pool"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type InternalServerError struct {
	Error string `json:"error"`
}

func onRequest(c *gin.Context) {
	var model models.Event
	err := c.ShouldBindBodyWithJSON(&model)
	if err != nil {
		log.Err(err)
		iSE := InternalServerError{
			Error: err.Error(),
		}
		c.JSON(500, &iSE)
		return

	}
	eventPool.Enqueue(&model)
	c.Status(201)
}

var eventPool = pool.EventPool{}

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	eventPool.Initialize()
	router := gin.Default()
	router.POST("/api/event", onRequest)
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("error starting server")
		}
	}()
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*100)
	defer cancel()
	time.Sleep(time.Second * 10)
	eventPool.Shutdown()
	srv.Shutdown(ctx)

}
