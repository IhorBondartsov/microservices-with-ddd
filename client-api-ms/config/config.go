package config


import (
	"flag"
	"fmt"
	"sync"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)
// clientAPIMSConfiguration implement pattern Singelton.
// We have only one config instance in our service.
// I think its good way for using this pattern.
// But I am doing its just for learning.
type clientAPIMSConfiguration struct {
	GRPCAddress  string `ignored:"true"`
	FilePath     string `required:"true"`
}

var config *clientAPIMSConfiguration
var once sync.Once

func GetConfig()*clientAPIMSConfiguration{
	once.Do(func(){
		config = &clientAPIMSConfiguration{}
	})
	return  config
}

func (cfg *clientAPIMSConfiguration) ResolveConfig() error {
	// Process CLI arguments
	var (
		grpcAddress string
		envFile     string
	)

	flag.StringVar(&grpcAddress, "grpc-address", ":50051", "Address to run server on")
	flag.StringVar(&envFile, "env", "", "External environment file")
	flag.Parse()

	//
	// Load env. from file if provided
	//
	if envFile != "" {
		if err := godotenv.Load(envFile); err != nil {
			return fmt.Errorf("failed to load environment from [%s], %s", envFile, err)
		}
	}

	if err := envconfig.Process("API_MS", cfg); err != nil {
		return err
	}

	cfg.GRPCAddress = grpcAddress

	return nil
}
