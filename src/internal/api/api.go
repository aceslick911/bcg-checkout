package api

import (
	"fmt"

	"github.com/aceslick911/bcg-checkout/internal/api/router"
	"github.com/aceslick911/bcg-checkout/internal/pkg/config"
	"github.com/aceslick911/bcg-checkout/internal/pkg/db"
	"github.com/aceslick911/bcg-checkout/internal/pkg/persistence"
	"github.com/gin-gonic/gin"
)

func setConfiguration(configPath string) {
	config.Setup(configPath)
	db.SetupDB()
	persistence.HydrateDatabase()
	gin.SetMode(config.GetConfig().Server.Mode)
}

func Run(configPath string) {
	if configPath == "" {
		configPath = "data/config.yml"
	}
	setConfiguration(configPath)
	conf := config.GetConfig()
	web := router.Setup()
	fmt.Println("Go API REST Running on port " + conf.Server.Port)
	fmt.Println("==================>")
	_ = web.Run(":" + conf.Server.Port)
}
