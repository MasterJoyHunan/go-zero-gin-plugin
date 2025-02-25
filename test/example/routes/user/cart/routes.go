// Code generated by goctl. DO NOT EDIT.
package cart

import (
	"github.com/MasterJoyHunan/gengin/test/example/handler/user/cart"
	"github.com/MasterJoyHunan/gengin/test/example/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterUserCartRoute(e *gin.Engine) {
	g := e.Group("/v1")
	g.Use(middleware.AuthMiddleware, middleware.SomeMiddleware, middleware.CorsMiddleware)
	g.GET("/cart", cart.GetCartListHandle)
	g.GET("/cart/:id", cart.GetCartHandle)
	g.POST("/cart", cart.AddCartHandle)
	g.PUT("/cart/:id", cart.EditCartHandle)
}
