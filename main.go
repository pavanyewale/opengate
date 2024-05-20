package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"opengate/cache"
	"opengate/config"
	"opengate/controller"
	"opengate/repository"
	"opengate/services"

	"github.com/gofreego/goutils/configutils"
)

var (
	configPath string
	env        string
)

func init() {
	flag.StringVar(&configPath, "configPath", "./", "Path to the configuration file")
	flag.StringVar(&env, "env", "dev", "environment")
	flag.Parse()
}

func main() {
	// Use the specified configPath or the default value if not provided
	fmt.Printf("Using Env: %s\n & Config file path : %s", env, configPath)
	var conf config.Config
	// Get the configuration
	err := configutils.ReadConfig(fmt.Sprintf("%s%s.yaml", configPath, env), &conf)
	if err != nil {
		log.Fatalf("Error getting configuration: %v", err)
	}

	bytes, _ := json.MarshalIndent(conf, "", "    ")
	log.Printf("config: \n %s", bytes)

	ctx := context.Background()

	cache := cache.NewCache(ctx, &conf.Cache)
	repo := repository.NewRepository(ctx, conf.Repository)
	srvFactory := services.NewServiceFactory(ctx, conf.Service, repo, cache)
	ctrl := controller.NewController(ctx, &conf.Controller, srvFactory)
	if err = ctrl.Listen(ctx); err != nil {
		log.Panic(err)
	}
}
