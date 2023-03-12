package test

import (
	"C0lliNN/card-game-service/internal/config"
	"C0lliNN/card-game-service/internal/game"
	"C0lliNN/card-game-service/internal/generator"
	"C0lliNN/card-game-service/internal/persistence"
	"C0lliNN/card-game-service/internal/server"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"net/http"
	"time"
)

type ServerTestSuite struct {
	RepositoryTestSuite
	BaseURL string
	Server  *server.Server
}

func (s *ServerTestSuite) SetupSuite() {
	s.RepositoryTestSuite.SetupSuite()

	testConfig := s.createTestConfig()
	testServer := s.createTestServer(testConfig)
	s.Server = testServer

	go func() {
		err := testServer.Start()
		require.Equal(s.T(), http.ErrServerClosed, err)
	}()

	require.Eventually(s.T(), func() bool {
		response, err := http.Get(s.BaseURL + "/decks/deck-id")
		if err != nil {
			return false
		}

		return response.StatusCode == http.StatusNotFound
	}, time.Second*20, time.Millisecond*200)
}

func (s *ServerTestSuite) createTestConfig() config.Config {
	baseUrl := "localhost:9000"
	s.BaseURL = fmt.Sprintf("http://%s", baseUrl)

	return config.Config{
		Database: struct {
			URI  string
			Name string
		}{
			URI:  s.Endpoint,
			Name: "card-game",
		},
		Server: struct {
			Address string
			Timeout int
		}{
			Address: baseUrl,
			Timeout: 30,
		},
	}
}

func (s *ServerTestSuite) createTestServer(c config.Config) *server.Server {
	db := config.NewMongoDatabase(c.Database.URI, c.Database.Name)
	deckRepo := persistence.NewDeckRepository(db)
	idGenerator := generator.NewUUIDGenerator()
	gameService := game.NewGame(deckRepo, idGenerator)
	engine := gin.New()
	return server.New(server.Config{
		Router:          engine,
		Addr:            c.Server.Address,
		Timeout:         time.Second * time.Duration(c.Server.Timeout),
		GameHandler:     server.NewGameHandler(gameService),
		ErrorMiddleware: server.NewErrorMiddleware(),
	})
}

func (s *ServerTestSuite) TearDownSuite() {
	s.RepositoryTestSuite.TearDownSuite()
	_ = s.Server.Shutdown(context.Background())
}
