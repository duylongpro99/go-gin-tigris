package main

import (
	"fmt"
	"log"
	"os"

	"server/config"
	. "server/config"
	docs "server/docs"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @BasePath /api/v1
// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Ping
// @Router /ping [get]
// r.GET("/api/v1/ping", Ping)
func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "success",
	})
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln("Cant not load env")
	}
	fmt.Println("Load env success")

	// get tigris config
	tigris_env := &config.TigrisEnv{}
	tigris_env = GetTigrisEnv()

	fmt.Println("")
	fmt.Println("**************TIGRIS*****************")
	fmt.Println("*************************************")
	fmt.Println("TIGRIS DATABASE:", tigris_env.Name)
	fmt.Println("TIGRIS BRANCH:", tigris_env.Branch)
	fmt.Println("*************************************")
	fmt.Println("*************************************")
	fmt.Println("")

	r := gin.Default()
	// Set up swagger
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		v1.GET("/ping", Ping)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
