package main

import (
	_ "github.com/aceslick911/bcg-checkout/docs"
	"github.com/aceslick911/bcg-checkout/internal/api"
)

// @BCG Api
// @version 1.0
// @description API REST for BCG

// @contact.name Angelo Perera
// @contact.url http://github.com/aceslick911
// @contact.email angeloperera@gmail.com

// @license.name MIT
// @license.url http://github.com/aceslick911

// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	api.Run("")
}
