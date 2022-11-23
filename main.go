package main

import (
	"gin_es-rabbit/controllers"
	"gin_es-rabbit/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Gin Rabbit Elastic
// @version         1.0
// @description     This is a sample server celler server.

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:2050
func main() {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	docs.SwaggerInfo.BasePath = ""

	apiV1 := router.Group("/api/v1")
	controllers.NewBookController(apiV1)

	err := router.Run(":2050")
	if err != nil {
		return
	}

}
