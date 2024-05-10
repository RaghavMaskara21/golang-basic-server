package config

import (
	"encoding/json"
	"fmt"
	"hayday/server/internal/logger"
	"os"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"

	"github.com/spf13/viper"
)

func LoadConfig() serverConfig {
	logger.InitiateLogger()
	var appEnv string = os.Getenv("APP_ENV")
	configFileName := CONFIG_FILENAME
	requestId, _ := uuid.NewUUID()
	log := logger.Log.WithFields(map[string]interface{}{
		"EVENT":      "LOADING_EVN_VALUES",
		"REQUEST_ID": requestId,
	})

	if strings.ToUpper(appEnv) != "LOCAL" {
		loadAWSSecreteValues(log)
	} else {
		//ignore aws secrete manager envs for local run
		// load the local config file for local runs
		configFileName = fmt.Sprintf("%s.%s", appEnv, configFileName)
	}

	log.Infof("Loading the configuration from the config file : FILE PATH : %s/%s.%s", CONFIG_FILEPATH, configFileName, CONFIG_FILE_EXTENSION)
	viper.AddConfigPath(CONFIG_FILEPATH)
	viper.SetConfigName(configFileName)
	viper.SetConfigType(CONFIG_FILE_EXTENSION)

	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("failed to read the configuration : ERROR : %s", err)
	}

	err = viper.Unmarshal(&EnvValues)
	if err != nil {
		log.Fatalf("failed to unmarshal the configuration value to the map : ERROR : %s", err)
	}

	if err := validator.New().Struct(&EnvValues); err != nil {
		log.Fatalf("Missing required configuration values : ERROR : %s", err)
	}

	log.Infof("Configuration loaded successfully")
	return EnvValues
}

func loadAWSSecreteValues(log *logger.LoggerEvent) {
	log.Infof("Loading the secrete manager env values to the process environment...")
	secreteManagerEnvValue, ok := os.LookupEnv("SECRETE_MANAGER_ENVS")
	if ok {
		err := json.Unmarshal([]byte(secreteManagerEnvValue), &EnvValues)
		if err != nil {
			log.Fatalf("failed to unmarshal the SECRET_MANAGER_ENVS to the environment variable : ERROR : %s", err)
		}
		v := reflect.ValueOf(EnvValues)
		for i := 0; i < v.NumField(); i++ {
			os.Setenv(reflect.TypeOf(EnvValues).Field(i).Name, v.Field(i).String())
		}
		log.Infof("Successfully loaded all the SECRET_MANAGER_ENVS key on the process environment...")
	} else {
		log.Fatalf("SECRET_MANAGER_ENVS is missing on the environment variable...")
	}
}
