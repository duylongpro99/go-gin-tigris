package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"server/config"
	. "server/config"
	docs "server/docs"
	models "server/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/tigrisdata/tigris-client-go/tigris"
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
	r := gin.Default()

	// Set up swagger
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		v1.GET("/ping", Ping)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
	fmt.Println("TIGRIS DATABASE:", tigris_env.Name)
	fmt.Println("TIGRIS BRANCH:", tigris_env.Branch)
	fmt.Println("*************************************")
	fmt.Println("")
	// Set up tigris
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	db, err := tigris.OpenDatabase(ctx, &tigris.Config{
		URL:          tigris_env.URL,
		ClientID:     tigris_env.ClientId,
		ClientSecret: tigris_env.ClientSecret,
		Project:      tigris_env.Name,
	}, &models.User{})

	if err != nil {
		panic(err)
	}

	fmt.Println("Connected db")
	models.SetCRUDRoutes[models.User](r, db, "users")

	r.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
