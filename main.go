package main

import (
	"fmt"
	"log"
	"os"

	"server/config"
	. "server/config"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

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
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "success",
		})
	})
	r.Run(fmt.Sprintf("localhost:%s", os.Getenv("PORT")))
	fmt.Println("Server is running")
}
