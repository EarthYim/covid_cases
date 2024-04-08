package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Context interface {
	OK(v any)
	InternalServerError(error)
}

type appContext struct {
	*gin.Context
}

func (c *appContext) OK(v any) {
	c.JSON(http.StatusOK, v)
}

func (c *appContext) InternalServerError(err error) {
	c.Error(err)
}

func NewHandler(handler func(Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler(&appContext{c})
	}
}

func NewRouter() *gin.Engine {
	r := gin.Default()

	config := cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET"},
		// AllowHeaders:     []string{"X-Requested-With", "Authorization", "Origin", "Content-Length", "Content-Type", "TransactionID"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}

	r.Use(cors.New(config))
	return r
}

func Run(r *gin.Engine, port string) {
	srv := http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownRelease()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		fmt.Printf("failed to shutdown gracefully: %v\n", err)
	}
	fmt.Println("shutting down")
}
