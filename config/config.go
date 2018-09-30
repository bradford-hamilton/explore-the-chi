package config

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"github.com/spf13/viper"
)

// EnvConfig is a struct of the applications ENV variables
type EnvConfig struct {
	PORT           string
	Region         string `yaml:"region"`
	DynamoEndpoint string `yaml:"dynamoEndpoint"`
}

// DBConfig is a struct containing DB which is a connection to the application's
// dynamo db database
type DBConfig struct {
	DB *dynamo.DB
}

// NewDynamo is used to generate a Dynamo DB connection to be passed around the
// application
func NewDynamo() *DBConfig {
	config := DBConfig{}
	envVariables, err := NewEnvConfig()

	if err != nil {
		log.Panicln("Configuration error", err)
	}

	config.DB = dynamo.New(session.New(), &aws.Config{
		Region:   aws.String(envVariables.Region),
		Endpoint: aws.String(envVariables.DynamoEndpoint),
	})

	createTablesIfNotExist(&config)

	return &config
}

func createTablesIfNotExist(dbConn *DBConfig) {
	type BtcTransactionTable struct {
		TransactionID string `dynamo:"Id,hash"`
	}

	err := dbConn.DB.
		CreateTable("btc-transaction", BtcTransactionTable{}).
		Provision(1, 1).
		Run()
	if err != nil {
		fmt.Println(err)
	}
}

// NewEnvConfig is used to generate a an EnvConfig struct with the applications
// env variables and returns the address
func NewEnvConfig() (*EnvConfig, error) {
	envVariables, err := initViper()

	if err != nil {
		return &envVariables, err
	}

	return &envVariables, err
}

func initViper() (EnvConfig, error) {
	viper.SetConfigName("env") // configuration filename without the extension
	viper.AddConfigPath(".")   // search root directory for the configuration file

	err := viper.ReadInConfig() // find and read the config file
	if err != nil {
		return EnvConfig{}, err
	}

	viper.SetDefault("PORT", "4000")
	viper.SetDefault("region", "us-west-2")
	viper.SetDefault("dynamoEndpoint", "http://localhost:8000")

	var envVariables EnvConfig
	err = viper.Unmarshal(&envVariables)

	return envVariables, err
}
