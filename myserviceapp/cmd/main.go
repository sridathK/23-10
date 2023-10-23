package main

import (
	"fmt"
	"myserviceapp/auth"
	"myserviceapp/handlers"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
)

func main() {
	// calling function to start app
	err := startApp()
	if err != nil {
		log.Panic().Err(err).Send()
	}

}

func startApp() error {
	log.Info().Msg("started main")
	privatePEM, err := os.ReadFile(`C:\Users\ORR Training 1\Desktop\myserviceapp\private.pem`)
	if err != nil {
		return fmt.Errorf("cannot find file private.pem %w", err)
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privatePEM)
	if err != nil {
		return fmt.Errorf("cannot convert byte to key %w", err)
	}

	publicPEM, err := os.ReadFile(`C:\Users\ORR Training 1\Desktop\myserviceapp\pubkey.pem`)
	if err != nil {
		return fmt.Errorf("cannot find file pubkey.pem %w", err)
	}
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicPEM)
	if err != nil {
		return fmt.Errorf("cannot convert byte to key %w", err)
	}
	a, err := auth.NewAuth(privateKey, publicKey)
	if err != nil {
		return fmt.Errorf("cannot create auth instance %w", err)
	}
	//server connections
	api := http.Server{ //server config and settimngs
		Addr:    ":8081",
		Handler: handlers.Api(a),
	}
	api.ListenAndServe()

	return nil
}
