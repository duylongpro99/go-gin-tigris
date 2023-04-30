package config

import (
	"context"
	"os"

	"github.com/tigrisdata/tigris-client-go/tigris"
)

type TigrisEnv struct {
	URL          string
	Name         string
	ClientId     string
	ClientSecret string
	Branch       string
}

type TigrisConfig interface {
	GetEnv() TigrisEnv
}

func GetTigrisEnv() *TigrisEnv {
	tigris_env := &TigrisEnv{
		URL:          os.Getenv("TG_URL"),
		Name:         os.Getenv("TG_NAME"),
		ClientId:     os.Getenv("TG_CLIENT_ID"),
		ClientSecret: os.Getenv("TG_CLIENT_SECRET"),
		Branch:       os.Getenv("TG_BRANCH"),
	}
	return tigris_env
}

func ConnectTigris(ctx context.Context) (*tigris.Client, error) {
	tigris_env := GetTigrisEnv()
	cfg := &tigris.Config{
		URL:          tigris_env.URL,
		ClientID:     tigris_env.ClientId,
		ClientSecret: tigris_env.ClientSecret,
		Project:      tigris_env.Name,
	}

	client, err := tigris.NewClient(ctx, cfg)
	return client, err
}
