package app

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Context interface {
	Bind(v any) error
	OK(v any)
	BadRequest(err error)
}

type context struct {
	*gin.Context
}

func (c *context) Bind(v any) error {
	return c.Context.ShouldBind(v)
}

func (c *context) OK(v any) {
	c.JSON(http.StatusOK, v)
}

func (c *context) BadRequest(err error) {
	c.AbortWithError(http.StatusBadRequest, err)
}

func NewHandler(handler func(Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler(&context{c})
	}
}

func NewRouter() *gin.Engine {
	r := gin.Default()

	config := cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"X-Requested-With", "Authorization", "Origin", "Content-Length", "Content-Type", "TransactionID"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}

	r.Use(cors.New(config))
	return r
}
