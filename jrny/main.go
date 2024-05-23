package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"

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
	signalCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	ctx, cancel := context.WithCancel(context.Background())
	factory := &pool.DefaultEventPoolWorkerFactory{}
	factory.UseQueueWorker("")
	eventPool.Initialize(factory, ctx)
	router := gin.Default()
	router.POST("/api/event", onRequest)
	srv := &http.Server{
		Addr:    ":8081",
		Handler: router,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("error starting server")
		}
	}()
	defer cancel()
	select {
	case <-signalCtx.Done():
	}
	log.Debug().Msg("Shutting down ...")
	cancel()
	srv.Shutdown(ctx)

}
