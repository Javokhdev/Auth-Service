package api

import (
	"github.com/Javokhdev/Auth-Service/api/handler"
	"github.com/Javokhdev/Auth-Service/api/middleware"
 	_ "github.com/Javokhdev/Auth-Service/docs"
	_ "github.com/gin-contrib/cors"


	"github.com/gin-gonic/gin"
	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Auth-Service API
// @version 1.0
// @description API for Auth-Service
// @host localhost:8081
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authourization

func NewGin(h *handler.Handler) *gin.Engine {
	r := gin.Default()
	
	
	r.Use(middleware.MiddleWare())
	u := r.Group("/user")
	u.POST("/registr", h.RegisterUser)
	u.PUT("/update/:id", h.UpdateUser)
	u.DELETE("/delete/:id", h.DeleteUser)
	u.GET("/getall", h.GetAllUser)
	u.GET("/getbyid/:id", h.GetbyIdUser)
	u.POST("/login", h.LoginUser)
	url := ginSwagger.URL("swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler, url))
	return r
}
