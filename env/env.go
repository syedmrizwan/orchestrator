package env

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type envFile struct {
	Domain         string
	Service     string
	Host     string
	ApplicationName         string
	ClientName string
}

var Env *envFile

func init(){
	_ = godotenv.Load()

	Env = &envFile{
		Domain:         os.Getenv("DOMAIN"),
		Service:     os.Getenv("SERVICE"),
		Host:     os.Getenv("HOST"),
		ApplicationName:         os.Getenv("APPLICATION_NAME"),
		ClientName: os.Getenv("CLIENT_NAME"),
	}

	fmt.Println(Env)
}