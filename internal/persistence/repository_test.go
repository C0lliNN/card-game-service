package persistence

import (
	"context"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const (
	mongoImage  = "mongo:5"
	exposedPort = "27017/tcp"
	dbName      = "card-game"
)

type mongodbContainer struct {
	testcontainers.Container
}

func newContainer(ctx context.Context) (*mongodbContainer, error) {
	req := testcontainers.ContainerRequest{
		Image:        mongoImage,
		ExposedPorts: []string{exposedPort},
		WaitingFor: wait.ForAll(
			wait.ForLog("Waiting for connections"),
			wait.ForListeningPort(exposedPort),
		),
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, err
	}

	return &mongodbContainer{Container: container}, nil
}

type RepositoryTestSuite struct {
	suite.Suite
	container *mongodbContainer
	db        *mongo.Database
}

func (s *RepositoryTestSuite) SetupSuite() {
	ctx := context.TODO()

	var err error
	s.container, err = newContainer(ctx)
	require.NoError(s.T(), err)

	endpoint, err := s.container.Endpoint(ctx, "mongodb")
	require.NoError(s.T(), err)

	mongoClient, err := mongo.NewClient(options.Client().ApplyURI(endpoint))
	require.NoError(s.T(), err)

	err = mongoClient.Connect(ctx)
	require.NoError(s.T(), err)

	require.Eventually(s.T(), func() bool {
		err = mongoClient.Ping(ctx, nil)
		return err == nil
	}, time.Second*10, time.Millisecond*100)

	s.db = mongoClient.Database(dbName)
}

func (s *RepositoryTestSuite) TearDownTest() {
	ctx := context.TODO()

	_ = s.db.Drop(ctx)
}

func (s *RepositoryTestSuite) TearDownSuite() {
	ctx := context.TODO()

	_ = s.container.Terminate(ctx)
}
