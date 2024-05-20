package config

import (
	"opengate/cache"
	"opengate/controller"
	"opengate/repository"
	"opengate/services"
)

type Config struct {
	Name       string
	Build      string
	Controller controller.Config
	Service    services.Config
	Repository repository.Config
	Cache      cache.Config
}
