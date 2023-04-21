package app

import (
	"github.com/gin-gonic/gin"
	"github.com/rmscoal/go-restful-monolith-boilerplate/config"
	"github.com/rmscoal/go-restful-monolith-boilerplate/internal/composer"
	v1 "github.com/rmscoal/go-restful-monolith-boilerplate/internal/delivery/v1"
	"github.com/rmscoal/go-restful-monolith-boilerplate/pkg/doorkeeper"
	httpserver "github.com/rmscoal/go-restful-monolith-boilerplate/pkg/http"
	"github.com/rmscoal/go-restful-monolith-boilerplate/pkg/logger"
	"github.com/rmscoal/go-restful-monolith-boilerplate/pkg/postgres"
)

func Run(cfg *config.Config) {
	// Postgres .-.
	pg := postgres.GetPostgres(
		cfg.Db.URL,
		postgres.MaxPoolSize(cfg.Db.MaxPoolSize()),
		postgres.MaxOpenCoon(cfg.Db.MaxOpenConn()),
	)

	// Logger .-.
	logger := logger.NewAppLogger(cfg.App.LogPath)

	dk := doorkeeper.GetDoorkeeper(
		doorkeeper.RegisterSecretKey(cfg.Doorkeeper.SecretKey()),
		doorkeeper.RegisterSalt(cfg.Doorkeeper.HashSalt()),
		doorkeeper.RegisterHashMethod(cfg.Doorkeeper.HashMethod()),
		doorkeeper.RegisterSignMethod(cfg.Doorkeeper.SigningMethod(), cfg.Doorkeeper.SigningSize()),
		doorkeeper.RegisterIssuer(cfg.Doorkeeper.Issuer),
		doorkeeper.RegisterDuration(cfg.Doorkeeper.Duration),
		doorkeeper.RegisterPath(cfg.Doorkeeper.Path),
	)

	// Composers .-.
	serviceComposer := composer.NewServiceComposer(dk)
	repoComposer := composer.NewRepoComposer(pg, cfg.App.Environment)
	usecaseComposer := composer.NewUseCaseComposer(repoComposer, serviceComposer)

	// Http
	deliveree := gin.Default()
	v1.NewRouter(deliveree, logger, usecaseComposer)
	httpserver.NewServer(deliveree)
}
