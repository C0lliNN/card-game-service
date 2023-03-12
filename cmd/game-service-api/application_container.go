package main

import (
	"C0lliNN/card-game-service/internal/config"
	"C0lliNN/card-game-service/internal/game"
	"C0lliNN/card-game-service/internal/generator"
	"C0lliNN/card-game-service/internal/persistence"
	"C0lliNN/card-game-service/internal/server"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
	"os"
	"time"
)

func newServer() *server.Server {
	appConfig := readConfig()
	db := config.NewMongoDatabase(appConfig.Database.URI, appConfig.Database.Name)
	deckRepo := persistence.NewDeckRepository(db)
	idGenerator := generator.NewUUIDGenerator()
	gameService := game.NewGame(deckRepo, idGenerator)
	engine := gin.New()
	return server.New(server.Config{
		Router:          engine,
		Addr:            appConfig.Server.Address,
		Timeout:         time.Second * time.Duration(appConfig.Server.Timeout),
		GameHandler:     server.NewGameHandler(gameService),
		ErrorMiddleware: server.NewErrorMiddleware(),
	})
}

func readConfig() config.Config {
	configPath := "./local.yml"
	if os.Getenv("CONFIG_PATH") != "" {
		configPath = os.Getenv("CONFIG_PATH")
	}

	f, err := os.Open(configPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var appConfig config.Config

	decoder := yaml.NewDecoder(f)
	if err = decoder.Decode(&appConfig); err != nil {
		panic(err)
	}

	return appConfig
}
