package server

import (
	"context"
	"net/http"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/L4B0MB4/JRNY/pkg/configuration"
	"github.com/L4B0MB4/JRNY/pkg/models"
	"github.com/L4B0MB4/JRNY/pkg/server/pool"
	"github.com/L4B0MB4/JRNY/pkg/server/pool/factory"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type InternalServerError struct {
	Error string `json:"error"`
}

func onHealthCheck(c *gin.Context) {

	c.Writer.Write([]byte("hello"))
	c.Status(200)
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
	model.ReceivedAt = time.Now().UTC()
	eventPool.Enqueue(&model)
	c.Status(201)
}

var eventPool = pool.EventPool{}

func Start(config *configuration.ServerConfiguration, factory factory.EventPoolWorkerFactory) {
	signalCtx, stop := signal.NotifyContext(context.Background(), syscall.SIGKILL, syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	ctx, cancel := context.WithCancel(context.Background())
	eventPool.Initialize(factory, ctx)
	router := gin.Default()
	router.GET("/", onHealthCheck)
	router.POST("/api/event", onRequest)
	srv := &http.Server{
		Addr:    config.HttpConfig.Host + ":" + strconv.Itoa(config.HttpConfig.Port),
		Handler: router,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("error starting server")
		}
	}()
	defer cancel()
	<-signalCtx.Done()
	log.Debug().Msg("Shutting down ...")
	cancel()
	srv.Shutdown(ctx)
}
