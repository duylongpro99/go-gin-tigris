package config

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/tigrisdata/tigris-client-go/tigris"
)

var DB *tigris.Database
var once sync.Once

func InitTigris() {
	if DB == nil {
		once.Do(func() {
			DB = connectDatabase()
			fmt.Println("Server has connected to tigris")
		})
	} else {
		fmt.Println("Server has already connect to tigris")
	}

}

func connectDatabase() *tigris.Database {
	// get tigris config
	tigris_env := &TigrisEnv{}
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
	})

	if err != nil {
		panic(err)
	}

	return db
}
