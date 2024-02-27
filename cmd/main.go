package main

import (
	"go-make-tests/internal/db"
	"go-make-tests/internal/env"
	creds "go-make-tests/internal/svc/credentials"
	"go-make-tests/internal/ui"

	"log"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	env.LoadEnv()
}

func main() {
	// Connect()
	ui.Run()

}

func Connect() {
	db := db.NewCredsDB()
	credsService := creds.New(db)

	allCreds, err := credsService.GetAllCredentials()
	env.LogFatalErr(err)

	for _, c := range allCreds {
		log.Printf("user: %s, access: %s, secret: %s\n", *c.User, *c.AccessKey, *c.SecretKey)
	}
}
