package testutils

import (
	"context"
	"fmt"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
)

const (
	pgImage    = "postgres:16.2"
	pgUsername = "root"
	pgPassword = "password"
	pgDBName   = "rd_clone_api"

	ecrHost = "public.ecr.aws/docker/library/"
)

type Container interface {
	ConnectionString() string
}

type PGContainer struct {
	*postgres.PostgresContainer
}

func (p *PGContainer) ConnectionString() string {
	ctx := context.TODO()
	connStr := p.MustConnectionString(ctx, "sslmode=disable")
	return connStr
}

func CreatePGContainer() Container {
	return createContainer(postgresContainer)
}

func postgresContainer() Container {
	ctx := context.TODO()
	pgCont, err := postgres.Run(ctx, pgImage,
		postgres.WithDatabase(pgDBName),
		postgres.WithUsername(pgUsername),
		postgres.WithPassword(pgPassword),
		postgres.BasicWaitStrategies(),
		withName("PGC"), mustReuse(),
		testcontainers.WithImageSubstitutors(hostSubstitutor{}))
	if err != nil {
		panic("Failed to start test postgres container: " + err.Error())
	}

	return &PGContainer{
		PostgresContainer: pgCont,
	}
}

type hostSubstitutor struct {
	ImageHost string
}

func (s hostSubstitutor) Description() string {
	return fmt.Sprintf("Prepending %s host to image name", s.ImageHost)
}

func (s hostSubstitutor) Substitute(image string) (string, error) {
	s.ImageHost = ecrHost
	return s.ImageHost + image, nil
}

func createContainer(createFunc func() Container) Container {
	return createFunc()
}

func withName(name string) testcontainers.CustomizeRequestOption {
	return func(req *testcontainers.GenericContainerRequest) error {
		req.Name = name
		return nil
	}
}

func mustReuse() testcontainers.CustomizeRequestOption {
	return func(req *testcontainers.GenericContainerRequest) error {
		req.Reuse = true
		return nil
	}
}
