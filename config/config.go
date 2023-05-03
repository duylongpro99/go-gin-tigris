package config

import (
	"os"
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
